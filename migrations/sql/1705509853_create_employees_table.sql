-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    hire_date DATE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- [your SQL script here]
DROP TABLE IF EXISTS employees;
