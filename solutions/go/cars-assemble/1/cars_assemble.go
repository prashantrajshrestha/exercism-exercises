package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var rateHour float64 = float64(productionRate) * (successRate / 100)

    return rateHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var rateMin = (float64(productionRate) * (successRate / 100)) / 60

    return int(rateMin)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groups int = carsCount / 10
    var individuals int = (carsCount - (groups) * 10) % 10

    var cost int = groups * 95000 + individuals * 10000

    return uint(cost)
 }
