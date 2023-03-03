#ad_serving_program
========================

## プログラム起動までの手順

1.以下コマンドを実行しネットワークを作成する。
docker network create golang_test_network

2.以下コマンドを実行しコンテナを起動する
###docker-compose up -d 

###3.以下コマンドを実行しgoコンテナにログインする
###docker compose exec go sh

###4.以下コマンドを実行し各テーブルの作成を行う
###goose up

###5.以下コマンドを実行しプログラムを起動する(/go/ad_serving_programディレクトリで実行する)
###go run main.go 

## Redisにデータを追加するまでの手順(advertise)
###1.以下コマンドを実行しコンテナにログインする
###docker exec -it ad_serving_program_redis_1 /bin/bash

###2.以下コマンドを実行しRedisに接続する
###redis-cli

###3.以下コマンドを実行しadvertiseにデータを追加する。created_atとupdated_atに関しては、任意の時間を入力する。
###set advertise id 1 name advertise1 image_url /img/img_1.jpeg redirect_url http://advertise1.s3-website-ap-northeast-1.amazonaws.com/index.html created_at “2023-02-04 00:11:50” updated_at “2023-02-04 00:11:50”


