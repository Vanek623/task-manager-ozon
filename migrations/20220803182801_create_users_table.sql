-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists tasks
(
    id          bigserial primary key,
    title       varchar(64) not null,
    description varchar(256),
    created     timestamp default now()
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists tasks;
-- +goose StatementEnd
