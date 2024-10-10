package main

import (
	"fmt"
	"mathTube/api"
	"mathTube/base"
	"mathTube/models"
	"net/http"
)

func main() {

	if err := base.Connect(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Db connected...")

	if err := models.AutoMigrate(); err != nil {
		fmt.Println(err.Error())
		return
	}

	http.ListenAndServe(":8090", api.Router())
}

//1) Authentication - 1. Login, 2. Password
