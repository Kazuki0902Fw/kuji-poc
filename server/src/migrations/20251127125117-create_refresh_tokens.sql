
-- +migrate Up
CREATE TABLE refresh_tokens (
  id                 VARCHAR(64)  NOT NULL,
  user_id            VARCHAR(64)  NOT NULL,
  token_hash         VARCHAR(255) NOT NULL,
  expires_at         DATETIME NOT NULL,
  
  created_at         DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at         DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,

  PRIMARY KEY (id),
  FOREIGN KEY fk_refresh_tokens_user_id (user_id) REFERENCES users(id),

  UNIQUE KEY uk_refresh_tokens_user_id (user_id)
);

-- +migrate Down
DROP TABLE refresh_tokens;
