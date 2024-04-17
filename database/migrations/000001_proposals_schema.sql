-- liquibase formatted sql
-- changeset omid:1
CREATE TYPE proposal_type AS ENUM ('withdrawal', 'donation');
CREATE TYPE proposal_status AS ENUM ('failed', 'success');
CREATE TABLE proposals (
    "id" UUID PRIMARY KEY,
    "propose_id" BYTEA NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    "candidate_id" TEXT NOT NULL,
    "candidate_address" BYTEA NOT NULL,
    "token_offer" bigint NOT NULL,
    "merkle_root" BYTEA NOT NULL,
    "tx_hash" BYTEA NOT NULL,
    "validator_count" bigint NOT NULL,
    "proposal_type" proposal_type NOT NULL,
    "proposal_status" proposal_status,
    "created_at" timestamptz DEFAULT (now()) NOT NULL,
    "updated_at" timestamptz DEFAULT (now()) NOT NULL
);
-- rollback DROP TABLE IF EXISTS proposals;
-- rollback DROP TYPE IF EXISTS proposal_status;
-- rollback DROP TYPE IF EXISTS proposal_type;