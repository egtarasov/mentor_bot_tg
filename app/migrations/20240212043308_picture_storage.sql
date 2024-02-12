-- +goose Up
-- +goose StatementBegin
create table pictures (
    id bigserial primary key,
    path varchar(1024) not null,
    command_id bigint references commands(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pictures;
-- +goose StatementEnd
