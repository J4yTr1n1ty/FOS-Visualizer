package lexer

import "github.com/J4yTr1n1ty/FOS-Visualizer/pkg/models"

func GetEmptyTestingFormation(formationString string) models.TrainFormationAtStop {
	trainFormation := models.NewTrainFormationAtStop(formationString)
	testSector := models.Sector{
		Name: "A",
	}

	testWagon := &models.Wagon{
		Status: "geschlossen",
		Type:   models.Engine,
		OrdNr:  0,
		Offers: []models.WagonOffer{models.FamilyZone},
	}

	testWagon2 := &models.Wagon{
		Status: "",
		Type:   models.SecondClassPassager,
		OrdNr:  0,
		Offers: []models.WagonOffer{models.Bicycle, models.LowFloorAccess},
	}

	testSector.Wagons = append(testSector.Wagons, *testWagon)
	testSector.Wagons = append(testSector.Wagons, *testWagon2)

	trainFormation.Sectors = append(trainFormation.Sectors, testSector)

	return *trainFormation
}
