package routers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
}

func NewRouter() *Router {
	return &Router{
		echo: echo.New(),
	}
}

func (router *Router) Get(path string, handler func(http.ResponseWriter, *http.Request) error) {
	router.echo.GET(path, func(c echo.Context) error {
		return handler(c.Response(), c.Request())
	})
}

func (router *Router) Post(path string, handler func(http.ResponseWriter, *http.Request) error) {
	router.echo.POST(path, func(c echo.Context) error {
		return handler(c.Response(), c.Request())
	})
}

func (router *Router) Run() {
	log.Fatal(router.echo.Start(":8080"))
}
