package vehicle

import(
	"fmt"
)

// List of Toll Exempt Vehicle Types
var TollFreeVehicles = []string{
	"Motorbike",
	"Tractor",
	"Military",
	"Emergency",
	"Diplomat",
	"Foreign",
}

// Interface of Vehicle types
type Vehicle interface{
	VehicleType() string
}

// Instance of Motorbike type
type Motorbike struct{
	
}

func(mo Motorbike) VehicleType() string {
	return "Motorbike"
}

// Instance of Military type
type Military struct{

}

func(m Military) VehicleType() string{
	return "Military"
}

// Instance of Tractor type
type Tractor struct{

}

func(t Tractor) VehicleType() string{
	return "Tractor"
}

// Instance of Emergency Vehicle type
type Emergency struct{

}

func(e Emergency) VehicleType() string{
	return "Emergency"
}

// Instance of Diplomat Vehicle type
type Diplomat struct{

}

func(d Diplomat) VehicleType() string{
	return "Diplomat"
}

// Instance of Foreign Vehicles type
type Foreign struct{

}

func(f Foreign) VehicleType() string{
	return "Foreign"
}

// Instance for all other vehicle types clubbed as common
type Vehicles struct{

}

func(m Vehicles) VehicleType() string{
	return "Common"
}

// Check if a vehicle type for which the toll is calculated is exempt from paying toll fee
func contains(e string)bool{
	for _,a := range TollFreeVehicles{
		if a == e{
			fmt.Println("It is a toll free exceptional vehicle type")
			return true
		}
	}
	return false
}

// Fetches the instance of vehicle type 
func CheckFreeVehicleType(vehicle Vehicle)bool{
	veh := vehicle.VehicleType()
	res := contains(veh)
	// fmt.Println("vehicle Type", vehicle.VehicleType())
	return res
}
