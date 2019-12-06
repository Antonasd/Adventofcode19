package main

import (
	"adventofcode19/util"
	"fmt"
	"math"
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

	programError := runProgram(&program)
	if programError != nil {
		fmt.Println(programError.Error())
		return
	}
}

func runProgram(program *[]int) error {
	var pError error
	var pFinish bool
	pointer := 0
	offset := 0

	for !pFinish {
		op := getOpCode((*program)[pointer])
		if op == 3 || op == 4 {
			offset = 2
		} else if op == 5 || op == 6 {
			offset = 3
		} else {
			offset = 4
		}

		pFinish, pError = interpret((*program)[pointer:pointer+offset], program, &pointer)
		if pError != nil {
			fmt.Println("Failed at instruction: ", (*program)[pointer:pointer+offset])
			return pError
		}
	}
	return pError

}

func interpret(instrct []int, memory *[]int, pointer *int) (bool, error) {
	op := getOpCode(instrct[0])
	switch op {
	case 1, 2, 7, 8:

		var value1 int
		var value2 int
		var memError error

		if instrct[3] > len(*memory) || instrct[3] < 0 {
			return false, fmt.Errorf("Invalid memory location: %v", instrct[3])
		}

		value1, memError = getValue(instrct[1], getDigitAt(instrct[0], 2), memory)
		if memError != nil {
			return false, memError
		}

		value2, memError = getValue(instrct[2], getDigitAt(instrct[0], 3), memory)
		if memError != nil {
			return false, memError
		}

		*pointer += 4

		if op == 1 {
			(*memory)[instrct[3]] = value1 + value2
		} else if op == 2 {
			(*memory)[instrct[3]] = value1 * value2
		} else if op == 7 {
			if value1 < value2 {
				(*memory)[instrct[3]] = 1
			} else {
				(*memory)[instrct[3]] = 0
			}
		} else {
			if value1 == value2 {
				(*memory)[instrct[3]] = 1
			} else {
				(*memory)[instrct[3]] = 0
			}
		}
		return false, nil
	case 3, 4:

		var value int

		if instrct[1] > len(*memory) || instrct[1] < 0 {
			return false, fmt.Errorf("Invalid memory location: %v", instrct[3])
		}

		*pointer += 2

		if op == 3 {
			hasInput := false
			for !hasInput {

				input, inputError := util.GetInput()
				if inputError != nil {
					fmt.Println("Failed to read input")
					return false, inputError
				}

				var convErr error
				value, convErr = strconv.Atoi(input)
				if convErr != nil {
					fmt.Println("Not a number!")
					fmt.Println(convErr.Error())
				} else {
					hasInput = true
				}
			}
			(*memory)[instrct[1]] = value
		} else {
			value, _ = getValue(instrct[1], getDigitAt(instrct[0], 2), memory)
			fmt.Println("Value: ", value)
		}
		return false, nil
	case 5, 6:

		var value1 int
		var value2 int
		var memError error

		value1, memError = getValue(instrct[1], getDigitAt(instrct[0], 2), memory)
		if memError != nil {
			return false, memError
		}

		value2, memError = getValue(instrct[2], getDigitAt(instrct[0], 3), memory)
		if memError != nil {
			return false, memError
		}

		*pointer += 3

		if op == 5 && value1 != 0 {
			*pointer = value2
		} else if op == 6 && value1 == 0 {
			*pointer = value2
		}

		return false, nil
	case 99:
		return true, nil
	default:
		return false, fmt.Errorf("Invalid OP-code: %v", op)
	}
}

func getValue(value, valueType int, memory *[]int) (int, error) {
	if valueType == 0 {
		if value > len(*memory) || value < 0 {
			return value, fmt.Errorf("Invalid memory location: %v", value)
		}
		return (*memory)[value], nil
	}
	return value, nil
}

func getDigitAt(number, index int) int {
	return number/int(math.Pow10(index)) - (number/int(math.Pow10(index+1)))*10
}

func getOpCode(instr int) int {
	return getDigitAt(instr, 1)*10 + getDigitAt(instr, 0)
}
