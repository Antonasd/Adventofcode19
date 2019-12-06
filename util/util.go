package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFileInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")

	path, inputError := reader.ReadString('\n')
	if inputError != nil {
		fmt.Println("Failed to read input path:")
		return "", inputError
	}

	path = strings.TrimSuffix(path, "\n")
	path = strings.TrimSuffix(path, "\r")

	input, readError := ioutil.ReadFile(path)
	if readError != nil {
		fmt.Println("Failed to read input file:")
		return "", readError
	}

	return string(input), nil
}

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input number")
	fmt.Print(">")

	input, inputError := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input, inputError
}
