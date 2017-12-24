package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	RegisterApi(e)
	//	e.Use(echomiddleware.ContextDB(db))
	e.Start(":5000")
}

func RegisterApi(e *echo.Echo) {
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "fruit-api")
	// })
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	v1 := e.Group("/v1")
	green := v1.Group("/green")
	green.POST("/pay", Paywx)
	green.POST("/query", Querywx)
	green.POST("/reverse", Reversewx)
	green.POST("/refund", Refundwx)
}

func Paywx(c echo.Context) error {
	return nil
}

func Querywx(c echo.Context) error {
	return nil
}

func Reversewx(c echo.Context) error {
	return nil
}

func Refundwx(c echo.Context) error {
	return nil
}
