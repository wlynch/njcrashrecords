package njcrash

import (
	"time"
)

type Driver struct {
	DepartmentCaseId   string    `json:"department_case_id,omitempty"`
	VehicleNumber      int       `json:"vehicle_id,omitempty"`
	Address            Address   `json:"address,omitempty"`
	LicenseState       string    `json:"license_state,omitempty"`
	DOB                time.Time `json:"dob,omitempty"`
	Sex                string    `json:"sex,omitempty"`
	AlcoholTestGiven   bool      `json:"alcohol_test_given"`
	AlcoholTestType    string    `json:"alcohol_test_type,omitempty"`
	AlcoholTestResults string    `json:"alcohol_test_results,omitempty"`
	Charge             string    `json:"charge,omitempty"`
	Summons            string    `json:"summons,omitempty"`
	MultiCharge        bool      `json:"multi_charge"`
	PhysicalStatus     string    `json:"physical_status,omitempty"`
}
