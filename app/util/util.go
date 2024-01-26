package util

import (
	"github.com/pokemon/AttributeRestraintQuery/app/model"
	"strconv"
)

func MergeResources(resource []model.ResultType) map[string]map[string][]string {
	res := make(map[string]map[string][]string, len(resource))

	for _, v := range resource {
		for k, vv := range v {
			var tempRes = make(map[string][]string, len(vv))
			for kk, vvv := range vv {
				tempRes[strconv.FormatFloat(float64(kk), 'f', 1, 64)] = vvv
			}
			res[k] = tempRes
		}
	}

	return res
}
