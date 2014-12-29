package main

import (
	"njcrash"

	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func parseAccident(entry string) *njcrash.Accident {
	accident := njcrash.Accident{
		Year:                        parseInt(entry[0:4]),
		CountyCode:                  parseString(entry[4:6]),
		MunicipalityCode:            parseString(entry[6:8]),
		DepartmentCaseId:            parseString(entry[8:31]),
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
		Latitude:                    parseFloat(entry[352:360]),
		Longitude:                   parseFloat(entry[361:369]),
		CellPhoneInUse:              parseBool(entry[370:371]),
		OtherPropertyDamage:         parseString(entry[372:452]),
		ReportingBadgeId:            parseString(entry[453:458]),
	}
	accident.Time = parseTime(accident.CrashDate, accident.CrashTime)
	return &accident
}

func parseOccupant(entry string) *njcrash.Occupant {
	occupant := njcrash.Occupant{
		DepartmentCaseId:         parseString(entry[8:31]),
		VehicleNumber:            parseInt(entry[32:34]),
		OccupantNumber:           parseInt(entry[35:37]),
		PhysicalStatus:           parseString(entry[38:40]),
		PositionInVehicle:        parseString(entry[41:43]),
		EjectionCode:             parseString(entry[44:46]),
		Age:                      parseInt(entry[47:50]),
		Sex:                      parseString(entry[51:52]),
		InjuryLocation:           parseString(entry[53:55]),
		InjuryType:               parseString(entry[56:58]),
		RefusedMedicalAttention:  parseBool(entry[59:60]),
		SafetyEquipmentAvailable: parseString(entry[61:63]),
		SafetyEquipmentUsed:      parseString(entry[64:66]),
		AirbagDeployment:         parseString(entry[67:69]),
		HospitalCode:             parseString(entry[70:74]),
	}
	return &occupant
}

func main() {
	for _, path := range os.Args[1:] {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Read through the file line by line.
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			entry := scanner.Text()
			log.Println(entry)
			// Determine the type of entry by it's size.
			switch len(entry) {
			case 74: // Occupant
				occupant := parseOccupant(entry)
				log.Printf("Occupant: %#v", occupant)
				data, err := json.Marshal(occupant)
				if err != nil {
					log.Println(err)
				}
				fmt.Println(string(data))
			case 161, 240, 200:
				log.Println("Filetype unimplemented.")
			case 458: // Accident
				accident := parseAccident(entry)
				log.Printf("Accident: %#v", accident)
				log.Println(accident.Time.String())
				data, err := json.Marshal(accident)
				if err != nil {
					log.Println(err)
				}
				fmt.Println(string(data))
			default:
				log.Printf("Unknown data of size %d", len(entry))
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
