package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const reproduction int = 2
const decomission int = 1

var alpha int
var beta int

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

	acrobots, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	bouncoids, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}
	alpha, err = strconv.Atoi(inputs[2])
	if err != nil {
		panic(err)
	}
	beta, err = strconv.Atoi(inputs[3])
	if err != nil {
		panic(err)
	}
	years, err := strconv.Atoi(inputs[4])
	if err != nil {
		panic(err)
	}

	for i := 0; i < years; i++ {
		newAcrobots, newBouncoids := calculateNewPopulation(acrobots, bouncoids)
		if newAcrobots == acrobots && newBouncoids == bouncoids {
			break
		}
		acrobots = newAcrobots
		bouncoids = newBouncoids
	}

	output += strconv.Itoa(acrobots)
	output += " "
	output += strconv.Itoa(bouncoids)
	return output
}

func calculateNewPopulation(acrobots, bouncoids int) (int, int) {

	updatedAcrobots := acrobots - ((acrobots * decomission) / 100)
	updatedBouncoids := bouncoids - ((bouncoids * decomission) / 100)

	var couples int

	if acrobots < bouncoids {
		couples = acrobots
	} else {
		couples = bouncoids
	}

	babyAcrobots := (((couples * reproduction) / 100) * alpha) / 100
	babyBouncoids := (((couples * reproduction) / 100) * beta) / 100

	remainingBabies := ((couples * reproduction) / 100) - babyAcrobots - babyBouncoids

	babyAcrobots += (remainingBabies / 2)
	remainingBabies -= (remainingBabies / 2)

	babyBouncoids += remainingBabies

	updatedAcrobots += babyAcrobots
	updatedBouncoids += babyBouncoids

	return updatedAcrobots, updatedBouncoids
}
