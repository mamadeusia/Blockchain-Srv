-- liquibase formatted sql
-- changeset omid:2
CREATE TABLE proposal_validators (
    "propose_id" BYTEA NOT NULL,
    "validator_address" VARCHAR(42) NOT NULL
);
ALTER TABLE proposal_validators
ADD FOREIGN KEY (propose_id) REFERENCES proposals (propose_id);
-- rollback  DROP TABLE IF EXISTS proposal_validators;