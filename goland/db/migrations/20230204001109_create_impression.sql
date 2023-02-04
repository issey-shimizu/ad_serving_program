
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE impression (
id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
adverrtise_id int NOT NULL,
impression int NOT NULL,
created_at datetime	NOT NULL,
updated_at datetime	NOT NULL
);

insert into impression values (1,1,0,now(),now());
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE impression;
