-- +goose Up
-- +goose StatementBegin
CREATE TABLE token (
    id SERIAL PRIMARY KEY,
    guid CHARACTER VARYING(255),
    refresh_token BYTEA NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS token CASCADE;
-- +goose StatementEnd
