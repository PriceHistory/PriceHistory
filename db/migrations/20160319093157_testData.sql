
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
insert into product VALUES (1, 1, 'Test product');
insert into price VALUES (1, 1, now() - interval '4 month', 1000);
insert into price VALUES (2, 1, now() - interval '3 month', 1100);
insert into price VALUES (3, 1, now() - interval '2 month', 1200);
insert into price VALUES (4, 1, now() - interval '1 month', 1000);
insert into price VALUES (5, 1, now() - interval '3 week', 1500);
insert into price VALUES (6, 1, now() - interval '2 week', 1700);
insert into price VALUES (7, 1, now() - interval '1 week', 1100);
insert into price VALUES (8, 1, now() - interval '1 day', 1300);
insert into price VALUES (9, 1, now(), 1350);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
delete from price where productfk = 1;
delete from product where productpk = 1;

