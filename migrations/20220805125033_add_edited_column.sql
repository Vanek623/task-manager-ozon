-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE tasks ADD IF NOT EXISTS edited timestamp;
UPDATE tasks SET edited = created WHERE tasks.edited ISNULL;
ALTER TABLE tasks ALTER edited SET DEFAULT now();

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE tasks DROP IF EXISTS edited;
