package main

import (
	"fmt"
	"math"
)

func main() {
	a := 158126
	b := 624574

	passwords := findPossiblePasswords(a, b)
	fmt.Println("Part 1, Number of passwords in interval: ", len(passwords))

	passwords2 := findPossiblePasswordsPart2(a, b)
	fmt.Println("Part 2, Number of passwords in interval: ", len(passwords2))

}

func findPossiblePasswords(start, end int) (possiblePasswords []int) {
	currentNumber, nAdjacent := getFirstPossiblePassword(start)

	if currentNumber > end {
		return
	}

	for currentNumber <= end {
		var nNewPasswords int
		if nAdjacent == 0 {
			nNewPasswords = 0
		} else {
			nNewPasswords = distanceFromOverflow(currentNumber)
		}

		if currentNumber+nNewPasswords > end {
			possiblePasswords = append(possiblePasswords, generatePasswords(nNewPasswords-(currentNumber+nNewPasswords-end), currentNumber)...)
		} else {
			possiblePasswords = append(possiblePasswords, generatePasswords(nNewPasswords, currentNumber)...)
		}

		currentNumber += distanceFromOverflow(currentNumber) + 1
		overFlowIndex := getIndexOfFirstDigit(currentNumber)
		currentNumber = fillRightOf(currentNumber, overFlowIndex)

		if overFlowIndex >= nAdjacent {
			nAdjacent = overFlowIndex - 1
		}

	}
	return
}

func getFirstPossiblePassword(number int) (password int, iAdjacent int) {
	hasAdjacent := false
	dec1 := getDigitAt(number, 5)

	for i := 4; i >= 0; i-- {
		dec2 := getDigitAt(number, i)

		if dec2 < dec1 {
			return fillRightOf(number, i+1), dec2
		} else if !hasAdjacent && dec2 == dec1 {
			iAdjacent = i
			hasAdjacent = true
		}
		dec1 = dec2
	}

	if !hasAdjacent {
		return number + 10, 0
	}

	return number, iAdjacent
}

func fillRightOf(number, index int) int {
	repeatNumber := getDigitAt(number, index)
	baseNumber := (number / int(math.Pow10(index))) * int(math.Pow10(index))
	for i := index - 1; i >= 0; i-- {
		baseNumber += repeatNumber * int(math.Pow10(i))
	}

	return baseNumber
}

func getDigitAt(number, index int) int {
	return number/int(math.Pow10(index)) - (number/int(math.Pow10(index+1)))*10
}

//Finds the index of the first non-zero digit
func getIndexOfFirstDigit(number int) (index int) {
	for index = 0; index <= 5; index++ {
		if getDigitAt(number, index) != 0 {
			if index == 0 {
				return 0
			}

			return index
		}
	}
	return
}

func distanceFromOverflow(number int) int {
	return 9 - getDigitAt(number, 0)
}

func generatePasswords(nPasswords, base int) (passwords []int) {
	for n := 0; n <= nPasswords; n++ {
		passwords = append(passwords, base+n)
	}
	return
}

func findPossiblePasswordsPart2(start, end int) (possiblePasswords []int) {
	potentiallyValidPass := findPossiblePasswords(start, end)

	for _, pass := range potentiallyValidPass {
		streak := 0
		valid := false

		dec1 := getDigitAt(pass, 0)
		for i := 1; i <= 5; i++ {
			dec2 := getDigitAt(pass, i)
			if dec1 == dec2 {
				streak++
			} else {
				if streak == 1 {
					valid = true
					break
				}
				streak = 0
			}
			dec1 = dec2
		}

		if valid || streak == 1 {
			possiblePasswords = append(possiblePasswords, pass)
		}

	}
	return
}
