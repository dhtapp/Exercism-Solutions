package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate / 100 * float64(productionRate)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount != 0 { // This avoids dividing by zero
		var groupOfTen int = carsCount / 10 // This creates the group of 10 cars
		if groupOfTen == 1 {                // If we end up with 1 group, we need calculate approriately
			return uint(groupOfTen) * 95000
		} else if groupOfTen > 1 { // This calculates the groups + remainder cost together
			var costForGroup int = groupOfTen * 95000
			var costforRemaining int = (carsCount - (groupOfTen * 10)) * 10000 // This line finds the reminder, I didn't use modulus b/c
			// I was having issues
			return uint(costForGroup) + uint(costforRemaining)
		} else { // This calculates if group is less than 0
			return uint(carsCount) * 10000
		}
	} else {
		return 0
	}
	panic("CalculateCost not implemented")
}
