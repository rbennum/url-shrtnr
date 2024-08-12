CREATE TABLE IF NOT EXISTS link_mappers (
    id SERIAL PRIMARY KEY,
    url VARCHAR(400) NOT NULL,
    short_tag VARCHAR(5) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);