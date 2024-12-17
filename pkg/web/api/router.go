package api

import (
	"net/http"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/formation"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	formationHandler := formation.NewHandler()

	mux.HandleFunc("GET /descriptions/{formationString}", formationHandler.GetFormation)

	return mux
}
