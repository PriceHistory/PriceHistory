CREATE TABLE product (
	ProductPK SERIAL PRIMARY KEY,
	ProductOuterID VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE price (
	PricePK SERIAL PRIMARY KEY,
	ProductFK INT NOT NULL REFERENCES product(ProductPK),
	PriceDate TIMESTAMP NOT NULL,
	Price INT
);

CREATE TABLE link (
	LinkPK serial PRIMARY KEY,
	LinkHref VARCHAR(255) NOT NULL UNIQUE,
	LinkText VARCHAR(255) NOT NULL
)

ALTER TABLE product
    ADD COLUMN ProductTitle VARCHAR(255);

select * from price where ProductFK = 6153;