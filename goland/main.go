package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"src/handler"
	"src/repository"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const tmplPath = "src/template/"

var db *sqlx.DB
var e = createMux()

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

	db = connectDB()
	repository.SetDB(db)

	e.GET("/impression/:id", handler.Impression)
	e.GET("/click/:id", handler.ShowCookie)
	e.GET("/conversion/:id", handler.Conversion)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	e.GET("/get", func(c echo.Context) error {
		key := "advertise"
		// Set
		/*
			for field, val := range map[string]string{"name": "shimizu", "age": "26"} {
				fmt.Println("Inserting", "field:", field, "val:", val)
				err := client.HSet(key, field, val).Err()
				if err != nil {
					fmt.Println("redis.Client.HSet Error:", err)
				}
			}
		*/

		// Get
		// HGet(key, field string) *StringCmd
		hgetVal, err := client.HGet(key, "image_url").Result()
		if err != nil {
			fmt.Println("redis.Client.HGet Error:", err)
		}
		fmt.Println(hgetVal)

		return nil
	})

	/*
		e.GET("/get", func(c echo.Context) error {
			fmt.Println("Go Redis Connection Test")

			ping, err := client.Ping().Result()
			fmt.Println(ping, err)

			key := c.QueryParam("key")
			val, err := client.Get(key).Result()

			if err != nil {
				return err
			}

			data := new(Request)
			data.Key = key
			data.Value = val
			log.Println("keyの値")
			log.Println(data.Key)
			log.Println("Valueの値")
			log.Println(data.Value)

			return c.JSON(http.StatusOK, data)
		})

		e.POST("/set", func(c echo.Context) error {
			key := c.FormValue("key")
			value := c.FormValue("value")
			err := client.Set(key, value, 0).Err()

			if err != nil {
				return err
			}

			data := new(Request)
			data.Key = key
			data.Value = value

			return c.JSON(http.StatusOK, data)
		})
	*/

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()
	e.Static("img", "img")
	e.Static("/js", "src/js")
	e.Static("/adv", "src/temlate/advertise")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://advertise1.s3-website-ap-northeast-1.amazonaws.com:80/"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.CORS())

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}
