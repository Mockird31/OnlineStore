INSERT INTO category (title) VALUES
('Electronics'),
('Smartphones'),
('Laptops'),
('Books'),
('Games'),
('Health'),
('Clothing'),
('Home'),
('Sports');

INSERT INTO item (title, description, price, image_url, quantity, is_active) VALUES
('Smartphone XYZ', 'Latest model with advanced features', 35999, '/default_item_image.png', 50, TRUE),
('Laptop ABC', 'High performance laptop for professionals', 89999.99, '/default_item_image.png', 30, TRUE),
('Gaming Console', 'Next-gen gaming console with exclusive games', 40000, '/default_item_image.png', 20, TRUE),
('Fitness Tracker', 'Monitor your health and fitness activities', 3000, '/default_item_image.png', 100, TRUE),
('Winter Jacket', 'Warm and stylish winter jacket', 2500, '/default_item_image.png', 75, TRUE);

INSERT INTO item_category (item_id, category_id) VALUES
(1, 2), -- Smartphone XYZ in Smartphones
(2, 3), -- Laptop ABC in Laptops
(3, 5), -- Gaming Console in Games
(4, 6), -- Fitness Tracker in Health
(5, 7); -- Winter Jacket in Clothing