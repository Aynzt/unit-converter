package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"test.io/conversion/length"
	"test.io/conversion/temperature"
	"test.io/conversion/weight"
)

var l = flag.Float64("l", 0., "Allow length conversions")
var t = flag.Float64("t", 0., "Allow temperature conversions")
var w = flag.Float64("w", 0., "Allow weight conversions")
var in = flag.Bool("i", false, "Interactive mode")

func main() {

	flag.Parse()

	if *in {
		fromStandardInput()
	}

	if t != nil {
		convertTemperature(*t)
	}

	if l != nil {
		convertLength(*l)
	}

	if w != nil {
		convertWeight(*w)
	}
}

// print conversion in celsius and fahrenheit
func convertTemperature(val float64) {
	fmt.Printf("%v = %v, %v = %v \n", temperature.Celsuis(val).String(), temperature.CtoF(temperature.Celsuis(val)).String(), temperature.Fahrenheit(val).String(), temperature.FtoC(temperature.Fahrenheit(val)).String())
}

// print conversion in Meter and Feet
func convertLength(val float64) {
	fmt.Printf("%v = %v, %v = %v \n", length.Feet(val).String(), length.FeetToMeter(length.Feet(val)).String(), length.Meter(val).String(), length.MeterToFeet(length.Meter(val)).String())
}

// print conversion in Kilogram and Pound
func convertWeight(val float64) {
	fmt.Printf("%v = %v, %v = %v \n", weight.Kilogram(val).String(), weight.KilogramtoP(weight.Kilogram(val)).String(), weight.Pound(val).String(), weight.PoundtoKg(weight.Pound(val)).String())
}

// Retrieve number from standard input to convert it through chosed conversion method
func fromStandardInput() {
setInput:
	fmt.Println("Insert a number :")

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {

		fmt.Println("Select a conversion : \n 1- Temperature \n 2- Length \n 3- Weight \n 0 - exit \n return - Retry")

		choice := bufio.NewScanner(os.Stdin)

		val, err := strconv.ParseFloat(input.Text(), 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while parsing to float : %v", err)
		}

		for choice.Scan() {
			switch {
			default:
				goto setInput
			case choice.Text() == "1":
				convertTemperature(val)
			case choice.Text() == "2":
				convertLength(val)
			case choice.Text() == "3":
				convertWeight(val)
			case choice.Text() == "0":
				os.Exit(0)
			}
		}

	}
}
