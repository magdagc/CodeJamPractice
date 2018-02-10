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
	output := "Case #" + strconv.Itoa(caseNumber) + ":"

	inputs := strings.Split(input, " ")
	N, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	K, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	A := N / K
	B := N % K

	bigTables := combinations(N, A+1)
	smallTables := 0
	if B > 0 {

	}

	return output
}

func combinations(n, p int) int {
	factorialN := factorial(n)
	factorialP := factorial(p)
	factorialNMinusP := factorial(n - p)

	result := (factorialN / (factorialP * factorialNMinusP))

	return result
}

func factorial(a int) int {

	result := 1

	if a < 3 {
		return a
	}

	for i := 2; i <= a; i++ {
		result *= i
	}
	return result
}
