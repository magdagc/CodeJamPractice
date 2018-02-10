package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	items := make([]int, 0)
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
		nbItems, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		items = append(items, nbItems)
		scanner.Scan()
		text = scanner.Text()
		input = append(input, text)
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	for i, v := range input {
		output = append(output, getSolution(i+1, v, items[i]))
	}

	for _, v := range output {
		fmt.Println(v)
	}

}

func getSolution(caseNumber int, input string, nbItems int) string {

	mixedItemPrices := make([]int, 0)
	output := "Case #" + strconv.Itoa(caseNumber) + ": "
	inputs := strings.Split(input, " ")

	for _, v := range inputs {
		price, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		mixedItemPrices = append(mixedItemPrices, price)
	}

	output += getSalePrices(nbItems, mixedItemPrices)

	return output
}

func getSalePrices(nbItems int, mixedItemPrices []int) string {

	var originalPrice int
	salePrices := ""

	if nbItems*2 != len(mixedItemPrices) {
		panic("Error in the input")
	}

	for i, v := range mixedItemPrices {
		originalPrice = int(float64(v) / 0.75)
		mixedItemPrices = removePrice(originalPrice, i, mixedItemPrices)
		if i+1 >= len(mixedItemPrices) {
			break
		}
	}

	for _, v := range mixedItemPrices {
		salePrices += strconv.Itoa(v) + " "
	}

	return salePrices[:len(salePrices)-1]
}

func removePrice(price, startIndex int, priceList []int) []int {

	foundPrice := false

	for i, v := range priceList {
		if i > startIndex {
			if v == price {
				priceList = append(priceList[:i], priceList[i+1:]...)
				foundPrice = true
				break
			}
		}
	}

	if !foundPrice {
		errMsg := fmt.Sprintf("Price not found %v", price)
		panic(errMsg)
	}

	return priceList
}
