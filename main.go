package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	gin.SetMode(gin.ReleaseMode)

	g.Use(MyMiddleware)

	g.GET("/route", MyHandlerGET)

	g.POST("/route", MyHandlerPOST)

	g.PUT("/route")

	http.ListenAndServe(":8090", g)
}

func MyMiddleware(c *gin.Context) {
	fmt.Print("My middleware")
}

func MyHandlerGET(c *gin.Context) {
	fmt.Println("omad zapros")

}

func MyHandlerPOST(c *gin.Context) {
	fmt.Println("omad zaprosi post")

}

func MyHandlerPUT(c *gin.Context) {
	fmt.Println("omad zaprosi post")

}
