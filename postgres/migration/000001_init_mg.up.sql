CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    email      VARCHAR(255),
    password   VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    is_deleted BOOLEAN
);

INSERT INTO users (first_name, last_name, email, password, created_at, updated_at, is_deleted)
VALUES ('Admin', 'User', 'admin@example.com', 'password', now(), now(), FALSE),
       ('John', 'Doe', 'john@doe.com', 'string-password', now() ,now(), FALSE),
       ('Jennifer', 'Lawrence', 'jen@star.com', 'secret', now() ,now(), TRUE)
;