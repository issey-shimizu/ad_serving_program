package repository

import (
	"log"
	"src/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Advertisedisplay(id int) ([]*model.Advertise, error) {

	//impressionテーブルのimpressionをカウントアップ＋更新日時を最新にして保存
	var impression []*model.Impression
	count_up := `UPDATE impression SET impression = impression + 1, updated_at = now() WHERE id = id;`
	if err := db.Select(&impression, count_up); err != nil {
		return nil, err
	}

	//advertiseテーブルの情報を取得
	var advertise []*model.Advertise
	advertise_reference := `SELECT * FROM advertise;`
	if err := db.Select(&advertise, advertise_reference); err != nil {
		return nil, err
	}
	log.Println(advertise)
	return advertise, nil

}
func ClickIdSet(click model.Click, id int, user_code string) ([]*model.Click, error) {

	now := time.Now()
	click.Created_at = now
	click.Updated_at = now
	click.Id = id
	click.User_code = user_code

	var click_table []*model.Click

	query := `insert into click (id, adverrtise_id, user_code,click,created_at,updated_at) values (:id,:id,:user_code,1,:created_at,:updated_at) on duplicate key update click = click + 1,updated_at = now();`

	tx := db.MustBegin()

	_, err := tx.NamedExec(query, click)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return nil, err
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	click_reference := `SELECT user_code FROM click;`
	if err := db.Select(&click_table, click_reference); err != nil {
		return nil, err
	}
	log.Println(click_table[0])
	return click_table, nil

}

func Conversion_count(click model.Click, conversion model.Conversion, id int, user_code string) error {

	//Clickテーブルのuser_codeに値が格納されているか確認する
	//上記の結果に応じて分岐。レコードが存在していればinsert

	now := time.Now()
	click.Created_at = now
	click.Updated_at = now
	click.Id = id
	click.User_code = user_code

	conversion.Created_at = now
	conversion.Updated_at = now
	conversion.Id = id
	conversion.User_code = user_code

	query1 := `select * from click where user_code = :user_code;`
	query2 := `insert into conversion (id, adverrtise_id, user_code,conversion,created_at,updated_at) values (:id,:id,:user_code,1,:created_at,:updated_at) on duplicate key update conversion = conversion + 1,updated_at = now();`

	tx := db.MustBegin()

	_, err := tx.NamedExec(query1, click)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return err
	} else {

		_, err := tx.NamedExec(query2, conversion)
		log.Println(err)
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	return nil

}
