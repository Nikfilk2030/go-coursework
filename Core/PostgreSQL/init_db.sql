CREATE TABLE messages
(
    id      bigserial PRIMARY KEY,
    login   varchar(255) NOT NULL,
    type    bigint NOT NULL,
    text    TEXT NOT NULL,
    parent  bigint,
    FOREIGN KEY (login) REFERENCES users (login),
    FOREIGN KEY (parent) REFERENCES messages (id)
);


CREATE TABLE users
(
    login    varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL
);
