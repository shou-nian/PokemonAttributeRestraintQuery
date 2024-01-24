package main

import (
	"github.com/pokemon/AttributeRestraintQuery/app/router"
	"log/slog"
)

func main() {
	sgi := router.New()

	err := sgi.Run()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
