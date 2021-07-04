package parkinglot

import (
	constant "SquadStack/base/constants"
	"strconv"
	"strings"
)

// parking query executor
func PrakingQuery(query string) (output string) {
	queryInfo := strings.Split(query, " ")

	switch queryInfo[0] {
	case constant.CREATE_PARKING_LOT:
		if len(queryInfo) == 2 {
			val, _ := strconv.Atoi(queryInfo[1])
			output = CreateParkingLot(val)
		} else {
			output = "Invalid Query"
		}
		break
	case constant.PARK:
		if len(queryInfo) == 4 {
			val, _ := strconv.Atoi(queryInfo[3])
			output = ParkVehicle(queryInfo[1], val)
		} else {
			output = "Invalid Query"
		}
		break
	case constant.SLOT_NUMBER_FOR_DRIVER_AGE:
		if len(queryInfo) == 2 {
			val, _ := strconv.Atoi(queryInfo[1])
			output = SlotNumbersForDriverOfAge(val)
		} else {
			output = "Invalid Query"
		}
		break
	case constant.SLOT_NUMBER_FOR_CAR_WITH_NUMBER:
		if len(queryInfo) == 2 {
			output = SlotNumberForCarWithNumber(queryInfo[1])
		} else {
			output = "Invalid Query"
		}
		break
	case constant.LEAVE:
		if len(queryInfo) == 2 {
			val, _ := strconv.Atoi(queryInfo[1])
			output = VehicleLeave(val)
		} else {
			output = "Invalid Query"
		}
		break
	case constant.VEHICLE_REGI_NUM_FOR_DRIVER_AGE:
		if len(queryInfo) == 2 {
			val, _ := strconv.Atoi(queryInfo[1])
			output = VehicleRegistrationNumberForDriverOfAge(val)
		} else {
			output = "Invalid Query"
		}
		break
	default:
		output = "Invalid Query"
	}
	return
}
