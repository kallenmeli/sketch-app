package canvas

import (
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (c *Handler) Draw(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *Handler) GetAll(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *Handler) GetById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	_, err := c.service.GetByID(ctx)
	if err != nil {
		return err
	}

	return nil
}
