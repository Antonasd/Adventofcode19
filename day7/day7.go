package main

import (
	"adventofcode19/intcode"
	"adventofcode19/util"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
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
				return
			}
			program = append(program, integer)
		}
	}

	fmt.Println(program)
	settings := getCombinations([]int{5, 6, 7, 8, 9})
	runCombinationsFeedback(settings, program)
}

func runCombinationsFeedback(settings [][]int, program []int) {
	aeIO := make(chan int)
	baIO := make(chan int)
	cbIO := make(chan int)
	dcIO := make(chan int)
	edIO := make(chan int)

	ampA := makeIntcodeInstance(aeIO, baIO)
	ampB := makeIntcodeInstance(baIO, cbIO)
	ampC := makeIntcodeInstance(cbIO, dcIO)
	ampD := makeIntcodeInstance(dcIO, edIO)
	ampE := makeIntcodeInstance(edIO, aeIO)

	maxValue := math.MinInt64
	bestSetting := []int{}
	fmt.Println("Amoutn of different settings: ", len(settings))
	for _, setting := range settings {
		var waitgroup sync.WaitGroup
		waitgroup.Add(4)

		go runInstance(ampE, program)
		edIO <- setting[4]

		go func() {
			runInstance(ampD, program)
			waitgroup.Done()
		}()
		dcIO <- setting[3]

		go func() {
			runInstance(ampC, program)
			waitgroup.Done()
		}()
		cbIO <- setting[2]

		go func() {
			runInstance(ampB, program)
			waitgroup.Done()
		}()
		baIO <- setting[1]

		go func() {
			runInstance(ampA, program)
			waitgroup.Done()
		}()
		aeIO <- setting[0]
		aeIO <- 0

		waitgroup.Wait()

		outPut := <-aeIO
		if outPut > maxValue {
			maxValue = outPut
			bestSetting = setting
		}
	}

	fmt.Println("Max thruster signal: ", maxValue)
	fmt.Println("Best setting: ", bestSetting)
}

func runCombinations(settings [][]int, program []int) {
	mainOut := make(chan int)
	baIO := make(chan int)
	cbIO := make(chan int)
	dcIO := make(chan int)
	edIO := make(chan int)
	mainIn := make(chan int)

	ampA := makeIntcodeInstance(mainOut, baIO)
	ampB := makeIntcodeInstance(baIO, cbIO)
	ampC := makeIntcodeInstance(cbIO, dcIO)
	ampD := makeIntcodeInstance(dcIO, edIO)
	ampE := makeIntcodeInstance(edIO, mainIn)

	maxValue := math.MinInt64
	bestSetting := []int{}
	fmt.Println("Amoutn of different settings: ", len(settings))
	for _, setting := range settings {
		go runInstance(ampE, program)
		edIO <- setting[4]

		go runInstance(ampD, program)
		dcIO <- setting[3]

		go runInstance(ampC, program)
		cbIO <- setting[2]

		go runInstance(ampB, program)
		baIO <- setting[1]

		go runInstance(ampA, program)
		mainOut <- setting[0]
		mainOut <- 0

		outPut := <-mainIn
		if outPut > maxValue {
			maxValue = outPut
			bestSetting = setting
		}
	}

	fmt.Println("Max thruster signal: ", maxValue)
	fmt.Println("Best setting: ", bestSetting)
}

func makeIntcodeInstance(i chan int, o chan int) intcode.Intcode {
	return intcode.NewIntocdeInstance(
		func() (int, error) {
			return <-i, nil
		}, func(val int) {
			o <- val
		})
}

func runInstance(instance intcode.Intcode, program []int) {
	programInstance := make([]int, len(program))
	copy(programInstance, program)
	instance.RunProgram(&programInstance)
}

func getCombinations(settings []int) (combinations [][]int) {
	if len(settings) == 0 {
		combinations = [][]int{{}}
		return
	} else if len(settings) == 1 {
		combinations = append(combinations, settings)
		return
	}

	for i := 0; i < len(settings); i++ {
		settingsCopy := make([]int, len(settings))
		copy(settingsCopy, settings)

		subsetCombos := [][]int{{}}
		if i == len(settings)-1 {
			subsetCombos = getCombinations(settingsCopy[:i])
		} else {
			subsetCombos = getCombinations(append(settingsCopy[:i], settingsCopy[i+1:]...))
		}
		for index, comb := range subsetCombos {
			subsetCombos[index] = append(comb, settings[i])
		}
		combinations = append(combinations, subsetCombos...)
	}
	return
}
