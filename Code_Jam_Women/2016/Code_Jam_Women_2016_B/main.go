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

	output += getDancers(int64(D), int64(K), int64(N))

	return output
}

func getDancers(dancers, dancer, turns int64) string {
	var dancerClockwise, dancerCounterClockwise int64

	if dancer%2 != 0 {
		dancerClockwise = ((dancer + (2 * turns)) % dancers) + 1
		dancerCounterClockwise = ((dancer - 2 + (2 * turns)) % dancers) + 1
	} else {
		aux := dancer - (2 * turns)
		if aux < dancers*(-1) {
			aux = (aux * -1) % dancers
			aux *= -1
		}
		dancerClockwise = ((dancers + aux) % dancers) + 1

		aux -= 2
		if aux < dancers*(-1) {
			aux = (aux * -1) % dancers
			aux *= -1
		}
		dancerCounterClockwise = ((dancers + aux) % dancers) + 1
	}

	result := fmt.Sprintf("%v %v", dancerClockwise, dancerCounterClockwise)

	return result
}
