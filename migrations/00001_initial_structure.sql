-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "type" TEXT NOT NULL CHECK(type IN ('paid', 'free'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
