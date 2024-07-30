// Package lasagna implements the logic for the lasagna exercise, and will not run as a normal Go package.
package lasagna

var OvenTime int = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	minutesToPrepare := 2
	return minutesToPrepare * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
