package util

import (
	"strings"
	"strconv"
	"log"
)

func ConvertPrice(price string) int {
	priceNumbers := price[0:len(price)-6]
	trimedPrice := removeSpaces(priceNumbers)
	convertedPrice, err := strconv.Atoi(trimedPrice)
	if (err != nil) {
		log.Fatal(err)
	}
	return convertedPrice
}

func removeSpaces(price string) string {
	priceWithoutSpaces := strings.Replace(price, " ", "", -1)
	priceWithoutThinSpaces := strings.Replace(priceWithoutSpaces, "\u2009", "", -1) //to remove thin space
	return priceWithoutThinSpaces
}
