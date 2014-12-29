package njcrash

type Vehicle struct {
	DepartmentCaseNumber      string
	VehicleNumber             string
	InsuranceCompanyCode      string
	OwnerState                string
	Make                      string
	Model                     string
	Color                     string
	Year                      int
	LicensePlateState         string
	VehicleWeightRating       string
	Towed                     string
	RemovedBy                 string
	InitialImpactLocation     string
	PrincipalDamageLocation   string
	TrafficControlsPresent    string
	VehicleType               string
	VehicleUse                string
	SpecialFunctionVehicles   string
	CargoBodyType             string
	ContributingCircumstances []string
	DirectionOfTravel         string
	PreCrashAction            string
	SequenceOfEvents          []string
	OversizePermit            string
	HazMatStatus              string
	HazMatPlacard             string
	USDOTFlag                 string
	USDOTNumber               string
	CarrierName               string
	HitAndRunDriver           bool
}
