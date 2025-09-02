// Package weather helps forecast the weather of Globinocus.
package weather

// CurrentCondition stores weather condition.
var CurrentCondition string

// CurrentLocation stores city name.
var CurrentLocation string

// Forecast returns a text message that specifies a location and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
