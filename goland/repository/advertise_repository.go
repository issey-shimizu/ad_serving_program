package repository

import (
	"database/sql"
	"src/model"
	"time"
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

	return advertise, nil
}

func ClickIdSet(click *model.Click, id int, user_code string) (sql.Result, error) {

	now := time.Now()
	click.Created_at = now
	click.Updated_at = now
	click.Id = id
	click.User_code = user_code

	// DBにレコードがない場合はINSERT、レコードがある場合はUPDATEでclick数をカウントアップ
	query := `insert into click (id, adverrtise_id, user_code,click,created_at,updated_at) values (:id,:id,:user_code,1,:created_at,:updated_at) on duplicate key update click = click + 1,updated_at = now();`

	tx := db.MustBegin()
	res, err := tx.NamedExec(query, click)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return nil, err
	}

	// エラーがない場合はコミットします。
	tx.Commit()
	return res, nil

}
