CREATE TABLE users
(
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    contact VARCHAR(255) NOT NULL,
    role INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);