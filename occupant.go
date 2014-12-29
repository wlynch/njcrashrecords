package njcrash

type Occupant struct {
	DepartmentCaseId         string `json:"department_case_id,omitempty"`
	VehicleNumber            int    `json:"vehicle_number,omitempty"`
	OccupantNumber           int    `json:"occupant_number,omitempty"`
	PhysicalCondition        string `json:"physical_condition,omitempty"`
	PositionInVehicle        string `json:"position_in_vehicle,omitempty"`
	EjectionCode             string `json:"ejection_code,omitempty"`
	Age                      int    `json:"age,omitempty"`
	Sex                      string `json:"sex,omitempty"`
	InjuryLocation           string `json:"injury_location,omitempty"`
	InjuryType               string `json:"injury_type,omitempty"`
	RefusedMedicalAttention  bool   `json:"refused_medical_attention"`
	SafetyEquipmentAvailable string `json:"safety_equipment_available,omitempty"`
	SafetyEquipmentUsed      string `json:"safety_equipment_used,omitempty"`
	AirbagDeployment         string `json:"airbag_deployment,omitempty"`
	HospitalCode             string `json:"hospital_code,omitempty"`
}
