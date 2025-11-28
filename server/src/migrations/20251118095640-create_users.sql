
-- +migrate Up
CREATE TABLE users (
    id VARCHAR(64) NOT NULL,
    mail_address VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL,
    gender ENUM('male', 'female') NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    is_mailmagazine BOOLEAN NOT NULL DEFAULT FALSE,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;
