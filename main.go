package main

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/handler"
)

func main() {
	e := echo.New()

	e.GET("/users", handler.Run())

	e.Start(":8080")
}
