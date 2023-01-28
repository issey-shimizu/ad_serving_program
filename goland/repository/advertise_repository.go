package repository

import (
	"database/sql"
	"src/model"
	"time"
)

// ArticleList ...
func Advertisedisplay(id int) ([]*model.Advertise, error) {

	//広告のimpressionをカウントアップして保存
	var impression []*model.Impression
	count_up := `UPDATE impression SET impression = impression + 1 WHERE id = id;`
	if err := db.Select(&impression, count_up); err != nil {
		return nil, err
	}

	//広告情報を参照する
	var advertise []*model.Advertise
	advertise_reference := `SELECT * FROM advertise;`
	if err := db.Select(&advertise, advertise_reference); err != nil {
		return nil, err
	}

	return advertise, nil
}

func ClickIdSet(click *model.Click, id int) (sql.Result, error) {
	//func ClickIdSet(click *model.Click, id int) ([]*model.Advertise, error) {

	now := time.Now()
	click.Created_at = now
	click.Updated_at = now
	click.Id = id
	// DBにレコードがない場合はINSERT、レコードがある場合はUPDATEでclick数をカウントアップ
	query := `insert into click (id, adverrtise_id, user_code,click,created_at,updated_at) values (:id,:id,"aa",1,:created_at,:updated_at) on duplicate key update click = click + 1,updated_at = now();`

	//query := `update click set click = click + 1 where id = 1;`
	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行します。
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付けされます。
	res, err := tx.NamedExec(query, click)
	//err := db.Select(&click, query)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return nil, err
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	//return res, nil

	/*広告情報を参照する
	var advertise []*model.Advertise
	advertise_reference := `SELECT * FROM advertise;`
	if err := db.Select(&advertise, advertise_reference); err != nil {
		return nil, err
	}

	return advertise, nil
	*/
	return res, nil

}
