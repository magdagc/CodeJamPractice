package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	N, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic("Error in input format")
	}
	S, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic("Error in input format")
	}
	p, err := strconv.Atoi(inputs[2])
	if err != nil {
		panic("Error in input format")
	}

	totalPoints := make([]int, 0)

	for _, v := range inputs[3:] {
		points, err := strconv.Atoi(v)
		if err != nil {
			panic("Error in input format")
		}
		totalPoints = append(totalPoints, points)
	}

	if N != len(totalPoints) {
		panic("Error in the input information")
	}

	output += strconv.Itoa(getMaximumGooglersWithP(S, p, totalPoints))

	return output
}

func getMaximumGooglersWithP(S, p int, totalPoints []int) int {

	sort.Ints(totalPoints)
	counterMoreThanP := 0
	counterS := 0

	if p == 0 && S == 0 {
		return len(totalPoints)
	}

	for _, v := range totalPoints {
		// Not enough surprising
		if counterS < S {
			if v >= 2 {
				if v >= (((p - 2) * 2) + p) {
					counterS++
					counterMoreThanP++
				}
			} else {
				if p <= v {
					counterMoreThanP++
				}
			}
			// Surprising points fulfilled
		} else {
			if v >= 1 {
				if v >= (((p - 1) * 2) + p) {
					counterMoreThanP++
				}
			} else {
				if p == 0 {
					counterMoreThanP++
				}
			}
		}

	}

	return counterMoreThanP
}
