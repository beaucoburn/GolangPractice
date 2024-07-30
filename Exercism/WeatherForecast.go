//Package weather is an exercism exercise and not a real Go package.
//Package weather takes the CurrentLocation and CurrentCondition variables and gives a weather forecast.
package weather

//CurrentCondition takes a string and identifies the location of the user.
var CurrentCondition string
//CurrentLocation takes a string and identifies the weather conditions of the user.
var CurrentLocation string

//Forecast takes the CurrentLocation of the user and combines with the CurrentCondition of the user to return the value of the conditions in the user's location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
