package constant

//constants
const (
	ENV_LOCAL                       = "local"
	LOCAL_PORT                      = "50051"
	CREATE_PARKING_LOT              = "Create_parking_lot"
	PARK                            = "Park"
	SLOT_NUMBER_FOR_DRIVER_AGE      = "Slot_numbers_for_driver_of_age"
	SLOT_NUMBER_FOR_CAR_WITH_NUMBER = "Slot_number_for_car_with_number"
	LEAVE                           = "Leave"
	VEHICLE_REGI_NUM_FOR_DRIVER_AGE = "Vehicle_registration_number_for_driver_of_age"
	DEFAULT                         = "default"
)

/**
*
* Getting the environment local
*
**/
func GetLocal() string {
	return ENV_LOCAL
}

/**
*
* Getting the port local
*
**/
func GetPort() string {
	return LOCAL_PORT
}
