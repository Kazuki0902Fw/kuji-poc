
-- +migrate Up
CREATE TABLE purchase_histories (
    id VARCHAR(64) NOT NULL,
    purchase_transaction_id VARCHAR(64) NOT NULL,
    intellectual_property_id VARCHAR(64) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (purchase_transaction_id) REFERENCES purchase_transactions(id) ON DELETE CASCADE,
    FOREIGN KEY (intellectual_property_id) REFERENCES intellectual_properties(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE purchase_histories;
