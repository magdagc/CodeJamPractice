package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
LetterReplacer replaces the google language letters to translate words to English
*/
var LetterReplacer = strings.NewReplacer(
	"a", "y",
	"b", "h",
	"c", "e",
	"d", "s",
	"e", "o",
	"f", "c",
	"g", "v",
	"h", "x",
	"i", "d",
	"j", "u",
	"k", "i",
	"l", "g",
	"m", "l",
	"n", "b",
	"o", "k",
	"p", "r",
	"q", "z",
	"r", "t",
	"s", "n",
	"t", "w",
	"u", "j",
	"v", "p",
	"w", "f",
	"x", "m",
	"y", "a",
	"z", "q")

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

	output += LetterReplacer.Replace(input)

	return output
}
