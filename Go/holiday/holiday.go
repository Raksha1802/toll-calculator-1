package holiday

import(
	// "fmt"
	"time"
)

type Cal struct {

}

//to calculae offset dates of holidays
func(c Cal) SetDate(sdate time.Time, offset int)  time.Time{
	// date := time.Date(sdate)
	// fmt.Println("date",sdate,offset)
	date1 := sdate.AddDate(0,0,offset)
	return date1
}

//calculate easter day for a given year
func(c Cal) EasterDate(year int) time.Time {
	
	if year < 1583 {
		year = 1583
	}

	month := 3
	// determine the Golden number
	golden := (year % 19) + 1
	// determine the century number
	century := year/100 + 1
	// correct for the years that are not leap years
	xx := (3*century)/4 - 12
	// moon correction
	yy := (8*century+5)/25 - 5
	// find Sunday
	zz := (5*year)/4 - xx - 10
	// determine epact
	// age of moon on January 1st of that year
	// (follows a cycle of 19 years)
	ee := (11*golden + 20 + yy - xx) % 30
	if ee == 24 {
		ee += 1
	}
	if (ee == 25) && (golden > 11) {
		ee += 1
	}
	// get the full moon
	moon := 44 - ee
	if moon < 21 {
		moon += 30
	}
	// up to Sunday
	day := (moon + 7) - ((zz + moon) % 7)
	// possibly up a month in easter_date
	if day > 31 {
		day -= 31
		month = 4
	}

	return time.Date(year,time.Month(month),int(day),0,0,0,0,time.UTC)
}

// Calculate Mid summer date holiday
func(c Cal) MidSummerDate(year int) time.Time {
	date := time.Date(year,time.Month(6),20,0,0,0,0,time.UTC)
	weekday := int(date.Weekday())
	return c.SetDate(date, 6 - weekday)
}

// Calculate All saints day holiday
func(c Cal) AllSaintsDay(year int) time.Time {
	date := time.Date(year,time.Month(10),31,0,0,0,0,time.UTC)
	weekday := int(date.Weekday())

	return c.SetDate(date, 6-weekday)
}

// Calculate non weekend holidays for a given year
func(c Cal) NonWeekendHol(year int) []string {

	easter := c.EasterDate(year)
	midsummer := c.MidSummerDate(year)
	allsaints := c.AllSaintsDay(year)

	return []string{
		time.Date(year,1,1,0,0,0,0,time.UTC).String(),//New Year's Day
		time.Date(year,1,6,0,0,0,0,time.UTC).String(),//Epiphany
		c.SetDate(easter, -2).String(),//Good Friday
		easter.String(),//Easter
		c.SetDate(easter,1).String(),//Easter Monday
		time.Date(year,5,1,0,0,0,0,time.UTC).String(),//Labors Day
		c.SetDate(easter,39).String(),//Ascesion Day
		c.SetDate(easter,49).String(),//Whitsun Day
		c.SetDate(easter,50).String(),
		time.Date(year,6,6,0,0,0,0,time.UTC).String(),//National Day of Sweden
		c.SetDate(midsummer, -1).String(), //MidSummer Eve
		midsummer.String(), //MidSummer Day
		allsaints.String(), //Allsaints Day
		time.Date(year,12,24,0,0,0,0,time.UTC).String(), //Christmas Eve
		time.Date(year,12,25,0,0,0,0,time.UTC).String(), //Christmas Day
		time.Date(year,12,26,0,0,0,0,time.UTC).String(), //St.Stephen's Day
		time.Date(year,12,31,0,0,0,0,time.UTC).String(), //New Years Eve
	}

}

