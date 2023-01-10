package handler

import (
	"log"
	"net/http"
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
