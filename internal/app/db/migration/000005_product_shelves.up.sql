CREATE TABLE IF NOT EXISTS Product_Shelves (
    id INT REFERENCES Products(id),
    shelf_id INT REFERENCES Shelves(shelf_id),
    is_primary BOOLEAN
);