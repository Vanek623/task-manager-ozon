-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

BEGIN;
DO $$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM information_schema.columns WHERE table_name = 'tasks' AND column_name='edited') THEN
            ALTER TABLE tasks ADD edited timestamp DEFAULT now();
            UPDATE tasks SET edited = created;
        END IF;
    END
$$;
COMMIT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE tasks DROP IF EXISTS edited;
