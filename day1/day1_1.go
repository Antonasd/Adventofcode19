package main

import (
	"adventofcode19/util"
	"fmt"
	"strconv"
	"strings"
)

type fuelCalc func(int) int

func main() {

	input, inputError := util.GetInput()
	if inputError != nil {
		fmt.Print(inputError.Error())
		return
	}

	var weights []int
	inputWeights := strings.Split(input, "\n")
	for _, weight := range inputWeights {
		if weight != "" {
			weightInterger, convError := strconv.Atoi(weight)
			if convError != nil {
				fmt.Println("Failed to convert weight to integer:")
				fmt.Print(convError.Error())
			}
			weights = append(weights, weightInterger)
		}
	}

	fmt.Println("The total amount of fuel required in part 1 is ", calculateTotalFuel(weights, calculateFuel))
	fmt.Println("The total amount of fuel required in part 2 is ", calculateTotalFuel(weights, calculateFuel2))
}

//Part 1
func calculateFuel(weight int) int {
	return weight/3 - 2
}

//Part 2
func calculateFuel2(weight int) int {
	fuel := weight/3 - 2
	if fuel > 0 {
		return fuel + calculateFuel2(fuel)
	}
	return 0
}

func calculateTotalFuel(moduelWeights []int, calc fuelCalc) int {
	var totalFuel int
	for _, weight := range moduelWeights {
		totalFuel += calc(weight)
	}

	return totalFuel
}
