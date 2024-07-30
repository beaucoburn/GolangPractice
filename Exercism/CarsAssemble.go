// Package cars provides functions for calculating the cost of cars, but is not a real Go package.
package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (float64(successRate) / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var value int = int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
    return value
}

func CalculateCost(carsCount int) uint {
    block := uint((carsCount)/10)
    ind := uint((carsCount)%10)
    blockPrice := uint((block)*95000)
    indPrice := uint ((ind)*10000)
    totalValue := uint(blockPrice + indPrice)
    return totalValue
}
