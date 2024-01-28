-- +goose Up
-- +goose StatementBegin
 CREATE TABLE tasks (
     id BIGSERIAL PRIMARY KEY,
     name VARCHAR(1024) NOT NULL,
     description VARCHAR(4096) DEFAULT NULL,
     story_points INT NOT NULL DEFAULT 0,
     completed BOOLEAN NOT NULL DEFAULT FALSE,
     employee_id BIGINT REFERENCES employees(id)
 );

CREATE TABLE todo_list (
    id BIGSERIAL PRIMARY KEY,
    label VARCHAR(1024) NOT NULL,
    priority INT NOT NULL DEFAULT 0,
    employee_id BIGINT REFERENCES employees(id),
    completed BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE goal_tracks (
  id BIGSERIAL PRIMARY KEY,
  track VARCHAR(1024) NOT NULL
);

CREATE TABLE goals (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    description VARCHAR(4096) DEFAULT NULL,
    employee_id BIGINT REFERENCES employees(id),
    track_id BIGINT NOT NUll REFERENCES goal_tracks(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  goals, goal_tracks, todo_list, tasks;
-- +goose StatementEnd
