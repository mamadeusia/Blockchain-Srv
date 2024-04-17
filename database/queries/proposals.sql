-- name: CreateProposal :exec
INSERT INTO proposals (
        "id",
        "propose_id",
        "description",
        "candidate_id",
        "candidate_address",
        "token_offer",
        "merkle_root",
        "tx_hash",
        "validator_count",
        "proposal_type",
        "proposal_status"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
-- TODO :: update proposals of the status has changed