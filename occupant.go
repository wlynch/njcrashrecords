package njcrash

type Occupant struct {
	Person
	VehicleNumber                  string
	OccupantNumber                 string
	PositionInVehicle              string
	EjectionCode                   string
	LocationOfMostSevereInjury     string
	TypeOfMostSeverePhysicalInjury string
	RefusedMedicalAttention        string
	SafetyEquipmentAvailable       string
	SafetyEquipmentUsed            string
	AirbagDeployment               string
	HospitalCode                   string
}
