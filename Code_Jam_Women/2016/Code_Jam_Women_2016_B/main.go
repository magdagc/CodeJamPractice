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

	K, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	N, err := strconv.Atoi(inputs[2])
	if err != nil {
		panic(err)
	}

	output += getDancers(D, K, N)

	return output
}

func getDancers(dancers, dancer, turns int) string {
	danceFloor := make([]int, dancers)
	var dancerClockwise, dancerCounterClockwise int

	for i := 0; i < dancers; i++ {
		danceFloor[i] = i + 1
	}

	// When the turn is equal to the number of dances,
	// the dancefloor as when it started
	turns = turns % dancers

	for i := 1; i <= turns; i++ {
		clockwise := false
		if i%2 != 0 {
			clockwise = true
		}
		danceFloor = swap(danceFloor, clockwise)
	}

	dancerIndex := 0

	for i, v := range danceFloor {
		if v == dancer {
			dancerIndex = i
		}
	}

	if dancerIndex == dancers-1 {
		dancerClockwise = danceFloor[0]
	} else {
		dancerClockwise = danceFloor[dancerIndex+1]
	}

	if dancerIndex == 0 {
		dancerCounterClockwise = danceFloor[dancers-1]
	} else {
		dancerCounterClockwise = danceFloor[dancerIndex-1]
	}
	result := strconv.Itoa(dancerClockwise) + " " + strconv.Itoa(dancerCounterClockwise)

	return result
}

func swap(danceFloor []int, clockwise bool) []int {
	var aux int
	startIndex := 0
	if !clockwise {
		aux = danceFloor[0]
		danceFloor[0] = danceFloor[len(danceFloor)-1]
		danceFloor[len(danceFloor)-1] = aux
		startIndex = 1
	}

	for i := startIndex; i < len(danceFloor)-1-startIndex; i += 2 {
		aux = danceFloor[i]
		danceFloor[i] = danceFloor[i+1]
		danceFloor[i+1] = aux
	}

	return danceFloor
}
