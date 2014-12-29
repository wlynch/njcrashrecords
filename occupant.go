package njcrash

type Occupant struct {
	Person
	VehicleId                string `json:"vehicle_id,omitempty"`
	OccupantId               string `json:"occupant_id,omitempty"`
	PositionInVehicle        string `json:"position_in_vehicle,omitempty"`
	EjectionCode             string `json:"ejection_code,omitempty"`
	InjuryLocation           string `json:"injury_location,omitempty"`
	InjuryType               string `json:"injury_type,omitempty"`
	RefusedMedicalAttention  bool   `json:"refused_medical_attention"`
	SafetyEquipmentAvailable bool   `json:"safety_equipment_available"`
	SafetyEquipmentUsed      bool   `json:"safety_equipment_used"`
	AirbagDeployment         bool   `json:"airbag_deployment"`
	HospitalCode             string `json:",omitempty"`
}
