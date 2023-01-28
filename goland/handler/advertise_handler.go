package handler

import (
	"log"
	"math/rand"
	"net/http"
	"src/model"
	"src/repository"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Cookie struct {
	Name  string
	Value string
}

func Impression(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	advertise, err := repository.Advertisedisplay(id)
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Message": "広告" + strconv.Itoa(id) + "のページです",
		"Now":     time.Now(),
		//"Advertise":    advertise, // 広告データをテンプレートエンジンに渡す
		"image_url":    advertise[id-1].Image_url,
		"redirect_url": advertise[id-1].Redirect_url,
	}
	log.Println(advertise[id-1].Image_url)
	return render(c, "advertise/advertise_"+strconv.Itoa(id)+".html", data)
}

func ShowCookie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var click model.Click
	cookie, err := c.Cookie("click_id")

	if cookie == nil {

		click_Cookie := &http.Cookie{}

		var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		random := make([]rune, 52)
		for i := range random {
			random[i] = alphabet[rand.Intn(len(alphabet))]
		}
		click_Cookie.Name = "click_id"
		click_Cookie.Value = string(random)

		http.SetCookie(c.Response().Writer, click_Cookie)
	}

	/*
		advertise, err := repository.ClickIdSet(&click, id)
		if err != nil {
			log.Println(err.Error())
		}


		//resのリダイレクト先を取得して、リダイレクト
		http.Redirect(c.Response().Writer, c.Request(), advertise[id-1].Redirect_url, 200)
		return nil
	*/

	res, err := repository.ClickIdSet(&click, id)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(res)
	return nil

}
