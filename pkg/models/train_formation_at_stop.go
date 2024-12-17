package models

type TrainFormationAtStop struct {
	FormationString string
	Sectors         []Sector
}

func NewTrainFormationAtStop(formationString string) *TrainFormationAtStop {
	return &TrainFormationAtStop{
		FormationString: formationString,
	}
}

type Sector struct {
	Name   string  `json:"name"`
	Wagons []Wagon `json:"wagons"`
}

type Wagon struct {
	Status string   `json:"status"`
	Type   string   `json:"type"`
	OrdNr  int      `json:"ord_nr"`
	Offers []string `json:"offers"`
}

func NewWagon() *Wagon {
	return &Wagon{}
}

func (w *Wagon) AddOffer(offer string) {
	w.Offers = append(w.Offers, offer)
}

func (w *Wagon) RemoveOffer(offer string) {
	for i, o := range w.Offers {
		if o == offer {
			w.Offers = append(w.Offers[:i], w.Offers[i+1:]...)
		}
	}
}

// Types
const (
	TypeFirstClassPassager           = "1"
	TypeSecondClassPassager          = "2"
	TypeFistAndSecondClassPassager   = "12"
	TypeCouchetteCoach               = "CC"
	TypeFamilyCoach                  = "FA"
	TypeSleepingCoach                = "WL"
	TypeRestaurant                   = "WR"
	TypeCombinedDiningAndFirstClass  = "W1"
	TypeCombinedDiningAndSecondClass = "W2"
	TypeEngine                       = "LK"
	TypeLuggageCoach                 = "D"
	TypeFicticiousCarriage           = "F" // On tracks with sectors, the delta between the length of the train and the stopping edge is filled with fictitious wagons at the front and/or rear.
	TypeClasslessVehicle             = "K"
	TypeParkedWagon                  = "X" // parked wagons influence the assignment of the vehicles of a train to the sectors, but are not part of the train in question.
)

// Offers
const (
	OfferWheelchairSpaces           = "BHP"
	OfferBusinessZone               = "BZ"
	OfferFamilyZone                 = "FZ"
	OfferBabyCarriagePlatform       = "KW"
	OfferLowFloorAccess             = "NF"
	OfferBicycle                    = "VH"
	OfferBicycleRequiresReservation = "VR"
)
