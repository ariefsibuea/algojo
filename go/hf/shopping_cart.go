package hf

import (
	"math"
	"strconv"
)

type Discount struct {
	Type   int32
	Amount int32
}

const EmptyDiscount = "EMPTY"

var Discounts = map[string]Discount{}

func FindLowestPrice(products [][]string, discounts [][]string) int32 {
	// set all of discount types
	setDiscounts(discounts)

	// count final price of every product
	totalPrice := int32(0)
	for i := range products {
		price, _ := strconv.Atoi(products[i][0])
		minPrice := countMinimumPrice(int32(price), products[i][1:])
		totalPrice = totalPrice + minPrice
	}

	return totalPrice
}

func countMinimumPrice(price int32, tags []string) int32 {
	minPrice := price
	for i := range tags {
		if tags[i] == EmptyDiscount {
			continue
		}
		tag, ok := Discounts[tags[i]]
		if !ok {
			continue
		}

		discPrice := int32(0)
		switch tag.Type {
		case 0:
			discPrice = tag.Amount
		case 1:
			discAmount := float64(price*tag.Amount) / 100
			discPrice = int32(math.Round((float64(price) - discAmount)))
		case 2:
			discPrice = price - tag.Amount
		}
		if minPrice > discPrice {
			minPrice = discPrice
		}
	}

	return minPrice
}

func setDiscounts(discounts [][]string) {
	for i := range discounts {
		// format discount [tag, type, amount]
		tag := discounts[i][0]
		typ, _ := strconv.Atoi(discounts[i][1])
		amn, _ := strconv.Atoi(discounts[i][2])
		disc := Discount{
			Type:   int32(typ),
			Amount: int32(amn),
		}
		Discounts[tag] = disc
	}
}
