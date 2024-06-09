-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS times (
    id SERIAL PRIMARY KEY,
    amount BIGINT NOT NULL,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS times;
-- +goose StatementEnd
