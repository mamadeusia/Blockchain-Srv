package postgres

import (
	"context"

	"github.com/google/uuid"
	postgresClient "github.com/mamadeusia/BlockchainSrv/client/postgres"
	"go-micro.dev/v4/errors"
)

type Blockchain struct {
	ID               uuid.UUID
	ProposeID        []byte
	Description      string
	CandidateID      string
	CandidateAddress []byte
	TokenOffer       int64
	MerkleRoot       []byte
	TxHash           []byte
	ValidatorCount   int64
	ProposalType     string
	ProposalStatus   string
}

func (p *PostgresRepository) CreateProposals(ctx context.Context, blockchain Blockchain) error {
	params := postgresClient.CreateProposalParams{
		ID:               blockchain.ID,
		ProposeID:        blockchain.ProposeID,
		Description:      blockchain.Description,
		CandidateID:      blockchain.CandidateID,
		CandidateAddress: blockchain.CandidateAddress,
		TokenOffer:       blockchain.TokenOffer,
		MerkleRoot:       blockchain.MerkleRoot,
		TxHash:           blockchain.TxHash,
		ValidatorCount:   blockchain.ValidatorCount,
		ProposalType:     postgresClient.ProposalType(blockchain.ProposalType),
		ProposalStatus:   postgresClient.NullProposalStatus{ProposalStatus: postgresClient.ProposalStatusSuccess, Valid: true},
	}

	if err := p.Postgres.Queries.CreateProposal(ctx, params); err != nil {
		return errors.InternalServerError("postgres-0001", "failed to create proposals: %v", err)
	}

	return nil
}

func (p *PostgresRepository) CreateValidators(ctx context.Context, proposeID []byte, validatorAddress string) error {
	params := postgresClient.CreateValidatorParams{
		ProposeID:        proposeID,
		ValidatorAddress: validatorAddress,
	}

	if err := p.Postgres.Queries.CreateValidator(ctx, params); err != nil {
		return errors.InternalServerError("postgres-0002", "failed to create validators: %v", err)
	}

	return nil
}

func (p *PostgresRepository) GetValidatorsByProposalID(ctx context.Context, proposeID []byte) ([]string, error) {
	validators, err := p.Postgres.Queries.ValidatorsByProposeID(ctx, proposeID)
	if err != nil {
		return nil, errors.InternalServerError("postgres-0001", "failed to get validators: %v", err)
	}

	return validators, nil
}
