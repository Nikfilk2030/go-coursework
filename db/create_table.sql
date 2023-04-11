CREATE TABLE tasks
(
    id            BIGSERIAL PRIMARY KEY,
    created_at    timestamptz NOT NULL DEFAULT now(),
    topic         TEXT        NOT NULL,
    description   TEXT        NOT NULL DEFAULT '',
    creator_login TEXT        NOT NULL,
    status_id     INT         NOT NULL DEFAULT 0
)