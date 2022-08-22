package routing

import (
	goerrors "errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"sketch/internal/errors"
)

type Router struct {
	router *httprouter.Router
}

type ErrorResult struct {
	Message string `json:"message"`
}

type RouterParams map[string]string

func errorHandler(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	log.Error(err)
	var businessErr errors.Error
	if ok := goerrors.As(err, &businessErr); ok {
		_ = ToJSON(w, http.StatusBadRequest, ErrorResult{
			Message: err.Error(),
		})
		return
	}

	_ = ToJSON(w, http.StatusInternalServerError, ErrorResult{
		Message: "failed to process the request",
	})
}

func NewRouter() *Router {
	return &Router{
		router: httprouter.New(),
	}
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request, httprouter.Params) error) {
	r.router.GET(path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := handler(writer, request, params)
		errorHandler(writer, err)
	})
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request, httprouter.Params) error) {
	r.router.POST(path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := handler(writer, request, params)
		errorHandler(writer, err)
	})
}

func (r *Router) Run() {
	port := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r.router))
}
