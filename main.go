package main

import (
	"log/slog"
	"net/http"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/boot"
	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/web/api"
	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/web/middleware"
)

func main() {
	slog.Info("Starting FOS-Visualizer")

	boot.LoadEnv()
	router := api.SetupRoutes()

	// Add middleware
	stack := middleware.CreateStack(middleware.Logging, middleware.Recovery)

	server := http.Server{
		Addr:    ":" + boot.Environment.GetEnv("PORT"),
		Handler: stack(router), // Wrap the router with the middleware stack created above
	}

	slog.Info("Listening on port :" + boot.Environment.GetEnv("PORT"))
	if err := server.ListenAndServe(); err != nil { // Blocking call to run the server
		slog.Error("Error starting server: " + err.Error())
	}
}
