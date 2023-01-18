package repository

import (
	"database/sql"
	"log"
	"src/model"
)

// ArticleList ...
func Adverdisplay() ([]*model.Advertise, error) {
	query1 := `UPDATE impression SET impression = impression + 1 WHERE id='1';`
	var impression []*model.Impression
	if err := db.Select(&impression, query1); err != nil {
		return nil, err
	}

	query2 := `SELECT * FROM advertise;`

	var advertise []*model.Advertise
	if err := db.Select(&advertise, query2); err != nil {
		return nil, err
	}
	log.Println(advertise[0].ID)

	type Profile struct {
		Name string
		Age  int
	}

	p := []*Profile{
		{"Tanaka", 31},
		{"Suzuki", 46},
	}
	log.Println(p[0].Name)
	log.Println(p[0].Name)

	return advertise, nil
}

func ClickIdSet(click *model.Click) (sql.Result, error) {

	// IDが1の広告のimpressionをカウントアップするクエリ文字列を保存
	query := `insert into impression values (1,1,0,'2019-10-04 15:25:07','2023-01-18 15:25:07');`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行します。
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付けされます。
	res, err := tx.NamedExec(query, click)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return nil, err
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	return res, nil

}
