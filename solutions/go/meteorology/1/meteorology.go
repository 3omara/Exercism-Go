package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tmpU TemperatureUnit) String() string {
	signs := []string{
		"°C",
		"°F",
	}
	return signs[tmpU]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (tmp Temperature) String() string {
	return fmt.Sprintf("%v %v", tmp.degree, tmp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (spdU SpeedUnit) String() string {
	units := []string{
		"km/h",
		"mph",
	}
	return units[spdU]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (spd Speed) String() string {
	return fmt.Sprintf("%v %v", spd.magnitude, spd.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (mtrData MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		mtrData.location,
		mtrData.temperature,
		mtrData.windDirection,
		mtrData.windSpeed,
		mtrData.humidity)
}
