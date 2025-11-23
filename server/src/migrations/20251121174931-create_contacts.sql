
-- +migrate Up
CREATE TABLE contacts (
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    contact_type ENUM('inquiry', 'complaint', 'suggestion', 'other') NOT NULL,
    text TEXT NOT NULL,
    email VARCHAR(255) NOT NULL,
    file VARCHAR(512),
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE contacts;
