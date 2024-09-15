-- +goose Up
CREATE TABLE kv (
    value TEXT,
    key TEXT GENERATED ALWAYS AS (json_extract(value, '$.key')) VIRTUAL UNIQUE NOT NULL
);

CREATE INDEX kv_key ON kv (key);

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL
);

-- +goose Down
DROP TABLE kv;
DROP TABLE users;