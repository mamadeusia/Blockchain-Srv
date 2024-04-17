-- name: CreateValidator :exec
INSERT INTO proposal_validators(
        "propose_id",
        "validator_address"
    )
VALUES ($1, $2);
-- name: ValidatorsByProposeID :many
SELECT v.validator_address
FROM Proposals AS p
    RIGHT JOIN proposal_validators AS v ON v.propose_id = p.propose_id
WHERE p.propose_id = $1;
-- TODO : select if this specific validator exist and blong to wich proposal