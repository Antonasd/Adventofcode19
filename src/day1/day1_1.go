package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type fuelCalc func(int) int

func main() {

	input, inputError := getInput()
	if inputError != nil {
		fmt.Print(inputError.Error())
		return
	}

	fmt.Println("The total amount of fuel required in part 1 is ", calculateTotalFuel(input, calculateFuel))
	fmt.Println("The total amount of fuel required in part 2 is ", calculateTotalFuel(input, calculateFuel2))
}

func getInput() ([]int, error) {
	var weights []int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please provide the path to your puzzle input: ")
	fmt.Print(">")

	path, inputError := reader.ReadString('\n')
	if inputError != nil {
		fmt.Println("Failed to read input path:")
		return weights, inputError
	}

	path = strings.TrimSuffix(path, "\n")
	path = strings.TrimSuffix(path, "\r")

	input, readError := ioutil.ReadFile(path)
	if readError != nil {
		fmt.Println("Failed to read input file:")
		return weights, readError
	}

	inputWeights := strings.Split(string(input), "\n")
	for _, weight := range inputWeights {
		if weight != "" {
			weightInterger, convError := strconv.Atoi(weight)
			if convError != nil {
				fmt.Println("Failed to convert weight to integer:")
				return weights, convError
			}
			weights = append(weights, weightInterger)
		}
	}

	return weights, nil
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
