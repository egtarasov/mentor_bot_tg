-- +goose Up
-- +goose StatementBegin
CREATE TABLE materials (
    id BIGSERIAL PRIMARY KEY,
-- maximum size supported by telegram
    message VARCHAR(4096) NOT NULL,
    command_id BIGINT REFERENCES commands(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE materials;
-- +goose StatementEnd
