-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts
(
    id         UUID PRIMARY KEY,
    email      TEXT UNIQUE              NOT NULL,
    password   TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE devices
(
    id          UUID PRIMARY KEY,
    account_id  UUID REFERENCES accounts (id) NOT NULL,
    os          TEXT                          NOT NULL,
    name        TEXT                          NOT NULL,
    fingerprint TEXT                          NOT NULL
);

CREATE TABLE refresh_tokens
(
    id         UUID PRIMARY KEY,
    device_id  UUID REFERENCES devices (id) NOT NULL,
    issued_at  TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;

DROP TABLE devices;

DROP TABLE accounts;
-- +goose StatementEnd
