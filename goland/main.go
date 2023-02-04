package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"src/handler"
	"src/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const tmplPath = "src/template/"

var db *sqlx.DB
var e = createMux()

func main() {

	db = connectDB()
	repository.SetDB(db)

	e.GET("/impression/:id", handler.Impression)
	e.GET("/click/:id", handler.ShowCookie)
	e.GET("/conversion/:id", handler.Conversion)

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
