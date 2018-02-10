package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	input := make([]string, 0)
	output := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfCases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < numberOfCases; i++ {
		scanner.Scan()
		text := scanner.Text()
		input = append(input, text)
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	for i, v := range input {
		output = append(output, getSolution(i+1, v))
	}

	for _, v := range output {
		fmt.Println(v)
	}

}

func getSolution(caseNumber int, input string) string {
	output := "Case #" + strconv.Itoa(caseNumber) + ": "

	digits, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	if digits <= 4 {
		output += "..."
	} else {
		output += "IT'S OVER 9000"
		if digits >= 31683 {
			output += "!"
		} else {
			output += getExclamationMarks(digits)
		}
	}

	return output
}

func getExclamationMarks(digits int) string {
	var middle, multifuncResultDigits, result int
	var multifuncResult string
	max := 31682
	maxExclamationMarks := 2
	min := 5
	minExclamationMarks := 8990
	finalExclamationMarks := ""

	if digits > 15842 {
		return "!!"
	}

	for (digits < max || digits == max) && digits > min {
		if minExclamationMarks-maxExclamationMarks == 1 {
			result = minExclamationMarks
			break
		}
		middle = (minExclamationMarks + maxExclamationMarks) / 2
		multifuncResult = multifunctional(middle)
		multifuncResultDigits = len(multifuncResult)
		if multifuncResultDigits == digits-1 {
			for middle--; multifuncResultDigits < digits; middle-- {
				result = middle + 1
				multifuncResult = multifunctional(middle)
				multifuncResultDigits = len(multifuncResult)
			}
			break
		} else {
			if multifuncResultDigits >= digits {
				max = multifuncResultDigits
				maxExclamationMarks = middle
			} else {
				min = multifuncResultDigits
				minExclamationMarks = middle
			}
		}
	}

	for i := 0; i < result; i++ {
		finalExclamationMarks += "!"
	}

	return finalExclamationMarks
}

func multifunctional(exclamationMarks int) string {

	var factor int64
	result := big.NewInt(9000)

	for i := 1; i*exclamationMarks < 9000; i++ {
		factor = int64(9000 - (i * exclamationMarks))
		result.Mul(result, big.NewInt(factor))
	}

	return result.String()

}
