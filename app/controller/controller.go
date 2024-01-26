package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon/AttributeRestraintQuery/app/config"
	"github.com/pokemon/AttributeRestraintQuery/app/services"
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
	//
	//if !strings.Contains(paramsStr, ",") {
	//	if !slices.Contains(config.AllAttributes, paramsStr) {
	//		ctx.JSON(
	//			http.StatusBadRequest,
	//			gin.H{
	//				"msg": "attr not found. please check again.",
	//			},
	//		)
	//	} else {
	//		ctx.JSON(
	//			http.StatusOK,
	//			gin.H{
	//				"msg": config.FuncMapping[paramsStr](),
	//			},
	//		)
	//	}
	//	return
	//}

	attrs := strings.Split(paramsStr, ",")
	var res []map[string]map[string][]string
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
		var attributeFunctions []config.FuncMap
		attributeFunctions = append(attributeFunctions, config.FuncMapping[attr])

		res = append(res, services.Main(attributeFunctions))
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg": res,
		},
	)
}
