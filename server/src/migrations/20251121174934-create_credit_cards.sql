
-- +migrate Up
CREATE TABLE credit_cards (
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE credit_cards;
