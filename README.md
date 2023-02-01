#ad_serving_program
========================

## プログラム起動までの手順
1.以下コマンドを実行しコンテナを起動する
docker-compose up -d 

2.以下コマンドを実行しgoコンテナにログインする
docker compose exec

3.以下コマンドを実行し各テーブルの作成を行う
goose up

4.以下コマンドを実行しプログラムを起動する(/go/ad_serving_programディレクトリで実行する)
go run main.go 
