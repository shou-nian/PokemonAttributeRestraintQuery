package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon/AttributeRestraintQuery/app/config"
	"net/http"
	"slices"
	"strings"
)

func Query(ctx *gin.Context) {
	paramsStr := ctx.Query("attr")
	if paramsStr == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"msg": "missing query string.",
			},
		)
		return
	}

	attrs := strings.Split(paramsStr, ",")

	var res = make([]map[string][]string, len(attrs)-1)
	for _, attr := range attrs {
		if !slices.Contains(config.AllAttributes, attr) {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": "attr not found. please check again.",
				},
			)
			return
		}
		res = append(res, config.FuncMapping[attr]())
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg": res,
		},
	)
}
