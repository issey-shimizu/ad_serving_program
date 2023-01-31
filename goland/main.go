package main

import (
	"log"
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
