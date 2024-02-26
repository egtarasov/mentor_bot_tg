-- +goose Up
-- +goose StatementBegin
create table questions(
    id bigserial primary key,
    text varchar(4096) not null,
    user_id bigint references employees(id),
    created_at timestamp default now(),
    answered_at timestamp default null,
    answered_by bigint default null,
    answer varchar(4096) default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table questions;
-- +goose StatementEnd
