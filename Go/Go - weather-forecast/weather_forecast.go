// Package weather provides a forecast of the current weather based on current conition and current location.
package weather

// CurrentCondition is where the current weather condition is stored.
var CurrentCondition string

// CurrentLocation is where the current city is stored.
var CurrentLocation string

// Forecast returns a statement that tells you the current cities weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
