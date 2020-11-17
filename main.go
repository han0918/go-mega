package main

import (
	"github.com/gorilla/context"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sample/controller"
	"sample/model"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
