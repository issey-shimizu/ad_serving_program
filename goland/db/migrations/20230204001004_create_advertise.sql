-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE advertise (
    id int NOT NULL NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    image_url varchar(255) NOT NULL,
    redirect_url varchar(255) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL
);

insert into advertise values (1,'広告1','/img/img_1.jpeg','http://advertise1.s3-website-ap-northeast-1.amazonaws.com',now(),now());
insert into advertise values (2,'広告2','/img/img_2.jpeg','https://www.yahoo.co.jp',now(),now());
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DROP TABLE articles;