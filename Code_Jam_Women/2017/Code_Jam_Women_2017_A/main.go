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

	output += getMaxFriendsPerRow(tickets, rows)

	return output
}

func getMaxFriendsPerRow(tickets []ticket, rows int) string {

	maxFriends := 1

	for i := 0; i < rows; i++ {
		maxFriends = getMaxFriends(i+1, tickets, maxFriends)
	}

	return strconv.Itoa(maxFriends)
}

func (t ticket) isRowPossible(r int) bool {
	return r == t.row || r == t.column
}

func (t ticket) exists(tickets []ticket) bool {
	appearedOneTime := false
	for _, v := range tickets {
		if v.row == t.row && v.column == t.column {
			if appearedOneTime {
				return true
			}
			appearedOneTime = true
		}
	}
	return false
}

func getMaxFriends(row int, tickets []ticket, currentMaxFriends int) int {
	maxFriends := 0
	repeatedTickets := 0
	for _, t := range tickets {
		if t.isRowPossible(row) {
			maxFriends++
			if t.exists(tickets) {
				repeatedTickets++
			}
		}
	}

	maxFriends -= repeatedTickets / 2

	if currentMaxFriends > maxFriends {
		return currentMaxFriends
	}
	return maxFriends

}
