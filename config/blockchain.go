package config

type Blockchain struct {
	Node       Node
	PrivateKey string
}

type Node struct {
	Address string
}

func BlockchainNodeAddress() string {
	return cfg.Blockchain.Node.Address
}
