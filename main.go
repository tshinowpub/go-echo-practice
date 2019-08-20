package main

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/dependencyinjection"
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/router"
)

func main() {
	e := echo.New()

	container, _ := dependencyinjection.New()

	router.NewRouter(e, container)

	e.Start(":8080")
}
