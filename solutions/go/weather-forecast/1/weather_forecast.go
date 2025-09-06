// Package weather checks and returns weather conditions.
package weather

// CurrentCondition ... stores what current condition is.
var CurrentCondition string
// CurrentLocation ... stores what current Location is.
var CurrentLocation string

// Forecast tells what the forecast is.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
