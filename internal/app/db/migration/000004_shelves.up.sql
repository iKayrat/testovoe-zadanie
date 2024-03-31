-- Create the Shelves table
CREATE TABLE IF NOT EXISTS Shelves (
    shelf_id SERIAL PRIMARY KEY,
    shelf_name VARCHAR(255) NOT NULL
);

