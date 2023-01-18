#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table article (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    title varchar(50) NOT NULL,
    body varchar(1000)
    );"
$CMD_MYSQL -e  "insert into article values (1, '記事1', '記事1です。');"
$CMD_MYSQL -e  "insert into article values (2, '記事2', '記事2です。');"



# advertiseテーブルの作成
CREATE TABLE advertise (
    id int NOT NULL NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL,
    image_url varchar(255) NOT NULL,
    redirect_url varchar(255) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL
);

# impressionテーブルの作成
CREATE TABLE impression (
id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
adverrtise_id int NOT NULL,
impression int NOT NULL,
created_at datetime	NOT NULL,
updated_at datetime	NOT NULL
);


# impressionテーブルの作成
CREATE TABLE click (
id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
adverrtise_id int NOT NULL,
user_code varchar(255) NOT NULL,
created_at datetime	NOT NULL,
updated_at datetime	NOT NULL
);



# advertiseテーブルに仮データの追加
insert into advertise values (1,'広告1','aaaa','aaaa','2019-10-04 15:25:07','2022-12-20 15:25:07');

# implessionテーブルに仮データの追加
insert into impression values (1,1,0,'2019-10-04 15:25:07','2022-12-20 15:25:07');