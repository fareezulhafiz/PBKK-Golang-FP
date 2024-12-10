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

INSERT INTO categories VALUES (1, 'furnitures', '');
INSERT INTO products   VALUES (1, 1, 'table', 12, '');
