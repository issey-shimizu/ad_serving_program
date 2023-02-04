
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE conversion (
id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
adverrtise_id int NOT NULL,
user_code varchar(255) NOT NULL,
conversion int NOT NULL,
created_at datetime	NOT NULL,
updated_at datetime	NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE conversion;