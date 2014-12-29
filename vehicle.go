package njcrash

type Vehicle struct {
	DepartmentCaseId          string   `json:"department_case_id,omitempty"`
	VehicleNumber             int      `json:"vehicle_id,omitempty"`
	InsuranceCompanyCode      string   `json:"insurance_company_code,omitempty"`
	OwnerState                string   `json:"owner_state,omitempty"`
	Make                      string   `json:"make,omitempty"`
	Model                     string   `json:"model,omitempty"`
	Color                     string   `json:"color,omitempty"`
	Year                      int      `json:"year,omitempty"`
	LicensePlateState         string   `json:"license_plate_state,omitempty"`
	VehicleWeightRating       string   `json:"vehicle_weight_rating,omitempty"`
	Towed                     bool     `json:"towed"`
	RemovedBy                 string   `json:"removed_by,omitempty"`
	InitialImpactLocation     string   `json:"initial_impact_location,omitempty"`
	PrincipalDamageLocation   string   `json:"principal_damage_location,omitempty"`
	TrafficControlsPresent    string   `json:"traffic_controls_present,omitempty"`
	Type                      string   `json:"type,omitempty"`
	Use                       string   `json:"use,omitempty"`
	SpecialFunction           string   `json:"special_function,omitempty"`
	CargoBodyType             string   `json:"cargo_body_type,omitempty"`
	ContributingCircumstances []string `json:"contributing_circumstances,omitempty"`
	DirectionOfTravel         string   `json:"direction_of_travel,omitempty"`
	PreCrashAction            string   `json:"precrash_action,omitempty"`
	SequenceOfEvents          []string `json:"sequence_of_events,omitempty"`
	OversizePermit            string   `json:"oversize_permit,omitempty"`
	HazMatStatus              string   `json:"hazmat_status,omitempty"`
	HazMatPlacard             string   `json:"mazmat_placard,omitempty"`
	USDOT                     bool     `json:"usdot"`
	USDOTId                   string   `json:"usdot_id,omitempty"`
	CarrierName               string   `json:"carrier_name,omitempty"`
	HitAndRunDriver           bool     `json:"hit_and_run_driver"`
}
