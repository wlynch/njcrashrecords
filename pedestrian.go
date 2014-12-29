package njcrash

type Pedestrian struct {
	Person
	PedestrianNumber               string
	PhysicalCondition              string
	TrafficControls                string
	ContributingCircumstances      []string
	DirectionOfTravel              string
	PreCrashAction                 string
	LocationOfMostSevereInjury     string
	TypeOfMostSeverePhysicalInjury string
	RefusedMedicalAttention        bool
	SafetyEquipmentUsed            bool
	HospitalCode                   string
	Bicyclist                      bool
	Other                          bool
}
