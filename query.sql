create database "Inventaris";


DROP TABLE IF EXISTS "Session" CASCADE;
DROP TABLE IF EXISTS "Items" CASCADE;
DROP TABLE IF EXISTS "Categories" CASCADE;
DROP TABLE IF EXISTS "Users" CASCADE;

CREATE TABLE IF NOT EXISTS "Users" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, 
    role VARCHAR(20) CHECK (role IN ('admin', 'user')) NOT NULL
);

CREATE TABLE IF NOT EXISTS "Categories" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS "Items" (
    id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
    photo_url VARCHAR(255)NOT NULL,
    price NUMERIC(12, 2) NOT NULL,
    purchase_date DATE NOT NULL,
    usage_days INTEGER NOT NULL,
	depreciation_rate NUMERIC(5, 2) DEFAULT 0,
    category_id INTEGER NOT NULL REFERENCES "Categories"(id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS "Session" (
    session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INTEGER NOT NULL REFERENCES "Users"(id) ON DELETE CASCADE, 
    login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);


-- Insert dummy categories into the Categories table with descriptions
INSERT INTO "Categories" (name, description)
VALUES
    ('Electronics', 'Devices and gadgets related to electronics, such as printers, laptops, etc.'),
    ('Computers', 'Includes desktops, laptops, and accessories related to computing.'),
    ('Mobile Phones', 'Smartphones and mobile devices of various brands.'),
    ('Furniture', 'Home and office furniture such as tables, chairs, desks, etc.'),
    ('Office Equipment', 'Items used in an office setting, including chairs, desks, and other supplies.'),
    ('Home Appliances', 'Household appliances including refrigerators, microwaves, and more.'),
    ('Entertainment', 'Items related to entertainment, such as TVs, projectors, and audio equipment.'),
    ('Kitchen Appliances', 'Appliances specifically used in the kitchen, such as blenders, ovens, etc.'),
    ('Cooling Systems', 'Air conditioners, fans, and other systems for cooling.');
    
-- Insert 20 dummy items into the Items table
INSERT INTO "Items" (name, category_id, photo_url, price, purchase_date, usage_days, depreciation_rate)
VALUES
    ('Printer Canon', 1, '/images/printer.png', 1500000, '2024-01-05', 10, 10),
    ('Laptop Dell', 2, '/images/laptop.png', 10000000, '2023-05-10', 20, 20),
    ('Smartphone Samsung', 3, '/images/smartphone.png', 8000000, '2023-08-15', 50, 15),
    ('Office Chair', 4, '/images/chair.png', 200000, '2023-11-01', 5, 15),
    ('Table Desk', 5, '/images/desk.png', 120000, '2023-02-20', 15, 17),
    ('Projector Epson', 1, '/images/projector.png', 4500000, '2022-09-25', 30, 12),
    ('Mouse Logitech', 2, '/images/mouse.png', 300000, '2023-06-18', 40, 10),
    ('Keyboard HP', 2, '/images/keyboard.png', 500000, '2023-07-22', 60, 10),
    ('Air Conditioner Panasonic', 6, '/images/ac.png', 5000000, '2022-12-10', 100, 5),
    ('Television Sony', 7, '/images/television.png', 8000000, '2024-02-14', 3, 12),
    ('Refrigerator LG', 6, '/images/refrigerator.png', 7000000, '2021-11-30', 150, 18),
    ('Microwave Samsung', 8, '/images/microwave.png', 1800000, '2023-04-10', 20, 1),
    ('Blender Philips', 8, '/images/blender.png', 400000, '2023-05-01', 50, 15),
    ('Fan Sharp', 9, '/images/fan.png', 600000, '2024-01-30', 10, 17),
    ('Washing Machine Toshiba', 6, '/images/washing_machine.png', 4000000, '2023-07-15', 80, 10),
    ('Speakers JBL', 3, '/images/speakers.png', 1000000, '2023-03-20', 25, 15),
    ('Laptop HP', 2, '/images/laptop_hp.png', 9500000, '2024-02-01', 15, 18),
    ('Microwave Sharp', 8, '/images/microwave_sharp.png', 1500000, '2022-05-12', 40, 12),
    ('Projector BenQ', 1, '/images/projector_benq.png', 6000000, '2023-11-22', 50, 12),
    ('Oven Panasonic', 7, '/images/oven.png', 3000000, '2023-12-05', 10, 12);
   
