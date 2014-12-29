package njcrash

import (
	"time"
)

type Pedestrian struct {
	DepartmentCaseId          string    `json:"department_case_id,omitempty"`
	PedestrianNumber          int       `json:"pedestrian_number,omitempty"`
	PhysicalCondition         string    `json:"physical_condition,omitempty"`
	Address                   Address   `json:"address,omitempty"`
	DOB                       time.Time `json:"dob,omitempty"`
	Age                       int       `json:"age,omitempty"`
	Sex                       string    `json:"sex,omitempty"`
	TrafficControls           string    `json:"traffic_controls,omitempty"`
	AlcoholTestGiven          bool      `json:"alcohol_test_given"`
	AlcoholTestType           string    `json:"alcohol_test_type,omitempty"`
	AlcoholTestResults        string    `json:"alcohol_test_results,omitempty"`
	Charge                    string    `json:"charge,omitempty"`
	Summons                   string    `json:"summons,omitempty"`
	MultiCharge               bool      `json:"multi_charge"`
	ContributingCircumstances []string  `json:"contributing_circumstances,omitempty"`
	DirectionOfTravel         string    `json:"direction_of_travel,omitempty"`
	PreCrashAction            string    `json:"precrash_action,omitempty"`
	InjuryLocation            string    `json:"injury_location,omitempty"`
	InjuryType                string    `json:"injury_type,omitempty"`
	RefusedMedicalAttention   bool      `json:"refused_medical_attention"`
	SafetyEquipmentUsed       string    `json:"safety_equipment_used,omitempty"`
	HospitalCode              string    `json:"hospital_code,omitempty"`
	PhysicalStatus            string    `json:"physical_status,omitempty"`
	Bicyclist                 bool      `json:"bicyclist"`
	Other                     bool      `json:"other"`
}
