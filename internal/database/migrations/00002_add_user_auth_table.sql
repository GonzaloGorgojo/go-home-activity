-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS UserToken (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "refreshToken" TEXT NOT NULL,
    "userEmail" TEXT NOT NULL UNIQUE,
    "createdAt" DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now')),
    "updatedAt" DATETIME NULL,
    FOREIGN KEY ("userEmail") REFERENCES User(email) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS UserToken;
-- +goose StatementEnd
