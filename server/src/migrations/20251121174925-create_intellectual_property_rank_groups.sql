
-- +migrate Up
CREATE TABLE intellectual_property_rank_groups (
    id VARCHAR(64) NOT NULL,
    ip_category_id VARCHAR(64) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `rank` ENUM('S', 'A', 'B', 'C', 'D', 'E') NOT NULL,
    emission_rate INT NOT NULL,
    explanation TEXT,
    is_hidden BOOLEAN NOT NULL DEFAULT FALSE,
    comments TEXT,
    img_url VARCHAR(512),
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (ip_category_id) REFERENCES intellectual_property_categories(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE intellectual_property_rank_groups;
