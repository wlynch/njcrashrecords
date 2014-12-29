package njcrash

type Driver struct {
	Person
	VehicleId    string `json:"vehicle_id,omitempty"`
	LicenseState string `json:"license_state,omitempty"`
}
