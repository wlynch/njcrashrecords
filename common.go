package njcrash

import (
	"time"
)

type Address struct {
	City  string
	State string
	Zip   string
}

type Person struct {
	DepartmentCaseNumber string    `json:"department_case_number,omitempty"`
	Address              Address   `json:"address,omitempty"`
	DOB                  time.Time `json:"dob,omitempty"`
	Sex                  string    `json:"sex,omitempty"`
	AlcoholTestGiven     bool      `json:"alcohol_test_given,omitempty"`
	AlcoholTestType      string    `json:"alcohol_test_type,omitempty"`
	AlcoholTestResults   string    `json:"alcohol_test_results,omitempty"`
	Charge               string    `json:"charge,omitempty"`
	Summons              string    `json:"summons,omitempty"`
	MultiCharge          bool      `json:"multi_charge,omitempty"`
	PhysicalStatus       string    `json:"physical_status,omitempty"`
}
