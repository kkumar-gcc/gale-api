CREATE TABLE password_reset_tokens (
   email VARCHAR(255) NOT NULL PRIMARY KEY,
   token VARCHAR(255) NOT NULL,
   created_at DATETIME(3) NOT NULL,
   updated_at DATETIME(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;