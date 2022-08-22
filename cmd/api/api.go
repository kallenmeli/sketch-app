package api

import (
	"sketch/db"
	"sketch/internal/canvas"
	"sketch/internal/routers"
)

func Start() {
	router := routers.NewRouter()
	connection := db.GetConnection()
	repository := canvas.NewRepository(connection)
	drawer := canvas.NewDrawer()
	service := canvas.NewService(repository, drawer)
	handler := canvas.NewHandler(service)

	router.Post("/", handler.Draw)
	router.Get("/:id", handler.GetById)
	router.Run()
}
