package tollfreecheck

import(
	"fmt"
	"time"
	"../holiday"
	"strings"
)
var value holiday.Cal
var hol []string
var y int

// Checks if a particular date is a weekend or a holidy and is exempt from toll fee
func CheckTollFree(date string) bool{

	layout := "2006-01-02 15:04:05 -0700 MST"
	
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	year := t.Year()
	weekday := int(t.Weekday())

	//check if weekday is Saturday or Sunday
	if(weekday == 0 || weekday == 6){
		return true
	}

	if(year != y){
	hol = value.NonWeekendHol(year) // fetch the holiday calender for the year specified
	}

	d1 := t.String()

	for _,item := range(hol){
		if (strings.Contains(item,d1) == true){
			fmt.Println("It is a National Holiday")
			return true
		}
	}
	return false
}

// Fetches the holiday calender for the present year by default
func GetHol(){	
	y = time.Now().Year()
	hol = value.NonWeekendHol(y)
}

