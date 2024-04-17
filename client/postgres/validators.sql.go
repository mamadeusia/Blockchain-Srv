// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: validators.sql

package postgres

import (
	"context"
)

const createValidator = `-- name: CreateValidator :exec
INSERT INTO proposal_validators(
        "propose_id",
        "validator_address"
    )
VALUES ($1, $2)
`

type CreateValidatorParams struct {
	ProposeID        []byte `db:"propose_id" json:"propose_id"`
	ValidatorAddress string `db:"validator_address" json:"validator_address"`
}

func (q *Queries) CreateValidator(ctx context.Context, arg CreateValidatorParams) error {
	_, err := q.db.Exec(ctx, createValidator, arg.ProposeID, arg.ValidatorAddress)
	return err
}

const validatorsByProposeID = `-- name: ValidatorsByProposeID :many
SELECT v.validator_address
FROM Proposals AS p
    RIGHT JOIN proposal_validators AS v ON v.propose_id = p.propose_id
WHERE p.propose_id = $1
`

func (q *Queries) ValidatorsByProposeID(ctx context.Context, proposeID []byte) ([]string, error) {
	rows, err := q.db.Query(ctx, validatorsByProposeID, proposeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var validator_address string
		if err := rows.Scan(&validator_address); err != nil {
			return nil, err
		}
		items = append(items, validator_address)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
