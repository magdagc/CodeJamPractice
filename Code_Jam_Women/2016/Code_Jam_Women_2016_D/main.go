package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

	nbOfWords, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}

	passwords := inputs[1:]

	if nbOfWords != len(passwords) {
		panic("Error in the input")
	}

	output += getPossibleAlphabet(passwords)

	return output
}

func getPossibleAlphabet(passwords []string) string {

	result := alphabet

	for _, v := range passwords {
		if len(v) == 1 {
			result = "IMPOSSIBLE"
			break
		}
		if lettersRepeated(v) {
			continue
		}
	}

	return result
}

func lettersRepeated(word string) bool {
	repeatedLetters := false

	for i := 0; i < len(alphabet); i++ {
		if strings.Count(word, alphabet[i:i+1]) > 1 {
			repeatedLetters = true
			break
		}
	}

	return repeatedLetters
}
