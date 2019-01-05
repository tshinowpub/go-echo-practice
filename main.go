package main

import (
	"github.com/labstack/echo"
	"github.com/tshinowpub/go-echo-practice/infrastructure/api/router"
	"github.com/tshinowpub/go-echo-practice/provider"
)

func main() {
	e := echo.New()

	provider := provider.Boot()

	router.NewRouter(e, provider)

	e.Start(":8080")
}
