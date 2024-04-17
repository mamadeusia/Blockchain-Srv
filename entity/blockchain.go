package entity

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	pb "github.com/mamadeusia/BlockchainSrv/proto"
	"go-micro.dev/v4/errors"
)

type ProposeRequest struct {
	Description      string
	CandidateId      string
	CandidateAddress common.Address
	TokenOffer       int32
	Validators       []string
}
type ProposeResponse struct {
	TXHash     string
	ProposalId string
}

func ProposeProtoToRequest(req *pb.ProposeRequest) (*ProposeRequest, error) {
	zeroAddress := common.Address{}
	candidateAddress := common.HexToAddress(req.Candidate)
	if candidateAddress == zeroAddress {
		return nil, errors.BadRequest("entity-0002", "candidate is not valid")
	}

	tokenOffer, err := strconv.ParseInt(req.TokenOffer, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("entity-0003", "tokenOffer is not valid")
	}

	return &ProposeRequest{
		Description:      req.Description,
		CandidateId:      req.CandidateID,
		CandidateAddress: candidateAddress,
		TokenOffer:       int32(tokenOffer),
		Validators:       req.Validators,
	}, nil

}

type ProofsRequest struct {
	VoterAddress common.Address
	ProposeID    *big.Int
	Support      uint8
}

func ProofsProtoToRequest(req *pb.ProofsRequest) (*ProofsRequest, error) {
	zeroAddress := common.Address{}
	voterAddress := common.HexToAddress(req.VoterAddress)
	if voterAddress == zeroAddress {
		return nil, errors.BadRequest("entity-0002", "candidate is not valid")
	}

	proposeID, ok := big.NewInt(0).SetString(req.ProposeID, 10)
	if !ok {
		return nil, errors.BadRequest("entity-0003", "tokenOffer is not valid")
	}

	return &ProofsRequest{
		VoterAddress: voterAddress,
		ProposeID:    proposeID,
	}, nil
}
