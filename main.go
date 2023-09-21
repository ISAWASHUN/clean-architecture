package main

import (
	"clean-artchitecture/controller"
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	return db, err
}

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS task (id INTEGER PRIMARY KEY, title TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskController := controller.TaskController{}

	e.GET("/tasks", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}