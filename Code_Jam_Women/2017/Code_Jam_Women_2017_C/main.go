package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const completeRow = "I/O/I/O/I/O/I/O\n"

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

	if N == 0 {
		return "IO"
	}

	removeFromTopBottom := (287 - N) % 3 //1
	removeFromMiddle := (287 - N) / 3    //94

	// Max amount of removed I/O in the middle is 273
	// That is 91 (273/3) slashes
	// After that, keep removing from top or bottom
	if removeFromMiddle > 91 {
		// Total amount of slashes in top+bottom minus N
		removeFromTopBottom = (14 - N)
		removeFromMiddle = 91
	}

	// Top and bottom lines
	// Remove at max 7 slashes
	top := completeRow
	bottom := completeRow
	if removeFromTopBottom > 7 {
		top = strings.Replace(top, "/", "O", 7)
		removeFromTopBottom -= 7
		bottom = strings.Replace(bottom, "/", "O", removeFromTopBottom)
	} else {
		top = strings.Replace(top, "/", "O", removeFromTopBottom)
	}

	result := top

	rowsToRemoveAllSlashes := removeFromMiddle / 7
	slashesToRemoveFromOneRow := removeFromMiddle % 7

	for i := 0; i < 13; i++ {
		if slashesToRemoveFromOneRow > 0 {
			result += strings.Replace(completeRow, "/", "O", slashesToRemoveFromOneRow)
			slashesToRemoveFromOneRow = 0
		} else {
			if rowsToRemoveAllSlashes > 0 {
				result += strings.Replace(completeRow, "/", "O", 7)
				rowsToRemoveAllSlashes--
			} else {
				result += completeRow
			}
		}
	}

	result += bottom

	return result[:len(result)-1]
}
