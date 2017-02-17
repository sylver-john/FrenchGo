package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"./services"
)


func HandlerRequest(c *gin.Context) {
	services.doInsert()
  	c.JSON(http.StatusOK, "hello world!")
}
