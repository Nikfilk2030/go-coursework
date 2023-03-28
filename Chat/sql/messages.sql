CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    username   TEXT NOT NULL,
    message    TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
