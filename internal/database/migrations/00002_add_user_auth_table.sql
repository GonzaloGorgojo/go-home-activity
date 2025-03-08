-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS UserToken (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "refreshToken" TEXT NOT NULL,
    "userId" INTEGER NOT NULL,
    FOREIGN KEY ("userId") REFERENCES User(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS UserToken;
-- +goose StatementEnd
