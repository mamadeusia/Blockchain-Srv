package config

type Contract struct {
	Governor Governor
	Deployer Deployer
}

type Governor struct {
	Address string
}

type Deployer struct {
	PrivateKey string
}

func ContractGovernorAddress() string {
	return cfg.Contract.Governor.Address
}

func ContractDeployerPrivateKey() string {
	return cfg.Contract.Deployer.PrivateKey
}
