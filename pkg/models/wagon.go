package models

type Wagon struct {
	Status string
	Type   string
	OrdNr  int
	Offers []string
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
