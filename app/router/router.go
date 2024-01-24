package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon/AttributeRestraintQuery/app/config"
	"github.com/pokemon/AttributeRestraintQuery/app/controller"
	"net/http"
)

func New() *gin.Engine {
	router := gin.Default()

	// load HTML files
	router.LoadHTMLGlob("./app/templates/*")

	router.GET("/", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"index.html",
			gin.H{},
		)
	})

	router.GET("/help", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"msg": config.AllAttributes,
			},
		)
	})

	router.GET("/pokemon", controller.Query)

	return router
}
