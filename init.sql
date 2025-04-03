CREATE TABLE IF NOT EXISTS series (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    seasons INT,
    episodes INT,
    genre VARCHAR(100),
    status VARCHAR(50) DEFAULT 'To Watch',
    current_episode INT DEFAULT 0,
    score INT DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);