package njcrash

type Driver struct {
	Person
	VehicleNumber string `json:"vehicle_number,omitempty"`
	LicenseState  string `json:"driver_license_state,omitempty"`
}
