-- +goose Up
-- +goose StatementBegin

create table notifications
(
    id                bigserial primary key,
    message           varchar(4096) not null,
    photo_path        varchar(4096) default null,
    notification_time time not null,
    -- 0 - Sunday, 6 - Saturday.
    day_of_week       int not null check (0 <= notifications.day_of_week)
        check ( notifications.day_of_week <= 6 ),
    repeat_time       interval not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table notifications;
-- +goose StatementEnd
