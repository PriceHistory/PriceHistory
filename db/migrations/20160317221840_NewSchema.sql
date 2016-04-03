
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE product (
	ProductPK SERIAL PRIMARY KEY,
	ProductOuterID VARCHAR(50) NOT NULL UNIQUE,
	ProductTitle VARCHAR(255)
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
);

CREATE TABLE linkprocess (
	LinkProcessPK SERIAL PRIMARY KEY,
	LinkFK INT NOT NULL REFERENCES link(LinkPK),
	Status INT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE link;
DROP TABLE price;
DROP TABLE product;
DROP TABLE linkprocess;
