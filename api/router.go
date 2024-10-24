package api

import (
	"fmt"
	"mathTube/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() http.Handler {
	g := gin.New()

	gin.SetMode(gin.ReleaseMode)
	//g.Use(MyMiddleware)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.POST("/authentication", AuthHandler)
	g.GET("/materials", GetAllMaterialsHandler)
	g.POST("/material", PostMaterialHandler)
	g.PUT("/material", PutMaterialHandler)
	g.DELETE("/material/:id", DeleteMaterialHandler)
	return g
}

// DeleteMaterialHandler godoc
//
//	@summary      Material
//	@description  delete material
//	@tags         Material
//	@accept       json
//	@produce      json
//
//	@param id path int true "material id"
//
//	@success      200  {object}  Error
//	@failure      400  {object}  Error
//	@failure      500  {object}  Error
//	@Router       /material/{id} [delete]
func DeleteMaterialHandler(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println(err.Error())
		BadRequest(c, "Incorrect id")
		return
	}
	m := models.Material{ID: id}
	//m.Votes++
	if err := m.Delete(); err != nil {
		fmt.Println(err.Error())
		InternalError(c, "Not deleted")
		return
	}
	Success(c, "Success")
}

// PutMaterialHandler godoc
//
//	@summary      Material
//	@description  update material
//	@tags         Material
//	@accept       json
//	@produce      json
//
//	@param request body models.Material true "material info"
//
//	@success      200  {object}  Error
//	@failure      400  {object}  Error
//	@failure      500  {object}  Error
//	@Router       /material [put]
func PutMaterialHandler(c *gin.Context) {
	m := models.Material{}
	if err := c.ShouldBind(&m); err != nil {
		fmt.Println(err.Error())
		BadRequest(c, "Incorrect json struct")
		return
	}
	//m.Votes++
	if err := m.Update(); err != nil {
		fmt.Println(err.Error())
		InternalError(c, "Not updated")
		return
	}
	Success(c, "Success")
}

// PostMaterialHandler godoc
//
//	@summary      Material
//	@description  create material
//	@tags         Material
//	@accept       json
//	@produce      json
//
//	@param request body models.Material true "material info"
//
//	@success      200  {object}  Error
//	@failure      400  {object}  Error
//	@failure      500  {object}  Error
//	@Router       /material [post]
func PostMaterialHandler(c *gin.Context) {
	m := models.Material{}
	if err := c.ShouldBind(&m); err != nil {
		fmt.Println(err.Error())
		BadRequest(c, "Incorrect json struct")
		return
	}
	if err := m.Create(); err != nil {
		fmt.Println(err.Error())
		InternalError(c, "Not created")
		return
	}
	Success(c, "Success")
}

// GetAllMaterialsHandler materials godoc
//
//	@summary      Material
//	@description  materials
//	@tags         Material
//	@accept       json
//	@produce      json
//
//	@success      200  {object}  models.Materials
//	@failure      404  {object}  Error
//	@Router       /materials [get]
func GetAllMaterialsHandler(c *gin.Context) {
	ms := models.Materials{}
	if err := ms.ReadAll(); err != nil {
		fmt.Println(err.Error())
		NotFound(c, "Materials not found")
		return
	}
	c.JSON(http.StatusOK, ms)
}

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
