package main

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/router"
)

func main() {
	e := echo.New()

	router.CreateRouter(e)

	e.Start(":8080")
}
