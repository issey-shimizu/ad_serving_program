package repository

import (
	"fmt"
	"log"
	"src/model"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func Count_up(id int) error {
	//impressionテーブルのimpressionをカウントアップ＋更新日時を最新にして保存
	var impression []*model.Impression
	count_up := `UPDATE impression SET impression = impression + 1, updated_at = now() WHERE id = id;`
	if err := db.Select(&impression, count_up); err != nil {
		return err
	}
	return nil
}

func Advertisedisplay(id int) ([]string, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	key := "advertise"

	//keyがadvetiseの情報を取得
	advertise, err := client.HVals(key).Result()
	if err != nil {
		fmt.Println("redis.Client.HGet Error:", err)
	}

	return advertise, nil
}

func ClickIdSet(click model.Click, id int, user_code string) ([]string, error) {

	now := time.Now()

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	key := "click"

	//keyがclickの情報が存在するか否かによって分岐
	val, err := client.HVals(key).Result()
	if len(val) == 0 {
		//新規で値を追加
		keys := [6]string{"id", "adverrtise_id", "user_code", "click", "created_at", "updated_at"}
		values := [6]string{strconv.Itoa(id), strconv.Itoa(id), user_code, "1", now.Format("2006-01-02 03:04:05"), now.Format("2006-01-02 03:04:05")}
		for i := 0; i < 6; i++ {
			client.HSet(key, keys[i], values[i])
		}

	} else if len(val) != 0 {
		//clickの値をカウントアップし、updated_atの値を最新にする
		click, err := client.HGet(key, "click").Result()
		if err != nil {
			log.Println(err)
		}

		click_count, _ := strconv.Atoi(click)

		client.HSet(key, "click", click_count+1)
		client.HSet(key, "updated_at", now.Format("2006-01-02 03:04:05"))

	} else {
		log.Println(err)
	}

	click_information, err := client.HVals(key).Result()
	if err != nil {
		fmt.Println("redis.Client.HVals Error:", err)
	}

	return click_information, err
}

func Conversion_count(conversion model.Conversion, id int, user_code string) error {

	now := time.Now()
	conversion.Created_at = now
	conversion.Updated_at = now
	conversion.Id = id
	conversion.User_code = user_code

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	key := "click"

	val, err := client.HGet(key, "user_code").Result()
	if err != nil {
		log.Println(err)
	}

	if len(val) == 0 {
		return nil
	} else {
		query := `insert into conversion (id, adverrtise_id, user_code,conversion,created_at,updated_at) values (:id,:id,:user_code,1,:created_at,:updated_at) on duplicate key update conversion = conversion + 1,updated_at = now();`

		tx := db.MustBegin()
		_, err := tx.NamedExec(query, conversion)
		if err != nil {
			log.Println(err)
		}

		tx.Commit()
	}

	if user_code == val {
		query := `insert into conversion (id, adverrtise_id, user_code,conversion,created_at,updated_at) values (:id,:id,:user_code,1,:created_at,:updated_at) on duplicate key update conversion = conversion + 1,updated_at = now();`

		tx := db.MustBegin()
		_, err := tx.NamedExec(query, conversion)
		if err != nil {
			log.Println(err)
		}
		tx.Commit()
	} else {
		return nil
	}
	return nil
}
