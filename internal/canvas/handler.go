package canvas

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sketch/internal/routers"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (c *Handler) Draw(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	requests, err := routers.FromJSON[DrawRequests](r)
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

	return routers.ToJSON(w, http.StatusOK, response)
}

func (c *Handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	id := params.ByName("id")
	canvas, err := c.service.GetByID(r.Context(), id)

	if errors.Is(err, ErrNotFound) {
		return routers.NotFound(w, err)
	}

	if err != nil {
		return err
	}

	return routers.ToJSON(w, http.StatusOK, canvas)
}
