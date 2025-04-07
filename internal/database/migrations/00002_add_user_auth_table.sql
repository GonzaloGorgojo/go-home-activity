-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS UserToken (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "refreshToken" TEXT NOT NULL,
    "userID" INTEGER NOT NULL UNIQUE,
    "createdAt" DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now')),
    "updatedAt" DATETIME NULL,
    FOREIGN KEY ("userID") REFERENCES User(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS UserToken;
-- +goose StatementEnd
