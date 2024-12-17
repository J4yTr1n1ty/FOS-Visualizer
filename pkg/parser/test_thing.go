package parser

import "github.com/J4yTr1n1ty/FOS-Visualizer/pkg/models"

func GetEmptyTestingFormation(formationString string) models.TrainFormationAtStop {
	trainFormation := models.NewTrainFormationAtStop(formationString)
	testSector := models.Sector{
		Name: "A",
	}

	testWagon := &models.Wagon{
		Status: "geschlossen",
		Type:   models.TypeEngine,
		OrdNr:  0,
		Offers: []string{models.OfferFamilyZone},
	}

	testWagon2 := &models.Wagon{
		Status: "geschlossen",
		Type:   models.TypeSecondClassPassager,
		OrdNr:  0,
		Offers: []string{models.OfferBicycle, models.OfferLowFloorAccess},
	}

	testSector.Wagons = append(testSector.Wagons, *testWagon)
	testSector.Wagons = append(testSector.Wagons, *testWagon2)

	trainFormation.Sectors = append(trainFormation.Sectors, testSector)

	return *trainFormation
}
