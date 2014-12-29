package njcrash

import (
	"time"
)

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Zip   string `json:"zip,omitempty"`
}

type Person struct {
	DepartmentCaseId   string    `json:"department_case_id,omitempty"`
	Address            Address   `json:"address,omitempty"`
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
