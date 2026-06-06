
//Package weather provides information about weather forecast for a particular city based on current condition.
package weather


var (
    //CurrentCondition represents current condition.
	CurrentCondition string
    //CurrentLocation represents current location.
	CurrentLocation  string
)

//Forecast returns a string which is weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
