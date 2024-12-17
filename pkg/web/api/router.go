package api

import (
	"net/http"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/web/formation"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	formationHandler := formation.NewHandler()

	mux.HandleFunc("GET /formation/describe", formationHandler.GetFormation)

	return mux
}
