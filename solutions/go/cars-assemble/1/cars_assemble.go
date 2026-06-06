package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var final float64 = float64(productionRate)*successRate/100.0
    return final
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var perHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
    var final float64 = perHour/60.0
    return int(final)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var bunch uint = uint(carsCount/10)
    var rem uint = uint(carsCount%10)
    var cost uint = bunch*95000 + rem*10000
    return cost
}
