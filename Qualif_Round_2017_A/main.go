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

	if len(inputs) != 2 {
		panic("Incorrect input")
	}

	pancakes := inputs[0]
	pancakeFlipper, err := strconv.Atoi(inputs[1])

	if err != nil {
		panic(err)
	}

	blankPancakes := strings.Count(pancakes, "-")

	if blankPancakes == 0 {
		output += "0"
	} else {
		output += calculateNeededFlips(pancakes, pancakeFlipper)
	}

	return output
}

func calculateNeededFlips(pancakes string, pancakeFlipper int) string {

	blankCounter := 0
	happyCounter := 0
	neededFlips := 0

	for _, v := range pancakes {
		if string(v) == "-" {
			if happyCounter > 0 {
				return "IMPOSSIBLE"
			}
			blankCounter++
			if blankCounter == pancakeFlipper {
				blankCounter = 0
				neededFlips++
			}
		} else {
			if blankCounter > 0 {
				happyCounter++
				if happyCounter+blankCounter == pancakeFlipper {
					neededFlips++
					blankCounter = happyCounter
					happyCounter = 0
				}
			}
		}
	}

	if blankCounter > 0 {
		return "IMPOSSIBLE"
	}

	return strconv.Itoa(neededFlips)
}
