// Package weather provides a solution to report on the current weather conditions based on the location.
package weather

// CurrentCondition saves the last known weather condition.
var CurrentCondition string

// CurrentLocation saves the last known location request.
var CurrentLocation string

// Forecast returns the curren weather condition for the city in question.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
