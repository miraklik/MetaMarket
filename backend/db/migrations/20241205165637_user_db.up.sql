CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(55) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);