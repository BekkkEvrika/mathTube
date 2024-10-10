package api

import "github.com/gin-gonic/gin"

type Error struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	MessageText string `json:"messageText,omitempty"`
}

func Success(c *gin.Context, mess string) {
	e := Error{Code: "200", Message: mess}
	c.JSON(200, e)
}

func SuccessText(c *gin.Context, mess string, text string) {
	e := Error{Code: "200", Message: mess, MessageText: text}
	c.JSON(200, e)
}

func InternalError(c *gin.Context, mess string) {
	e := Error{Code: "500", Message: mess}
	c.JSON(500, e)
}

func InternalErrorText(c *gin.Context, mess string, text string) {
	e := Error{Code: "500", Message: mess, MessageText: text}
	c.JSON(500, e)
}

func NotFound(c *gin.Context, mess string) {
	e := Error{Code: "404", Message: mess}
	c.JSON(404, e)
}

func NotFoundText(c *gin.Context, mess string, text string) {
	e := Error{Code: "404", Message: mess, MessageText: text}
	c.JSON(404, e)
}

func MethodNotAllowed(c *gin.Context, mess string) {
	e := Error{Code: "405", Message: mess}
	c.JSON(405, e)
}

func MethodNotAllowedText(c *gin.Context, mess string, text string) {
	e := Error{Code: "405", Message: mess, MessageText: text}
	c.JSON(405, e)
}

func Created(c *gin.Context, mess string) {
	e := Error{Code: "201", Message: mess}
	c.JSON(201, e)
}

func BadRequest(c *gin.Context, mess string) {
	e := Error{Code: "400", Message: mess}
	c.JSON(400, e)
}

func BadRequestText(c *gin.Context, mess string, text string) {
	e := Error{Code: "400", Message: mess, MessageText: text}
	c.JSON(400, e)
}
