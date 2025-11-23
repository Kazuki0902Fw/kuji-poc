
-- +migrate Up
CREATE TABLE purchase_transactions (
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    ip_category_id VARCHAR(64) NOT NULL,
    purchase_quantity INT NOT NULL,
    purchase_price DECIMAL(10, 2) NOT NULL,
    payment_method ENUM('credit_card', 'bank_transfer', 'other') NOT NULL,
    provider_transaction_id VARCHAR(255),
    status ENUM('pending', 'paid', 'failed', 'cancelled') NOT NULL DEFAULT 'pending',
    paid_at DATETIME,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (ip_category_id) REFERENCES intellectual_property_categories(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE purchase_transactions;
