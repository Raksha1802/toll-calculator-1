package main

import(
	"fmt"
	"toll-calculator/tollfreecheck"
	"toll-calculator/feeCal"
	"toll-calculator/vehicle"
)

func createInstance(a string) vehicle.Vehicle{
	if(a == "motorbike"){
		return &vehicle.Motorbike{}
	}else if(a == "military"){
		return &vehicle.Military{}
	}
	return &vehicle.Vehicles{}
}


func main(){

	tollfreecheck.GetHol()
	//generate a list of dummy datetime trip events
	dates := []string{"2021-06-10 07:27:03 +0000 UTC","2021-06-08 17:02:03 +0000 UTC","2021-06-08 12:54:03 +0000 UTC", "2021-06-08 09:30:03 +0000 UTC","2021-06-08 10:32:03 +0000 UTC","2021-06-08 11:35:03 +0000 UTC","2021-06-08 13:54:03 +0000 UTC","2021-06-08 07:02:03 +0000 UTC"}	
	//calculate total toll fee for the vehicle specified
	totalfee := feeCal.GetTotalTollFee(createInstance("car"),dates)
	fmt.Println("Total Toll Fee that needs to be paid by the vehicle is =>",totalfee)
}