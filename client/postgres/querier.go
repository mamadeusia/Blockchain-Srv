// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgres

import (
	"context"
)

type Querier interface {
	CreateProposal(ctx context.Context, arg CreateProposalParams) error
	CreateValidator(ctx context.Context, arg CreateValidatorParams) error
	ValidatorsByProposeID(ctx context.Context, proposeID []byte) ([]string, error)
}

var _ Querier = (*Queries)(nil)
