CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    password VARCHAR(20) NOT NULL
);
INSERT INTO users VALUES(1, 'password1');