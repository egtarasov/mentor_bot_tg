-- +goose Up
-- +goose StatementBegin

-- Create tables
CREATE TABLE actions (
    id  BIGSERIAL PRIMARY KEY ,
    action VARCHAR(30) NOT NULL
);

CREATE TABLE commands (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE,
    action_id BIGINT REFERENCES actions(id),
    parent_id BIGINT REFERENCES commands(id)
);

CREATE TABLE occupations (
    id BIGSERIAL PRIMARY KEY,
    occupation VARCHAR(100) NOT NULL
);

CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) DEFAULT NULL,
    telegram_id BIGINT NOT NULL UNIQUE,
    occupation_id BIGINT REFERENCES occupations(id)
);

-- Insert values
INSERT INTO actions (action)
-- TODO should use something better as names
VALUES ('get data'), ('show subsections'), ('complex');

INSERT INTO occupations (occupation)
-- TODO add all possible occupations
VALUES ('developer');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employees, occupations, commands, actions;
-- +goose StatementEnd
