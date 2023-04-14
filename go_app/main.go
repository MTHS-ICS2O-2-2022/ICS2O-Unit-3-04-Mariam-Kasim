// Created by: Mariam Kasim
// Created on: Apr 2023
//
//This program calculates the conversion of temperature from fahrenheit to celsius

package main

import (
	"fmt"
)

func main() {
	//This program calculates the conversion of temperature from fahrenheit to celsius
	var fahrenheit float64

	// input
	fmt.Println("This program calculates the conversion of temperature from fahrenheit to celsius.")
	fmt.Println()
	fmt.Print("Enter the fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// process
	celsius := (fahrenheit - 32) * 5 / 9
	roundCelsius := fmt.Sprintf("%.2f", celsius)

	// output
	fmt.Println()
	fmt.Println("The roundCelsius is: " + roundCelsius + " °C")
	fmt.Println("\nDone.")
}
