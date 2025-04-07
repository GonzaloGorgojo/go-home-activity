-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS User (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now')),
    "updatedAt" DATETIME NULL,
    "type" TEXT NOT NULL CHECK(type IN ('paid', 'free')),
    "status" TEXT NOT NULL CHECK(status IN ('active', 'suspended')) DEFAULT ('active')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS User;
-- +goose StatementEnd
