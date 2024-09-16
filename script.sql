CREATE TABLE IF NOT EXISTS stocks (
    stockid SERIAL PRIMARY KEY,
    name TEXT,
    price INT,
    company TEXT
);
