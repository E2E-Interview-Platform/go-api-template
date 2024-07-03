CREATE TABLE `users` (
    id VARCHAR(64) PRIMARY KEY,
    email VARCHAR(320) NOT NULL UNIQUE, -- 320 is the actual limit of a email
    password VARCHAR(64) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);
