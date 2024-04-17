package main

import (
	"context"

	"github.com/mamadeusia/BlockchainSrv/contract"
	"github.com/mamadeusia/BlockchainSrv/handler"
	blockchainService "github.com/mamadeusia/BlockchainSrv/service/blockchain"

	"github.com/mamadeusia/BlockchainSrv/config"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	postgresClient "github.com/mamadeusia/BlockchainSrv/client/postgres"
	blockchainPostgres "github.com/mamadeusia/BlockchainSrv/domain/blockchain/postgres"
	pb "github.com/mamadeusia/BlockchainSrv/proto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "blockchainsrv"
	version = "latest"
)

func main() {
	// load config from .env
	if err := config.Load(); err != nil {
		logger.Fatal("failed to load config")
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create service
	srv := micro.NewService(
		micro.Context(ctx),
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.BeforeStop(func() error {
			logger.Infof("shutting down %s service", service)
			cancel()
			return nil
		}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	provider, err := ethClient.Dial(config.BlockchainNodeAddress())
	if err != nil {
		logger.Fatal("failed to dial node: ", err)
	}

	chainID, err := provider.ChainID(ctx)
	if err != nil {
		logger.Fatal("failed to dial node: ", err)
	}

	privateKey, err := crypto.HexToECDSA(config.ContractDeployerPrivateKey())
	if err != nil {
		logger.Fatal("failed to convert privateKey: ", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	logger.Error("address: ", address.Hex())

	// create instance of StrikeGovernor contract
	governorContractAddress := common.HexToAddress(config.ContractGovernorAddress())
	governorContract, err := contract.NewStrikeGovernor(governorContractAddress, provider)
	if err != nil {
		logger.Fatal("failed to create NewStrikeGovernor: ", err)
	}

	blockchainPostgresInstance := blockchainPostgres.NewRepository(postgresClient.NewClient(ctx, config.PostgresURL()))
	// create instance of blockchain grpc handler
	blockchainGRPCHandler := &handler.BlockchainHandler{
		BlockchainService: &blockchainService.BlockchainServiceStruct{
			Contract:           governorContract,
			Provider:           provider,
			ChainID:            chainID,
			PrivateKey:         privateKey,
			BlockchainPostgres: blockchainPostgresInstance,
		},
	}

	// Register handler
	err = pb.RegisterBlockchainSrvHandler(srv.Server(), blockchainGRPCHandler)
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
