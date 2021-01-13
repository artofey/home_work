-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE EXTENSION "uuid-ossp";

create table events (
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user" uuid NOT NULL,
    title character varying(200),
    descr text,
    start_date timestamp NOT NULL,
    end_date timestamp NOT NULL,
    notify_time timestamp
);

CREATE INDEX user_idx ON events ("user");
CREATE INDEX start_idx ON events (start_date);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP INDEX user_idx;
DROP INDEX start_idx;
DROP TABLE events;
DROP EXTENSION "uuid-ossp";
