package main

import (
	"bufio"
	"fmt"
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
	output := "Case #" + strconv.Itoa(caseNumber) + ":"

	output += "\n"

	headerFooter := "+"
	headerFooter += dashFilling(len(input) + 2)
	headerFooter += "+"
	headerFooter += "\n"

	output += headerFooter
	output += "|"
	output += " " + input + " "
	output += "|\n"
	output += headerFooter

	return output
}

func dashFilling(length int) string {
	output := ""
	for i := 0; i < length; i++ {
		output += "-"
	}
	return output
}
