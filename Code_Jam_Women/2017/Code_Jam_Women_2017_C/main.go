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
	output := "Case #" + strconv.Itoa(caseNumber) + ":\n"

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

	var grid []string

	if N == 0 {
		return "IO"
	}

	if N == 1 {
		return "I/O"
	}

	if N <= 5 {
		grid = make([]string, 3)
		grid[0] = "OOOOO"
		grid[2] = "O/I/O"

		switch N {
		case 2:
			grid[1] = "OOOOO"
		case 3:
			grid[1] = "OO/OO"
		case 4:
			grid[1] = "O//OO"
		case 5:
			grid[1] = "O///O"
		}

	}

	if N <= 8 {
		grid = make([]string, 5)
		grid[0] = "OOOOO"
		grid[1] = "O///O"
		grid[2] = "O/I/O"
		grid[4] = "OOOOO"

		switch N {
		case 6:
			grid[3] = "OO/OO"
		case 7:
			grid[3] = "O//OO"
		case 8:
			grid[3] = "O///O"
		}

	}

	result := ""

	for _, v := range grid {
		result += v + "\n"
	}

	return result[:len(result)-1]
}
