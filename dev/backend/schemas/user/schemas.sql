CREATE TABLE IF NOT EXISTS tb_user (
    id SERIAL PRIMARY KEY,
    
    username VARCHAR(255) NOT NULL UNIQUE,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,

    session_token TEXT,
    csrf_token    TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);