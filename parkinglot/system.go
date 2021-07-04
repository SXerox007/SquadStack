package parkinglot

import (
	"encoding/json"
	"strconv"
	"strings"
)

type ParkingSystem struct {
	SlotId          int    `json:"slotId" param:"slotId"`
	VehicleNumber   string `json:"vehicleNumber" param:"vehicleNumber"`
	DriverAge       int    `json:"driverAge" param:"driverAge"`
	IsSlotAvailable bool   `json:"isSlotAvailable" param:"isSlotAvailable"`
}

var pLot []ParkingSystem
var vehicleMap map[string]int

// create parking lot
func CreateParkingLot(size int) string {
	if len(pLot) != 0 {
		return "Parking Lot already creeated."
	}
	for i := 1; i <= size; i++ {
		pLot = append(pLot, ParkingSystem{SlotId: i, IsSlotAvailable: true})
	}
	vehicleMap = make(map[string]int)
	return "Created parking of " + strconv.Itoa(size) + " slots"
}

// park vehicle
func ParkVehicle(vehicleNumber string, driverAge int) string {
	if len(pLot) == 0 {
		return "No Parking Lot exists. Please create parking lot first."
	}
	if val, ok := vehicleMap[vehicleNumber]; ok {
		return "Car with vehicle registration number " + vehicleNumber + " is already parked at slot number " + strconv.Itoa(val)
	}
	lotNumber := -1
	for index, attr := range pLot {
		if attr.IsSlotAvailable {
			lotNumber = attr.SlotId
			pLot[index].DriverAge = driverAge
			pLot[index].VehicleNumber = vehicleNumber
			pLot[index].IsSlotAvailable = false
			break
		}
	}

	// add into vehicle map
	vehicleMap[vehicleNumber] = lotNumber
	return "Car with vehicle registration number " + vehicleNumber + " has been parked at slot number " + strconv.Itoa(lotNumber)
}

// slot number for driver of age
func SlotNumbersForDriverOfAge(driverAge int) string {
	if len(pLot) == 0 {
		return "No Parking Lot exists. Please create parking lot first."
	}
	arr := []int{}

	for _, attr := range pLot {
		if attr.DriverAge == driverAge {
			arr = append(arr, attr.SlotId)
		}
	}
	if len(arr) == 0 {
		return "No parked car matches the query"
	}
	s, _ := json.Marshal(arr)
	return strings.Trim(string(s), "[]")
}

// slot number for car with number
func SlotNumberForCarWithNumber(vehicleNumber string) string {
	if len(pLot) == 0 {
		return "No Parking Lot exists. Please create parking lot first."
	}

	if val, ok := vehicleMap[vehicleNumber]; ok {
		return "Car with vehicle registration number " + vehicleNumber + " has slot number " + strconv.Itoa(val)
	} else {
		return "No parked car matches the query"
	}
}

// vehicle leave
func VehicleLeave(slotNumber int) (result string) {
	if len(pLot) == 0 {
		return "No Parking Lot exists. Please create parking lot first."
	}

	for index, attr := range pLot {
		if attr.SlotId != slotNumber {
			continue
		}

		if attr.IsSlotAvailable {
			result = "There is no vehicle at slot " + strconv.Itoa(slotNumber)
			return
		}
		result = "Slot number " + strconv.Itoa(attr.SlotId) + " vacated, the car with vehicle registration number " + attr.VehicleNumber + " left the space, the driver of the car was of age " + strconv.Itoa(attr.DriverAge)
		pLot[index].DriverAge = -1
		pLot[index].VehicleNumber = ""
		pLot[index].IsSlotAvailable = true
		break
	}
	return
}

//
func VehicleRegistrationNumberForDriverOfAge(driverAge int) string {
	if len(pLot) == 0 {
		return "No Parking Lot exists. Please create parking lot first."
	}
	arr := []string{}

	for _, attr := range pLot {
		if attr.DriverAge == driverAge {
			arr = append(arr, attr.VehicleNumber)
		}
	}
	if len(arr) == 0 {
		return "No parked car matches the query"
	}
	s, _ := json.Marshal(arr)
	return strings.Trim(string(s), "[]")
}
