//Package weather contains a function which can be used to get the current weather condition of a city.
package weather

//CurrentCondition is a variable that stores the current weather condition.
var CurrentCondition string

//CurrentLocation is a variable that stores the current location of the place of forecast.
var CurrentLocation string

//Forecast function while providing the city and the condition will provide an user readable output regarding the weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
