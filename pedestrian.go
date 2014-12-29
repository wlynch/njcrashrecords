package njcrash

type Pedestrian struct {
	Person
	PedestrianId              string   `json"pedestrian_id,omitempty"`
	PhysicalCondition         string   `json"physical_condition,omitempty"`
	TrafficControls           string   `json"traffic_controls,omitempty"`
	ContributingCircumstances []string `json"contributing_circumstances,omitempty"`
	DirectionOfTravel         string   `json"direction_of_travel,omitempty"`
	PreCrashAction            string   `json"precrash_action,omitempty"`
	InjuryLocation            string   `json"injury_location,omitempty"`
	InjuryType                string   `json"injury_type,omitempty"`
	RefusedMedicalAttention   bool     `json"refused_medical_attention"`
	SafetyEquipmentUsed       bool     `json"safety_equipment_used"`
	HospitalCode              string   `json"hospital_code,omitempty"`
	Bicyclist                 bool     `json"bicyclist"`
	Other                     bool     `json"other"`
}
