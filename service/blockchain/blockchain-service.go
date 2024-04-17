package blockchainservice

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	merkleTree "github.com/mamadeusia/simplemerkletree/merkletree"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/mamadeusia/BlockchainSrv/contract"
	blockchainPostgres "github.com/mamadeusia/BlockchainSrv/domain/blockchain/postgres"
	"github.com/mamadeusia/BlockchainSrv/entity"
	"go-micro.dev/v4/errors"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/sha3"
)

type BlockchainServiceStruct struct {
	Contract           *contract.StrikeGovernor
	Provider           *ethClient.Client
	ChainID            *big.Int
	PrivateKey         *ecdsa.PrivateKey
	BlockchainPostgres *blockchainPostgres.PostgresRepository
}

func (b *BlockchainServiceStruct) Propose(
	ctx context.Context,
	propose *entity.ProposeRequest,
) (*entity.ProposeResponse, error) {

	root := merkleTree.BuildTree(propose.Validators).GetHash()

	opts, err := bind.NewKeyedTransactorWithChainID(b.PrivateKey, b.ChainID)
	if err != nil {
		return nil, errors.InternalServerError("bc-0001", "failed to create tx option: %v", err)
	}

	tx, err := b.Contract.Propose(
		opts,
		propose.Description,
		sha3.Sum256([]byte(propose.CandidateId)),
		propose.CandidateAddress,
		uint32(propose.TokenOffer),
		root,
		uint32(len(propose.Validators)),
		1,
	)
	if err != nil {
		return nil, errors.InternalServerError("bc-0002", "failed to propose: %v", err)
	}

	receipt, err := bind.WaitMined(ctx, b.Provider, tx)
	if err != nil {
		return nil, errors.InternalServerError("bc-0003", "failed to get receipt: %v", err)
	}

	var proposal *contract.StrikeGovernorProposalCreated
	for _, log := range receipt.Logs {
		proposal, err = b.Contract.ParseProposalCreated(*log)
		if err == nil {
			break
		}
	}

	if proposal == nil {
		return nil, errors.InternalServerError("bc-0004", "failed to parse ProposalCreated: %v", err)
	}
	proposeUUID := uuid.New()
	txHash := tx.Hash()
	err = b.BlockchainPostgres.CreateProposals(ctx, blockchainPostgres.Blockchain{
		ID:               proposeUUID,
		ProposeID:        proposal.ProposalId.Bytes(),
		Description:      propose.Description,
		CandidateID:      propose.CandidateId,
		CandidateAddress: proposal.Candidate[:],
		TokenOffer:       int64(proposal.TokenOffer),
		MerkleRoot:       root[:],
		TxHash:           txHash[:],
		ValidatorCount:   int64(len(propose.Validators)),
		ProposalType:     entity.DONATION,
	})
	if err != nil {
		return nil, errors.InternalServerError("bc-0005", "failed to store data in blockchain: %v", err)
	}

	for _, validator := range propose.Validators {
		if err = b.BlockchainPostgres.CreateValidators(ctx, proposal.ProposalId.Bytes(), validator); err != nil {
			return nil, errors.InternalServerError("bc-0006", "failed to store data in Validators: %v", err)
		}
	}

	return &entity.ProposeResponse{
		TXHash:     tx.Hash().Hex(),
		ProposalId: proposal.ProposalId.String(),
	}, nil
}

func (b *BlockchainServiceStruct) Proofs(
	ctx context.Context,
	vote *entity.ProofsRequest,
) ([][]byte, error) {
	validators, err := b.BlockchainPostgres.GetValidatorsByProposalID(ctx, vote.ProposeID.Bytes())
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	proofs := merkleTree.GetProof(vote.VoterAddress.Hex(), validators)

	pfs := make([][]byte, len(proofs))
	for i, proof := range proofs {
		proofHash := proof.GetHash()
		pfs[i] = proofHash[:]
	}

	return pfs, nil
}
