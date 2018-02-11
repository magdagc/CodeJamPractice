package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	row    int
	column int
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
		text := scanner.Text() + " "
		friends, _ := strconv.Atoi(strings.Split(text, " ")[0])
		for i := 0; i < friends; i++ {
			// Tickets
			scanner.Scan()
			text += scanner.Text() + " "
		}
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

	friends, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}

	rows, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	tickets := make([]ticket, friends)

	inputs = inputs[2:]

	for i := 0; i < friends*2; i += 2 {
		inputRow, _ := strconv.Atoi(inputs[i])
		inputColumn, _ := strconv.Atoi(inputs[i+1])
		tickets[i/2] = ticket{row: inputRow, column: inputColumn}
	}

	friendsPerRow := make([]int, rows)

	output += getMaxFriendsPerRow(tickets, friendsPerRow)

	return output
}

func getMaxFriendsPerRow(tickets []ticket, friendsPerRow []int) string {

	return ""
}
