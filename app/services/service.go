package services

import (
	"github.com/pokemon/AttributeRestraintQuery/app/config"
	"github.com/pokemon/AttributeRestraintQuery/app/model"
	"github.com/pokemon/AttributeRestraintQuery/app/util"
)

func Main[F []config.FuncMap, R map[string]map[string][]string](f F) R {
	var source []model.ResultType

	for _, v := range f {
		source = append(source, v())
	}

	res := util.MergeResources(source)

	return res
}
