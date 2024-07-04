CREATE TABLE public.product (
                                id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                                name VARCHAR(100) NOT NULL,
                                price integer NOT NULL,
                                count integer NOT NULL
);

INSERT INTO product (name, price, count) values ('банан', 30, 2);
INSERT INTO product (name, price, count) values ('яблоко', 20, 6);
INSERT INTO product (name, price, count) values ('мандарины', 45, 4);
INSERT INTO product (name, price, count) values ('виноград', 60, 1);