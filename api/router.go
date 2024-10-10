package api

import (
	"fmt"
	"mathTube/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() http.Handler {
	g := gin.New()

	gin.SetMode(gin.ReleaseMode)
	//g.Use(MyMiddleware)
	g.POST("/authentication", AuthHandler)

	return g
}

/*
POST /authentication
{
	"login":6,
	"password":"1234"
}
	Person{login:admin,password:1234}
*/

func AuthHandler(c *gin.Context) {
	p := models.Person{}
	if err := c.ShouldBind(&p); err != nil {
		fmt.Println(err.Error())
		BadRequest(c, "Incorrect json struct")
		return
	}
	if err := p.Authentication(); err != nil {
		fmt.Println(err.Error())
		BadRequest(c, "Login or password incorrect")
		return
	}
	Success(c, "Success")
}
