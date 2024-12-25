package utils

import "fmt"

func ConvertToIDR(amount int64) string {

	amountStr := fmt.Sprintf("%d", amount)
	integerPart := amountStr[:len(amountStr)-2]
	// decimalPart := amountStr[len(amountStr)-2:]

	var formattedIntegerPart string
	count := 0
	for i := len(integerPart) - 1; i >= 0; i-- {
		formattedIntegerPart = string(integerPart[i]) + formattedIntegerPart
		count++

		if count%3 == 0 && i != 0 {
			formattedIntegerPart = "." + formattedIntegerPart
		}
	}

	return formattedIntegerPart
}
