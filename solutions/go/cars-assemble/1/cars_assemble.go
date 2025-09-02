package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	res := float64(productionRate) * successRate / 100.0
	return res
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	res := int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
	return res
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	res := carsCount/10*95_000 + (carsCount%10)*10_000
	return uint(res)
}
