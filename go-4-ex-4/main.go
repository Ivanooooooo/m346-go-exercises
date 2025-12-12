package main

func convertCelsiusToFahrenheit(celsius int) int {
	fahrenheit := (celsius * 9 / 5) + 32
	return fahrenheit
}
func convertFahrenheitToCelsius(fahrenheit int) int {
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

func main() {
	fahrenheit := convertCelsiusToFahrenheit(30)
	celsius := convertFahrenheitToCelsius(90)
	println(fahrenheit, "F")
	println(celsius, "°C")

}
