DROP TABLE products;
DROP TABLE categories;

CREATE TABLE IF NOT EXISTS categories(
    id          serial       PRIMARY KEY,
    name        varchar(255) NOT NULL,
    description text         NOT NULL,
    type        varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
    id          serial       PRIMARY KEY,
    category    int          NOT NULL REFERENCES categories(id),
    name        varchar(255) NOT NULL,
    price       float(32)    NOT NULL,
    description text         NOT NULL
);

INSERT INTO categories VALUES (DEFAULT, 'furnitures', 'stuff you put in your house', 'foo');
INSERT INTO products   VALUES (DEFAULT, 1, 'table', 120, 'thing you put things on');
INSERT INTO products   VALUES (DEFAULT, 1, 'chair', 40, 'thing you sit on');
INSERT INTO products   VALUES (DEFAULT, 1, 'stool', 39.99, 'thing you sit on but less nice');
INSERT INTO products   VALUES (DEFAULT, 1, 'desk',  20000, 'thing you put things on to work');

INSERT INTO categories VALUES (DEFAULT, 'food', 'stuff you put in your mouth', 'bar');
INSERT INTO products   VALUES (DEFAULT, 2, 'apple', 1, 'thing you eat thats red');
INSERT INTO products   VALUES (DEFAULT, 2, 'banana', 1.1, 'thing you eat thats yellow');

INSERT INTO categories VALUES (DEFAULT, 'clothes', 'stuff you put on yourself', 'baz');
INSERT INTO products   VALUES (DEFAULT, 3, 'shirt', 20, 'thing you put on to go to work');
INSERT INTO products   VALUES (DEFAULT, 3, 'pants', 30, 'thing you put on to go to work that you go to jail if you dont have');
