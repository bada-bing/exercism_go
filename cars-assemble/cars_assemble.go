package cars

// int is signed integer at least 32 bits, but usually 64 bits
//
// TypeConversion to convert to different numeric types
//
// in Go, you can't do arithmetic operations on different types of variables

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const costOfSingleCar uint = 10000
	const costOfTenCars uint = 95000

	return uint(carsCount/10)*costOfTenCars + uint(carsCount%10)*costOfSingleCar
}
