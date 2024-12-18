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
	Status WagonStatus  `json:"status"`
	Type   WagonType    `json:"type"`
	OrdNr  int          `json:"ord_nr"`
	Offers []WagonOffer `json:"offers"`
}

func NewWagon() *Wagon {
	return &Wagon{}
}

type WagonStatus string

const (
	Closed           WagonStatus = "-"
	GroupsAtThisStop WagonStatus = ">"
	Reserved         WagonStatus = "="
	OpenButNotServed WagonStatus = "%" // Only for Restaurant Wagons
)

func GetWagonStatusRegex() string {
	return `-|>|=|%`
}

type WagonType string

const (
	FirstClassPassager           WagonType = "1"
	SecondClassPassager          WagonType = "2"
	FistAndSecondClassPassager   WagonType = "12"
	CouchetteCoach               WagonType = "CC"
	FamilyCoach                  WagonType = "FA"
	SleepingCoach                WagonType = "WL"
	Restaurant                   WagonType = "WR"
	CombinedDiningAndFirstClass  WagonType = "W1"
	CombinedDiningAndSecondClass WagonType = "W2"
	Engine                       WagonType = "LK"
	LuggageCoach                 WagonType = "D"
	FicticiousCarriage           WagonType = "F" // On tracks with sectors, the delta between the length of the train and the stopping edge is filled with fictitious wagons at the front and/or rear.
	ClasslessVehicle             WagonType = "K"
	ParkedWagon                  WagonType = "X" // parked wagons influence the assignment of the vehicles of a train to the sectors, but are not part of the train in question.
)

func GetWagonTypesRegex() string {
	return `1|2|12|CC|FA|WL|WR|W1|W2|LK|D|F|K|X`
}

type WagonOffer string

const (
	WheelchairSpaces           WagonOffer = "BHP"
	BusinessZone               WagonOffer = "BZ"
	FamilyZone                 WagonOffer = "FZ"
	BabyCarriagePlatform       WagonOffer = "KW"
	LowFloorAccess             WagonOffer = "NF"
	Bicycle                    WagonOffer = "VH"
	BicycleRequiresReservation WagonOffer = "VR"
)

func GetWagonOfferRegex() string {
	return `BHP|BZ|FZ|KW|NF|VH|VR`
}
