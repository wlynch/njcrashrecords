package main

import (
	"njcrash"

	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseAccident(entry string) *njcrash.Accident {
	accident := njcrash.Accident{
		CountyName:                  strings.TrimSpace(entry[32:44]),
		MunicipalityName:            strings.TrimSpace(entry[45:69]),
		CrashDate:                   strings.TrimSpace(entry[70:80]),
		CrashDayOfWeek:              strings.TrimSpace(entry[81:83]),
		CrashTime:                   strings.TrimSpace(entry[84:88]),
		PoliceDeptCode:              strings.TrimSpace(entry[89:91]),
		PoliceDepartment:            strings.TrimSpace(entry[92:117]),
		PoliceStation:               strings.TrimSpace(entry[119:133]),
		Severity:                    strings.TrimSpace(entry[146:147]),
		Intersection:                strings.TrimSpace(entry[148:149]),
		AlcoholInvolved:             strings.TrimSpace(entry[150:151]),
		HazMatInvolved:              strings.TrimSpace(entry[152:153]),
		CrashTypeCode:               strings.TrimSpace(entry[154:156]),
		TotalVehiclesInvolved:       strings.TrimSpace(entry[157:159]),
		CrashLocation:               strings.TrimSpace(entry[160:210]),
		LocationDirection:           strings.TrimSpace(entry[211:212]),
		Route:                       strings.TrimSpace(entry[213:217]),
		RouteSuffix:                 strings.TrimSpace(entry[218:219]),
		StandardRouteIdentifier:     strings.TrimSpace(entry[220:236]),
		MilePost:                    strings.TrimSpace(entry[237:244]),
		RoadSystem:                  strings.TrimSpace(entry[245:247]),
		RoadCharacter:               strings.TrimSpace(entry[248:250]),
		RoadSurfaceType:             strings.TrimSpace(entry[251:253]),
		SurfaceCondition:            strings.TrimSpace(entry[254:256]),
		LightCondition:              strings.TrimSpace(entry[257:259]),
		EnvironmentalCondition:      strings.TrimSpace(entry[260:262]),
		RoadDividedBy:               strings.TrimSpace(entry[263:265]),
		TemporaryTrafficControlZone: strings.TrimSpace(entry[266:268]),
		DistanceToCrossStreet:       strings.TrimSpace(entry[269:273]),
		UnitOfMeasurement:           strings.TrimSpace(entry[274:276]),
		DirectionFromCrossStreet:    strings.TrimSpace(entry[277:278]),
		CrossStreetName:             strings.TrimSpace(entry[279:314]),
		IsRamp:                      strings.TrimSpace(entry[315:316]),
		RampRouteName:               strings.TrimSpace(entry[317:342]),
		RampRouteDirection:          strings.TrimSpace(entry[343:345]),
		PostedSpeed:                 strings.TrimSpace(entry[346:348]),
		PostedSpeedCrossStreet:      strings.TrimSpace(entry[349:351]),
		Latitude:                    strings.TrimSpace(entry[352:360]),
		Longitude:                   strings.TrimSpace(entry[361:369]),
		CellPhoneInUseFlag:          strings.TrimSpace(entry[370:371]),
		OtherPropertyDamage:         strings.TrimSpace(entry[372:452]),
		ReportingBadgeNo:            strings.TrimSpace(entry[453:458]),
	}
	// Parse all int values
	if year, err := strconv.Atoi(strings.TrimSpace(entry[0:4])); err == nil {
		accident.Year = year
	} else {
		log.Printf("Year: %v", err)
	}

	if totalKilled, err := strconv.Atoi(strings.TrimSpace(entry[134:136])); err == nil {
		accident.TotalKilled = totalKilled
	} else {
		log.Printf("TotalKilled: %v", err)
	}

	if totalInjured, err := strconv.Atoi(strings.TrimSpace(entry[137:139])); err == nil {
		accident.TotalInjured = totalInjured
	} else {
		log.Printf("TotalInjured: %v", err)
	}

	if pedKilled, err := strconv.Atoi(strings.TrimSpace(entry[140:142])); err == nil {
		accident.PedestriansKilled = pedKilled
	} else {
		log.Printf("PedestriansKilled: %v", err)
	}

	if pedInjured, err := strconv.Atoi(strings.TrimSpace(entry[143:145])); err == nil {
		accident.PedestriansInjured = pedInjured
	} else {
		log.Printf("PedestriansInjured: %v", err)
	}

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
		fmt.Printf("Accident: %#v", accident)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
