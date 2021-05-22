package feeCal

import(
	"fmt"
	"time"
	"math"
	"sort"
	"../tollfreecheck"
	"../vehicle"
)

// function to convert date string to type time.Time
func toTime(s string) time.Time{
	layout := "2006-01-02 15:04:05 -0700 MST"
		t, err := time.Parse(layout, s)

		if err != nil {
			fmt.Println(err)
		}
	return t	
}

//Get TollFee for Single date-time event
func GetTollFee(date string) int{

	t := toTime(date)

	var hour = t.Hour()
	var minute = int(t.Minute())

	b := tollfreecheck.CheckTollFree(date) //Check if the date is a holiday or a weekend

	if b == true{
		fmt.Println("It is a weekend",date)
		return 0
	}

	//calculate/fetch tool rate according to the time
	switch(true){
	case (hour == 6):
		if minute <= 29{
			return 0
		}else{
			return 8
		}
	case (hour == 7):
		return 18
	case (hour ==8):
		return 13
	case (hour >=9 && hour <=14):
		return 8
	case (hour ==15):
		if minute <= 29{
			return 13
		}else{
			return 18
		}
	case (hour == 16):
		return 18
	case (hour == 17):
		return 13
	case (hour == 18):
		if minute <= 29{
			return 8
		}else{
			return 0
		}
	default:
		return 0						
	}
}

// function to fetch year,month and date from a datetime 
func getYMD(date time.Time) (int,int,int){
	year, month, day := date.Date()
	return year,int(month),day
}

// Calculate total toll fee to be paid by a vehicle of a particular type
// Calculates total fee ranging over multiple days and datetime entries
func GetTotalTollFee(ve vehicle.Vehicle, dates []string) int{

	sort.Strings(dates) // Sorts the toll entries based on their dates
	
	a:= vehicle.CheckFreeVehicleType(ve) //checks if the given vehicle type is excempt from paying toll
	
	if(a == true){
		return 0 // returns 0 for vehicles that are exempt from paying fees
	}

	var istart = dates[0] //start from the first datetime entry for which toll needs to be calculated
	var feeftrip = GetTollFee(dates[0]) //get the toll rate charged for the first trip(first datetime) entry

	var total = feeftrip 
	var fullTotal = 0 // total fee for a vehicle ranging over mulltiple trips across multiple dates(if applicable)
	var currfee int

	for i:= 1;i< len(dates);i++ { //range over the other trip entries from the list

		y1,m1,d1 := getYMD(toTime(dates[i]))
		y2,m2,d2 := getYMD(toTime(istart))

		if (y1 != y2 || m1 != m2 || d1 != d2){ //checks if the toll fee is calculated for the same day or the next/other date
			
			istart = dates[i]
			total = int(math.Min(float64(total),60))
			fullTotal = fullTotal + total
			total = 0
			
		}

		currfee = GetTollFee(dates[i])
		var tempfee = 0
		
		var dur = toTime(istart).Sub(toTime(dates[i])).Minutes()
		//check if the duration between trips is less than 1 hour since a vehicle can be charged only once in an hour
		if dur > 60{
			istart = dates[i]
			total = total + currfee

		}else{
			if (total >=0){

				total = total - tempfee				
				tempfee = int(math.Max(float64(tempfee), float64(currfee)))	//to charge the maximum fee of a trip within an hour			
				total = total + tempfee
			}			
		}
	}

	total = int(math.Min(float64(total),60)) // total for all trips in a single day(cant be more than 60SEK)
	fullTotal = fullTotal + total //Total of all the trips of a vehicle ranging over all the trips of multiple dates and days

	return fullTotal
}
