package formation

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/parser"
)

type FormationHandler struct{}

func NewHandler() *FormationHandler {
	return &FormationHandler{}
}

type DescribeFormationRequest struct {
	FormationString string `json:"formation_string"`
}

func (f *FormationHandler) GetFormation(w http.ResponseWriter, r *http.Request) {
	requestBody := DescribeFormationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestBody.FormationString == "" {
		slog.Error("No formation string provided")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	parsedFormation := parser.GetEmptyTestingFormation(requestBody.FormationString)

	json, err := json.Marshal(parsedFormation)
	if err != nil {
		slog.Error("Failed to marshal JSON", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(json)
}

func GetFormationImage(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
