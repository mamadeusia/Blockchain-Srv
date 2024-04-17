package postgres

import postgresClient "github.com/mamadeusia/BlockchainSrv/client/postgres"

type PostgresRepository struct {
	Postgres *postgresClient.PostgresRepository
}

func NewRepository(postgres *postgresClient.PostgresRepository) *PostgresRepository {
	return &PostgresRepository{Postgres: postgres}
}
