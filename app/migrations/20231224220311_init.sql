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
    is_admin          BOOLEAN               DEFAULT FALSE NOT NULL,
    parent_id BIGINT DEFAULT NULL
);

CREATE TABLE occupations (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    material VARCHAR(4096) NOT NULL
);

CREATE TABLE employees
(
    id                BIGSERIAL PRIMARY KEY,
    name              VARCHAR(100) NOT NULL,
    surname           VARCHAR(100)          DEFAULT NULL,
    telegram_id       BIGINT       NOT NULL UNIQUE,
    occupation_id     BIGINT REFERENCES occupations (id),
    grade             INT                   DEFAULT 17 NOT NULL,
    is_admin          BOOLEAN               DEFAULT FALSE NOT NULL,
    first_working_day date         NOT NULL DEFAULT NOW(),
    adaptation_end_at date         NOT NULL DEFAULT CURRENT_DATE + 90
);

-- Insert values
INSERT INTO actions (action)
-- TODO should use something better as names
VALUES ('get data'), ('show subsections'), ('complex');

INSERT INTO occupations (name, material)
-- TODO add all possible occupations
VALUES ('developer', 'Матерьял для разработчиков');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE commands, actions, employees, occupations;
-- +goose StatementEnd
