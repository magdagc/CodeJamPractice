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

	amountOfBytes, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	pseudoBinaryString := inputs[1]

	if amountOfBytes*8 != len(pseudoBinaryString) {
		panic("Error in the format")
	}

	pseudoBinaryString = strings.Replace(pseudoBinaryString, "I", "1", -1)
	pseudoBinaryString = strings.Replace(pseudoBinaryString, "O", "0", -1)

	inputBytes := make([]byte, 0)

	for i := 0; i < len(pseudoBinaryString); i += 8 {
		a, _ := strconv.ParseInt(pseudoBinaryString[i:8+i], 2, 8)
		b := uint8(a)
		inputBytes = append(inputBytes, b)
	}

	output += string(inputBytes)

	return output
}
