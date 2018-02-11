package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var solutions []*big.Int

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

	C, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}

	V, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	L, err := strconv.Atoi(inputs[2])
	if err != nil {
		panic(err)
	}

	solutions = make([]*big.Int, L+1)

	zero := big.NewInt(0)

	for i := 0; i <= L; i++ {
		solutions[i] = zero
	}

	solutions[0] = big.NewInt(1)
	if L >= 1 {
		solutions[1] = big.NewInt(int64(V))
	}
	combinationsInt := getCombinations(big.NewInt(int64(C)), big.NewInt(int64(V)), int64(L))
	combinationsInt = combinationsInt.Mod(combinationsInt, big.NewInt(1000000007))
	combinationsString := fmt.Sprint(combinationsInt)

	output += combinationsString

	return output
}

func getCombinations(C, V *big.Int, L int64) *big.Int {

	if solutions[L].Cmp(big.NewInt(0)) != 0 {
		return solutions[L]
	}

	result := big.NewInt(0)
	a := big.NewInt(0)
	b := big.NewInt(0)
	d := big.NewInt(0)

	a.Mul(getCombinations(C, V, L-1), V)

	d.Mul(C, V)
	b.Mul(getCombinations(C, V, L-2), d)

	result.Add(a, b)

	solutions[L] = result

	return result
}
