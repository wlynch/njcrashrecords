package njcrash

import (
	"time"
)

type Accident struct {
	Year                        int
	CountyCode                  int
	MunicipalityCode            int
	DepartmentCaseNumber        int
	CountyName                  string
	MunicipalityName            string
	CrashDate                   string
	CrashDayOfWeek              string
	CrashTime                   string
	PoliceDeptCode              string
	PoliceDepartment            string
	PoliceStation               string
	TotalKilled                 int
	TotalInjured                int
	PedestriansKilled           int
	PedestriansInjured          int
	Severity                    string
	Intersection                string
	AlcoholInvolved             bool
	HazMatInvolved              bool
	CrashTypeCode               string
	TotalVehiclesInvolved       int
	CrashLocation               string
	LocationDirection           string
	Route                       string
	RouteSuffix                 string
	StandardRouteIdentifier     string
	MilePost                    string
	RoadSystem                  string
	RoadCharacter               string
	RoadSurfaceType             string
	SurfaceCondition            string
	LightCondition              string
	EnvironmentalCondition      string
	RoadDividedBy               string
	TemporaryTrafficControlZone string
	DistanceToCrossStreet       string
	UnitOfMeasurement           string
	DirectionFromCrossStreet    string
	CrossStreetName             string
	IsRamp                      bool
	RampRouteName               string
	RampRouteDirection          string
	PostedSpeed                 int
	PostedSpeedCrossStreet      int
	Latitude                    string
	Longitude                   string
	CellPhoneInUse              bool
	OtherPropertyDamage         string
	ReportingBadgeNo            string

	// Custom defined
	Time time.Time
}
