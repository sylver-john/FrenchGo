package router

import (
  	"github.com/gin-gonic/gin"
	"../handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HandlerRequest",
		"GET",
		"/helloworld",
		handler.HandlerRequest,
	},
}