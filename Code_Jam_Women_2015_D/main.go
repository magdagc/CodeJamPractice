package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stage [][]int

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

	rows, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}

	columns, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	var max int

	if rows > columns {
		max = rows
	} else {
		max = columns
	}

	stage = make([][]int, max)
	for i := 0; i < max; i++ {
		stage[i] = make([]int, max)
		for j := 0; j < max; j++ {
			stage[i][j] = -1
		}
	}

	output += strconv.Itoa(getDifferentPaths(rows, columns))

	return output
}

func getDifferentPaths(r, c int) int {
	if r == 1 || c == 1 {
		return 1
	}

	if stage[r-1][c-1] != -1 {
		return stage[r-1][c-1]
	}

	ans := 0

	for i := 1; i <= r; i++ {
		ans += getDifferentPaths(c-1, i)
	}

	stage[r-1][c-1] = ans

	return ans

}
