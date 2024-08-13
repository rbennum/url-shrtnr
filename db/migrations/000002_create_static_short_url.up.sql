CREATE TABLE IF NOT EXISTS static_url (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL
);
INSERT INTO static_url (url)
VALUES ('localhost:8088');