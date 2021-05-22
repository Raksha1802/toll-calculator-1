Before we start, the gif is from the movie Hackers,1995 ( Frankly I had to google it!!)

# Toll fee calculator 1.0 in Go
A calculator for vehicle toll fees implemented in Golang.

# Background

Our city has decided to implement toll fees in order to reduce traffic congestion during rush hours. This is the current draft of requirements:

    * Fees will differ between 8 SEK and 18 SEK, depending on the time of day
    * Rush-hour traffic will render the highest fee
    * The maximum fee for one day is 60 SEK
    * A vehicle should only be charged once an hour
        In the case of multiple fees in the same hour period, the highest one applies.
    * Some vehicle types are fee-free
    * Weekends and holidays are fee-free

# Solution Delivered

A program/solution that calculates the total toll fee that needs to be paid by a vehicle(of a given type) for their trips.

+  ## Improvements made:

    * The solution calculates toll fee for trips ranging over different dates rather than one particular day. 
    * Can't just hardcode the fee-free vehicle type into the code, instead instace of any new vehicle types added needs to be created.
    * The solution doesn't need the holiday dates for a particular year to be hardcoded instead, it automaticaly and dynamically calculates the holiday dates for a year based on the timestamp of the trips for which the fee is calculated.
    * In the previous code, both the conditions of fee-free vehicle and holiday/weeend dates were checked before returning the fee of a trip, but in this solution, the vehicle type alone is checked as a foremost step,before getting into further calculations.
    * The list of trips(in form of timestamp) for which the total tollfee needs to be calculated need not be order sensitive, the solution itself will take care of sorting the order of trips before calculations are carried out.

+   ## Assumptions Made:

    * Have assumed the rush hour traffic and thereby the toll fee for those hours based on the information from [website](https://www.transportstyrelsen.se/en/road/road-tolls/Congestion-taxes-in-Stockholm-and-Goteborg/congestion-tax-in-gothenburg/hours-and-amounts-in-gothenburg/).
    * The trips are recorded in form of standard UTC dateTime format.
    * Vehicle type value for calculation is to be parsed as standard vehicle struct type(for Go implementation).
    * Have assumed vehicle types Emergency, Military, Diplomat, Foreign, Tractor to be toll free vehicles and all other categories of vehicles are termed as "Common"
    * The trips for which the toll fee are to be calculated are parsed as a list of datetime strings.

# How to run the program:

(As a prerequisite step, should have go installed, and download the code packages and run "go mod init")

The solution contains a "main.go" file, which acts a feeder of data to carry out toll fee calculation. This particular file can be modified as per use.

## File structure of the Implemented Solution

1. feeCal\ ==> toll fee calculations.
2. holiday\ ==> holiday dates/weekened dates calculation
3. tollfreecheck\ ==> check for toll free criterias
4. vehicles\ ==> fee-free vehicle types
5. main.go ==> program to verify the solution

# Further Additions:

* Carry out tests to verify the solution implemented.

* Handling errors.

* As part of future enhancement the solution implemented can be extended to a webapp/web service which takes in an input of a particular unique Id of a user, maps the userId to a database containing all the uncharged/uncalculated trips of the past for the Id, vehicle type mapped to the Id. And the solution implemented here can be called, wherein the above mentioned uncharged trips can be parsed into as a list along with the vehicle type to calculate the pending toll fee for a User.

