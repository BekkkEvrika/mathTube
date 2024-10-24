package main

import (
	"fmt"
	"mathTube/api"
	"mathTube/base"
	"mathTube/models"
	_ "mathTube/docs"
	"net/http"
)

//	@title           MathTube API
//	@version         1.0
//	@description     mathtube for video materials
//	@contact.name    test
//	@contact.email  test
//	@license.name  Apache 2.0
//	@license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//	@host      localhost:8090
//	@BasePath	/

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

// CRUD-
