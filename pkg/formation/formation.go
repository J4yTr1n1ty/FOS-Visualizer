package formation

import (
	"net/http"
)

type FormationHandler struct{}

func NewHandler() *FormationHandler {
	return &FormationHandler{}
}

func (f *FormationHandler) GetFormation(w http.ResponseWriter, r *http.Request) {
	formationString := r.PathValue("formationString")

	w.Write([]byte(formationString))
}

func GetFormationImage(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
