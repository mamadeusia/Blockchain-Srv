package handler

import (
	"context"

	"github.com/mamadeusia/BlockchainSrv/entity"
	pb "github.com/mamadeusia/BlockchainSrv/proto"
	blockchainService "github.com/mamadeusia/BlockchainSrv/service/blockchain"

	"go-micro.dev/v4/logger"
)

type BlockchainHandler struct {
	BlockchainService *blockchainService.BlockchainServiceStruct
}

func (g *BlockchainHandler) Propose(ctx context.Context, req *pb.ProposeRequest, rsp *pb.ProposeResponse) error {
	data, err := entity.ProposeProtoToRequest(req)
	if err != nil {
		logger.Error(err)
		return err
	}

	res, err := g.BlockchainService.Propose(ctx, data)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.TxHash = res.TXHash
	rsp.ProposeID = res.ProposalId

	return nil
}

func (g *BlockchainHandler) Proofs(ctx context.Context, req *pb.ProofsRequest, rsp *pb.ProofsResponse) error {
	data, err := entity.ProofsProtoToRequest(req)
	if err != nil {
		logger.Error(err)
		return err
	}

	res, err := g.BlockchainService.Proofs(ctx, data)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.Proofs = res

	return nil
}
