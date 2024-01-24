package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pokemon/AttributeRestraintQuery/app/controller"
	"log/slog"
)

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		_, err := fmt.Fprintf(context.Writer, "hello welcome to pokemon query web site")
		if err != nil {
			slog.Error(err.Error())
			panic(err)
		}
	})
	router.GET("/pokemon", controller.Query)

	return router
}
