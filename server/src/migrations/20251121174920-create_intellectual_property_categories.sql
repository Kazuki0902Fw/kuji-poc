
-- +migrate Up
CREATE TABLE intellectual_property_categories (
    id VARCHAR(64) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    price DECIMAL NOT NULL,
    sales_start_date DATE NOT NULL,
    sales_end_date DATE NOT NULL,
    fee DECIMAL NOT NULL,
    precautions TEXT,
    is_hidden BOOLEAN NOT NULL DEFAULT FALSE,
    img_url VARCHAR(512),
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE intellectual_property_categories;
