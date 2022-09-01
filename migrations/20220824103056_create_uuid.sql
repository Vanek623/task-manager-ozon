-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE IF EXISTS tasks DROP IF EXISTS id;
ALTER TABLE IF EXISTS tasks ADD IF NOT EXISTS id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid();

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE IF EXISTS tasks DROP IF EXISTS uuid;
ALTER TABLE IF EXISTS tasks ADD IF NOT EXISTS id bigserial PRIMARY KEY;

