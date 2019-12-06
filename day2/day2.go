package main

import (
	"adventofcode19/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please provide the path to your puzzle input: ")
	input, inputError := util.GetFileInput()
	if inputError != nil {
		fmt.Println(inputError.Error())
		return
	}

	var program []int
	sProgram := strings.Split(input, "\n")
	sProgram = strings.Split(sProgram[0], ",")
	for _, sInteger := range sProgram {
		if sInteger != "" {
			integer, convError := strconv.Atoi(sInteger)
			if convError != nil {
				fmt.Println("Failed to convert to integer:")
				fmt.Print(convError.Error())
			}
			program = append(program, integer)
		}
	}

	programError := part1(program)
	if programError != nil {
		fmt.Println(programError.Error())
		return
	}

	programError = part2(program)
	if programError != nil {
		fmt.Println(programError.Error())
		return
	}
}

func part1(programRef []int) error {

	program := make([]int, len(programRef))
	copy(program, programRef)

	program[1] = 12
	program[2] = 2

	programError := runProgram(&program)
	if programError == nil {
		fmt.Println("Value of position 0 after execution in part 1: ", program[0])
	}
	return programError

}

func part2(programRef []int) error {
	currentRun := make([]int, len(programRef))

	var verb, noun int

	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			copy(currentRun, programRef)
			currentRun[1] = noun
			currentRun[2] = verb

			programError := runProgram(&currentRun)
			if programError != nil {
				return programError
			}
			if currentRun[0] == 19690720 {
				fmt.Printf("\nNoun: %v \nVerb: %v\nAnswer: %v", noun, verb, 100*noun+verb)
				return programError
			}
		}
	}
	return fmt.Errorf("Could not find a combination of noun and verb that produce the number 19690720")
}

func runProgram(program *[]int) error {
	var pError error
	var pFinish bool
	lOffset := 0
	hOffset := 4

	for !pFinish {
		pFinish, pError = interpret((*program)[lOffset:hOffset], program)
		if pError != nil {
			return pError
		}

		lOffset += 4
		hOffset += 4
	}

	return pError

}

func interpret(instrct []int, memory *[]int) (bool, error) {

	switch op := instrct[0]; op {
	case 1, 2:
		for i := 1; i < 4; i++ {
			if instrct[i] > len(*memory) || instrct[i] < 0 {
				return false, fmt.Errorf("Invalid memory location: %v", instrct[i])
			}
		}

		if instrct[0] == 1 {
			(*memory)[instrct[3]] = (*memory)[instrct[1]] + (*memory)[instrct[2]]
		}

		if instrct[0] == 2 {
			(*memory)[instrct[3]] = (*memory)[instrct[1]] * (*memory)[instrct[2]]
		}
		return false, nil
	case 99:
		return true, nil
	default:
		return false, fmt.Errorf("Invalid OP-code : %v", op)
	}
}
