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
		"Message":      "広告" + strconv.Itoa(id) + "のページです",
		"Now":          time.Now(),
		"image_url":    advertise[id-1].Image_url,
		"redirect_url": advertise[id-1].Redirect_url,
	}
	return render(c, "advertise/advertise_"+strconv.Itoa(id)+".html", data)
}

func ShowCookie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var click model.Click
	cookie, err := c.Cookie("click_id")

	click_Cookie := &http.Cookie{}

	if cookie == nil {

		var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		random := make([]rune, 52)
		for i := range random {
			random[i] = alphabet[rand.Intn(len(alphabet))]
		}
		click_Cookie.Name = "click_id"
		click_Cookie.Value = string(random)
		http.SetCookie(c.Response().Writer, click_Cookie)
	}

	user_code := cookie.Value

	click_table, err := repository.ClickIdSet(click, id, user_code)
	if err != nil {
		log.Println(err.Error())
	}

	advertise, err := repository.Advertisedisplay(id)
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	log.Println("click_tableの中身")
	log.Println(click_table[0].User_code)
	log.Println("advertiseの中身")
	log.Println(advertise[0].Redirect_url)

	c.Redirect(http.StatusFound, advertise[0].Redirect_url+"?advertise_id="+strconv.Itoa(id)+"&click_id="+click_table[0].User_code)
	return nil
}

func Conversion(c echo.Context) error {
	//v := c.Request().Response.Request.URL.Query()
	//log.Println(v)
	user_code := c.Request().URL.Query().Get("user_code")
	id, _ := strconv.Atoi(c.Param("id"))
	var click model.Click
	var conversion model.Conversion
	err := repository.Conversion_count(click, conversion, id, user_code)

	return err
}
