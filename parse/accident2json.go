package main

import (
	"njcrash"

	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseBool(s string) bool {
	return strings.ToUpper(strings.TrimSpace(s)) == "Y"
}

func parseInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Printf("parseInt: %v", err)
	}
	return i
}

func parseString(s string) string {
	return strings.TrimSpace(s)
}

func parseTime(d, t string) time.Time {
	val, err := time.Parse("01/02/2006 1504 MST", fmt.Sprintf("%s %s EST", d, t))
	if err != nil {
		log.Printf("parseTime: %v", err)
	}
	return val
}

func parseAccident(entry string) *njcrash.Accident {
	accident := njcrash.Accident{
		Year:                        parseInt(entry[0:4]),
		CountyName:                  parseString(entry[32:44]),
		MunicipalityName:            parseString(entry[45:69]),
		CrashDate:                   parseString(entry[70:80]),
		CrashDayOfWeek:              parseString(entry[81:83]),
		CrashTime:                   parseString(entry[84:88]),
		PoliceDeptCode:              parseString(entry[89:91]),
		PoliceDepartment:            parseString(entry[92:117]),
		PoliceStation:               parseString(entry[119:133]),
		TotalKilled:                 parseInt(entry[134:136]),
		TotalInjured:                parseInt(entry[137:139]),
		PedestriansKilled:           parseInt(entry[140:142]),
		PedestriansInjured:          parseInt(entry[143:145]),
		Severity:                    parseString(entry[146:147]),
		Intersection:                parseString(entry[148:149]),
		AlcoholInvolved:             parseBool(entry[150:151]),
		HazMatInvolved:              parseBool(entry[152:153]),
		CrashTypeCode:               parseString(entry[154:156]),
		TotalVehiclesInvolved:       parseInt(entry[157:159]),
		CrashLocation:               parseString(entry[160:210]),
		LocationDirection:           parseString(entry[211:212]),
		Route:                       parseString(entry[213:217]),
		RouteSuffix:                 parseString(entry[218:219]),
		StandardRouteIdentifier:     parseString(entry[220:236]),
		MilePost:                    parseString(entry[237:244]),
		RoadSystem:                  parseString(entry[245:247]),
		RoadCharacter:               parseString(entry[248:250]),
		RoadSurfaceType:             parseString(entry[251:253]),
		SurfaceCondition:            parseString(entry[254:256]),
		LightCondition:              parseString(entry[257:259]),
		EnvironmentalCondition:      parseString(entry[260:262]),
		RoadDividedBy:               parseString(entry[263:265]),
		TemporaryTrafficControlZone: parseString(entry[266:268]),
		DistanceToCrossStreet:       parseString(entry[269:273]),
		UnitOfMeasurement:           parseString(entry[274:276]),
		DirectionFromCrossStreet:    parseString(entry[277:278]),
		CrossStreetName:             parseString(entry[279:314]),
		IsRamp:                      parseBool(entry[315:316]),
		RampRouteName:               parseString(entry[317:342]),
		RampRouteDirection:          parseString(entry[343:345]),
		PostedSpeed:                 parseInt(entry[346:348]),
		PostedSpeedCrossStreet:      parseInt(entry[349:351]),
		Latitude:                    parseString(entry[352:360]),
		Longitude:                   parseString(entry[361:369]),
		CellPhoneInUse:              parseBool(entry[370:371]),
		OtherPropertyDamage:         parseString(entry[372:452]),
		ReportingBadgeNo:            parseString(entry[453:458]),
	}
	accident.Time = parseTime(accident.CrashDate, accident.CrashTime)
	return &accident
}

func main() {
	file, err := os.Open("accidents.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		fmt.Println(entry)
		accident := parseAccident(entry)
		fmt.Printf("Accident: %#v\n", accident)
		fmt.Println(accident.Time.String())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
