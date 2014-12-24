package njcrash

import (
	"time"
)

type Driver struct {
	Year                 int       `json:"year,omitempty"`
	CountyCode           string    `json:"county_code,omitempty"`
	MunicipalityCode     string    `json:"municipality_code,omitempty"`
	DepartmentCaseNumber string    `json:"department_case_number,omitempty"`
	VehicleNumber        int       `json:"vehicle_number,omitempty"`
	DriverCity           string    `json:"driver_city,omitempty"`
	DriverState          string    `json:"driver_state,omitempty"`
	DriverZipCode        string    `json:"driver_zip_code,omitempty"`
	DriverLicenseState   string    `json:"driver_license_state,omitempty"`
	DriverDOB            time.Time `json:"driver_dob,omitempty"`
	DriverSex            string    `json:"driver_sex,omitempty"`
	AlcoholTestGiven     bool      `json:"alcohol_test_given,omitempty"`
	AlcoholTestType      string    `json:"alcohol_test_type,omitempty"`
	AlcoholTestResults   string    `json:"alcohol_test_results,omitempty"`
	Charge               string    `json:"charge,omitempty"`
	Summons              string    `json:"summons,omitempty"`
	MultiCharge          bool      `json:"multi_charge,omitempty"`
	DriverPhysicalStatus string    `json:"driver_physical_status,omitempty"`
}
