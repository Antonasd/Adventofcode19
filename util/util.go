package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please provide the path to your puzzle input: ")
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
