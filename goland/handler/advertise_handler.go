package handler

import (
	"log"
	"math/rand"
	"net/http"
	"src/model"
	"src/repository"
	"time"

	"github.com/labstack/echo/v4"
)

func Impression(c echo.Context) error {
	advertise, err := repository.Adverdisplay()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Message":   "広告1のページです",
		"Now":       time.Now(),
		"Advertise": advertise, // 記事データをテンプレートエンジンに渡す
		"URL":       advertise[0].Image_url,
	}

	return render(c, "advertise/advertise_1.html", data)
}

func showCookie(c echo.Context) {
	var advertise model.Advertise
	cookie, err := c.Cookie("click_id")

	if err != nil {
		log.Fatal("Cookie: ", err)
	}

	if cookie.Value == "" {
		var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		random := make([]rune, 52)
		for i := range random {
			random[i] = alphabet[rand.Intn(len(alphabet))]
		}
		log.Println(random)

		cookie := &http.Cookie{
			Name:  "click_id", // ここにcookieの名前を記述
			Value: "random",   // ここにcookieの値を記述
		}

		http.SetCookie(c.Response().Writer, cookie)
	}

	res, err := repository.ClickIdSet(&advertise)
	if err != nil {
		log.Println(err.Error())
	}

	//resのリダイレクト先を取得して、リダイえk
	http.Redirect(c.Response().Writer, c.Request().Response.Request, "/", 302)
}
