package njcrash

import (
	"time"
)

type Accident struct {
	Year                        int     `json:"year"`
	CountyCode                  string  `json:"county_code,omitempty"`
	MunicipalityCode            string  `json:"municipality_code,omitempty"`
	DepartmentCaseNumber        string  `json:"department_case_number,omitempty"`
	CountyName                  string  `json:"county_name,omitempty"`
	MunicipalityName            string  `json:"municipality_name,omitempty"`
	CrashDate                   string  `json:"crash_date,omitempty"`
	CrashDayOfWeek              string  `json:"crash_day_of_week,omitempty"`
	CrashTime                   string  `json:"crash_time,omitempty"`
	PoliceDeptCode              string  `json:"police_dept_code,omitempty"`
	PoliceDepartment            string  `json:"police_department,omitempty"`
	PoliceStation               string  `json:"police_station,omitempty"`
	TotalKilled                 int     `json:"total_killed"`
	TotalInjured                int     `json:"total_injured"`
	PedestriansKilled           int     `json:"pedestrains_killed"`
	PedestriansInjured          int     `json:"pedestrians_injured"`
	Severity                    string  `json:"severity,omitempty"`
	Intersection                string  `json:"intersection,omitempty"`
	AlcoholInvolved             bool    `json:"alcohol_involved"`
	HazMatInvolved              bool    `json:"haz_mat_involved"`
	CrashTypeCode               string  `json:"crash_type_code,omitempty"`
	TotalVehiclesInvolved       int     `json:"total_vehicles_involved"`
	CrashLocation               string  `json:"crash_location,omitempty"`
	LocationDirection           string  `json:"location_direction,omitempty"`
	Route                       string  `json:"route,omitempty"`
	RouteSuffix                 string  `json:"route_suffix,omitempty"`
	StandardRouteIdentifier     string  `json:"standard_route_identifier,omitempty"`
	MilePost                    string  `json:"mile_post,omitempty"`
	RoadSystem                  string  `json:"road_system,omitempty"`
	RoadCharacter               string  `json:"road_character,omitempty"`
	RoadSurfaceType             string  `json:"road_surface_type,omitempty"`
	SurfaceCondition            string  `json:"surface_condition,omitempty"`
	LightCondition              string  `json:"light_condition,omitempty"`
	EnvironmentalCondition      string  `json:"environmental_condition,omitempty"`
	RoadDividedBy               string  `json:"road_divided_by,omitempty"`
	TemporaryTrafficControlZone string  `json:"temporary_traffic_control_zone,omitempty"`
	DistanceToCrossStreet       string  `json:"distance_to_cross_street,omitempty"`
	UnitOfMeasurement           string  `json:"unit_of_measurement,omitempty"`
	DirectionFromCrossStreet    string  `json:"direction_from_cross_street,omitempty"`
	CrossStreetName             string  `json:"cross_street_name,omitempty"`
	IsRamp                      bool    `json:"is_ramp"`
	RampRouteName               string  `json:"ramp_route_name,omitempty"`
	RampRouteDirection          string  `json:"ramp_route_direction,omitempty"`
	PostedSpeed                 int     `json:"posted_speed,omitempty"`
	PostedSpeedCrossStreet      int     `json:"posted_speed_cross_street,omitempty"`
	Latitude                    float64 `json:"latitude,omitempty"`
	Longitude                   float64 `json:"longitude,omitempty"`
	CellPhoneInUse              bool    `json:"cell_phone_in_use"`
	OtherPropertyDamage         string  `json:"other_property_damage,omitempty"`
	ReportingBadgeNumber        string  `json:"reporting_badge_number,omitempty"`

	// Custom defined
	Time time.Time `json:"time,omitempty"`
}
