package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type role struct {
	biggerProbability float64
	lowerProbability  float64
}

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
		scanner.Scan()
		text += " " + scanner.Text()
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
		panic(err)
	}

	inputs = inputs[1:]
	probabilities := make([]float64, N*2)

	for i := 0; i < N*2; i++ {
		probabilities[i], _ = strconv.ParseFloat(inputs[i], 64)
	}

	output += getSucceedProbability(probabilities, N)

	return output
}

func getSucceedProbability(probabilities []float64, roles int) string {

	var result float64 = 1
	roleArray := make([]role, roles)

	sort.Float64s(probabilities)

	for i, v := range probabilities {
		if i < roles {
			roleArray[i].lowerProbability = v
		} else {
			roleArray[roles*2-1-i].biggerProbability = v
		}
	}

	for _, r := range roleArray {
		result *= (1 - (r.biggerProbability * r.lowerProbability))
	}

	return fmt.Sprintf("%6f", result)

}
