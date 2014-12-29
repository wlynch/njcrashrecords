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
	log.Printf("Accident: %#v", accident)
	return &accident
}

func parseDriver(entry string) *njcrash.Driver {
	driver := njcrash.Driver{
		DepartmentCaseId: parseString(entry[8:31]),
		VehicleNumber:    parseInt(entry[32:34]),
		Address: njcrash.Address{
			City:  parseString(entry[35:60]),
			State: parseString(entry[61:63]),
			Zip:   parseString(entry[64:69]),
		},
		DOB:                parseTime(parseString(entry[73:83]), "0000"),
		Sex:                parseString(entry[84:85]),
		AlcoholTestGiven:   parseBool(entry[86:87]),
		AlcoholTestType:    parseString(entry[88:90]),
		AlcoholTestResults: parseString(entry[91:94]),
		Charge:             parseString(entry[95:125]),
		Summons:            parseString(entry[126:156]),
		MultiCharge:        parseBool(entry[157:158]),
		PhysicalStatus:     parseString(entry[159:161]),
	}
	log.Printf("Driver: %#v", driver)
	return &driver
}

func parseOccupant(entry string) *njcrash.Occupant {
	occupant := njcrash.Occupant{
		DepartmentCaseId:         parseString(entry[8:31]),
		VehicleNumber:            parseInt(entry[32:34]),
		OccupantNumber:           parseInt(entry[35:37]),
		PhysicalCondition:        parseString(entry[38:40]),
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
	log.Printf("Occupant: %#v", occupant)
	return &occupant
}

func parsePedestrian(entry string) *njcrash.Pedestrian {
	pedestrian := njcrash.Pedestrian{
		DepartmentCaseId:  parseString(entry[8:31]),
		PedestrianNumber:  parseInt(entry[32:34]),
		PhysicalCondition: parseString(entry[35:37]),
		Address: njcrash.Address{
			City:  parseString(entry[38:63]),
			State: parseString(entry[64:66]),
			Zip:   parseString(entry[67:72]),
		},
		DOB:                       parseTime(parseString(entry[73:83]), "0000"),
		Age:                       parseInt(entry[84:87]),
		Sex:                       parseString(entry[88:89]),
		AlcoholTestGiven:          parseBool(entry[90:91]),
		AlcoholTestType:           parseString(entry[92:94]),
		AlcoholTestResults:        parseString(entry[95:98]),
		Charge:                    parseString(entry[99:129]),
		Summons:                   parseString(entry[130:160]),
		MultiCharge:               parseBool(entry[161:162]),
		TrafficControls:           parseString(entry[163:165]),
		ContributingCircumstances: parseStringSequence(entry[166:168], entry[169:171]),
		DirectionOfTravel:         parseString(entry[172:174]),
		PreCrashAction:            parseString(entry[175:177]),
		InjuryLocation:            parseString(entry[178:180]),
		InjuryType:                parseString(entry[181:183]),
		RefusedMedicalAttention:   parseBool(entry[184:185]),
		SafetyEquipmentUsed:       parseString(entry[186:188]),
		HospitalCode:              parseString(entry[189:193]),
		PhysicalStatus:            parseString(entry[194:196]),
		Bicyclist:                 parseBool(entry[197:198]),
		Other:                     parseBool(entry[199:200]),
	}
	log.Printf("Pedestrian: %#v", pedestrian)
	return &pedestrian
}

func parseVehicle(entry string) *njcrash.Vehicle {
	vehicle := njcrash.Vehicle{
		DepartmentCaseId:        parseString(entry[8:31]),
		VehicleNumber:           parseInt(entry[32:34]),
		InsuranceCompanyCode:    parseString(entry[35:39]),
		OwnerState:              parseString(entry[40:42]),
		Make:                    parseString(entry[43:73]),
		Model:                   parseString(entry[74:94]),
		Color:                   parseString(entry[95:98]),
		Year:                    parseInt(entry[99:103]),
		LicensePlateState:       parseString(entry[104:106]),
		VehicleWeightRating:     parseString(entry[107:108]),
		Towed:                   parseBool(entry[109:110]),
		RemovedBy:               parseString(entry[111:113]),
		InitialImpactLocation:   parseString(entry[114:116]),
		PrincipalDamageLocation: parseString(entry[117:119]),
		TrafficControlsPresent:  parseString(entry[120:122]),
		Type:                      parseString(entry[123:125]),
		Use:                       parseString(entry[126:128]),
		SpecialFunction:           parseString(entry[129:131]),
		CargoBodyType:             parseString(entry[132:134]),
		ContributingCircumstances: parseStringSequence(entry[135:137], entry[138:140]),
		DirectionOfTravel:         parseString(entry[141:143]),
		PreCrashAction:            parseString(entry[144:146]),
		SequenceOfEvents: parseStringSequence(
			entry[147:149],
			entry[150:152],
			entry[153:155],
			entry[156:158]),
		OversizePermit:  parseString(entry[159:161]),
		HazMatStatus:    parseString(entry[162:163]),
		HazMatPlacard:   parseString(entry[164:174]),
		USDOT:           parseBool(entry[175:176]),
		USDOTId:         parseString(entry[177:187]),
		CarrierName:     parseString(entry[188:238]),
		HitAndRunDriver: parseBool(entry[239:240]),
	}
	log.Printf("Vehicle: #%v", vehicle)
	return &vehicle
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
			var data []byte
			var err error
			entry := scanner.Text()
			log.Println(entry)
			// Determine the type of entry by it's size.
			switch len(entry) {
			case 74: // Occupant
				data, err = json.Marshal(parseOccupant(entry))
			case 161: // Driver
				data, err = json.Marshal(parseDriver(entry))
			case 200: // Pedestrian
				data, err = json.Marshal(parsePedestrian(entry))
			case 240: // Vehicle
				data, err = json.Marshal(parseVehicle(entry))
			case 458: // Accident
				data, err = json.Marshal(parseAccident(entry))
			default:
				err = fmt.Errorf("Unknown data of size %d", len(entry))
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(string(data))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
