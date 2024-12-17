package parser

import "github.com/J4yTr1n1ty/FOS-Visualizer/pkg/models"

func ParseFormation(formationString string) models.TrainFormationAtStop {
	trainFormation := models.NewTrainFormationAtStop(formationString)

	return *trainFormation
}
