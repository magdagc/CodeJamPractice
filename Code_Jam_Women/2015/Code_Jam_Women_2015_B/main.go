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

	K, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	V, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	output += getBlandColors(K, V)

	return output
}

func getBlandColors(K, V int) string {

	blandColorAmount := powerOf((V + 1), 3)

	if K > V {
		blandColorAmount *= (K + 1 - V)
		blandColorAmount -= (powerOf(V, 3) * (K - V))
	}

	result := strconv.Itoa(blandColorAmount)

	return result

}

func powerOf(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
