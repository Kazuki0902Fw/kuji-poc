
-- +migrate Up
CREATE TABLE shipping_addresses (
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    recipient_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    post_code VARCHAR(10) NOT NULL,
    prefecture VARCHAR(50) NOT NULL,
    address_1 VARCHAR(255) NOT NULL,
    address_2 VARCHAR(255),
    apartment_name VARCHAR(255),
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE shipping_addresses;
