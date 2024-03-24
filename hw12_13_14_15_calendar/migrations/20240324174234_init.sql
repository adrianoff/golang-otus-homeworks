-- +goose Up
-- +goose StatementBegin
CREATE TABLE events
(
    id BIGSERIAL CONSTRAINT events_pk PRIMARY KEY,
    title VARCHAR (255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
