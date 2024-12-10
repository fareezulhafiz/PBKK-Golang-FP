CREATE TABLE IF NOT EXISTS categories(
    id          serial       PRIMARY KEY,
    name        varchar(255) NOT NULL,
    description text         NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
    id          serial       PRIMARY KEY,
    category    int          NOT NULL REFERENCES categories(id),
    name        varchar(255) NOT NULL,
    price       float(32)    NOT NULL,
    description text         NOT NULL
);

DELETE FROM products;
DELETE FROM categories;

INSERT INTO categories VALUES (1, 'furnitures', 'stuff you put in your house');
INSERT INTO products   VALUES (1, 1, 'table', 120, 'thing you put things on');
INSERT INTO products   VALUES (2, 1, 'chair', 40, 'thing you sit on');
INSERT INTO products   VALUES (3, 1, 'stool', 39.99, 'thing you sit on but less nice');
INSERT INTO products   VALUES (4, 1, 'desk',  20000, 'thing you put things on to work');

INSERT INTO categories VALUES (2, 'food', 'stuff you put in your mouth');
INSERT INTO products   VALUES (5, 2, 'apple', 1, 'thing you eat thats red');
INSERT INTO products   VALUES (6, 2, 'banana', 1.1, 'thing you eat thats yellow');

INSERT INTO categories VALUES (3, 'clothes', 'stuff you put on yourself');
INSERT INTO products   VALUES (7, 3, 'shirt', 20, 'thing you put on to go to work');
INSERT INTO products   VALUES (8, 3, 'pants', 30, 'thing you put on to go to work that you go to jail if you dont have');
