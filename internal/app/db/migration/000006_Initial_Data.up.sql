-- Insert sample data into the Items table
INSERT INTO Products (title) VALUES
    ('ноутбук'),
    ('телевизор'),
    ('телефон'),
    ('системный блок'),
    ('часы'),
    ('микрофон');

-- Insert sample data into the Shelves table
INSERT INTO Shelves (shelf_name) VALUES
    ('A'),
    ('Б'),
    ('В'),
    ('Г'),
    ('Д'),
    ('Е'),
    ('Ж'),
    ('З'),
    ('Й'),
    ('К'),
    ('Л'),
    ('М'),
    ('Н'),
    ('О'),
    ('П');

-- Insert sample data into the Orders table
INSERT INTO Orders (
order_number, product_name, quantity, product_id
) VALUES
    ('10', 'ноутбук', 2, 1),
    ('11', 'monitor', 3, 2),
    ('10', 'телефон', 1, 3),
    ('14', 'ноутбук', 3, 1),
    ('14', 'системный блок', 4, 4),
    ('15', 'часы', 1, 5),
    ('10', 'микрофон', 1, 6);

-- Insert sample data into the Item_Shelves table
INSERT INTO Product_Shelves (id, shelf_id, is_primary) VALUES
    (1, 1, true),
    (1, 1, true),
    (3, 2, true),
    (3, 8, false),
    (3, 3, false),
    (2, 1, true),
    (4, 7, true),
    (5, 7, true),
    (5, 1, false),
    (6, 7, true);
