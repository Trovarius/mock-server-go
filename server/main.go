package main 

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Word")
	}

	e.GET("/", handler)
	e.GET("/teste?status=1", handler)
	e.Logger.Fatal(e.Start(":1323"))
}