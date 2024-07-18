-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE chats
(
    id         serial primary key,
    created_at timestamp not null default now(),
    updated_at timestamp
);

CREATE TABLE message
(
    id         serial primary key,
    author     int       not null,
    content    text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    chat_id    integer REFERENCES chats (id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
