package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	inputs := strings.Split(input, " ")

	D, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}

	N, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	output += getGrid(D, N)

	return output
}

func getGrid(D, N int) string {

	if N == 0 {
		return "IO"
	}

	return ""
}
