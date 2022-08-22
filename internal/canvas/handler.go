package canvas

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"sketch/internal/routing"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (c *Handler) Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	id := r.URL.Query().Get("id")
	tmpl := template.Must(template.ParseFiles("./pages/home.html"))
	drawing, err := c.service.GetByID(r.Context(), id)

	if errors.Is(err, ErrNotFound) {
		return template.
			Must(template.ParseFiles("./pages/404.html")).
			Execute(w, nil)
	}

	if err != nil {
		return tmpl.Execute(w, err)
	}

	return tmpl.Execute(w, drawing)
}

func (c *Handler) Draw(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	requests, err := routing.FromJSON[DrawRequests](r)
	if err != nil {
		return fmt.Errorf("failed to get json body: %w", err)
	}

	if err := requests.Validate(); err != nil {
		return err
	}

	response, err := c.service.Save(r.Context(), requests)
	if err != nil {
		return err
	}

	return routing.ToJSON(w, http.StatusOK, response)
}

func (c *Handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	id := params.ByName("id")
	canvas, err := c.service.GetByID(r.Context(), id)

	if errors.Is(err, ErrNotFound) {
		return routing.NotFound(w, err)
	}

	if err != nil {
		return err
	}

	return routing.ToJSON(w, http.StatusOK, canvas)
}
