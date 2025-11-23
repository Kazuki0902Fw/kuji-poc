
-- +migrate Up
CREATE TABLE intellectual_properties (
    id VARCHAR(64) NOT NULL,
    ip_rank_group_id VARCHAR(64) NOT NULL,
    name VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    is_hidden BOOLEAN NOT NULL DEFAULT FALSE,
    img_url VARCHAR(512),
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (ip_rank_group_id) REFERENCES intellectual_property_rank_groups(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE intellectual_properties;
