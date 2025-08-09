package main

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {

	for {
		currentCurrency, currentAmount, targetCurrency, err := getUserCurrentCurrencyAndAmount()
		if err != nil {
			fmt.Println(err)
			continue
		}

		calculate(currentCurrency, currentAmount, targetCurrency)
	}

}

func getUserCurrentCurrencyAndAmount() (string, int32, string, error) {
	var currentCurrency string
	var currentAmount int32
	var targetCurrency string
	fmt.Print("Введите исходную валюту (USD, EUR, RUR): ")
	fmt.Scan(&currentCurrency)
	fmt.Print("Сколько исходной валюты хотите сконвертировать: ")
	fmt.Scan(&currentAmount)
	fmt.Print("Введите целевую валюту (USD, EUR, RUR): ")
	fmt.Scan(&targetCurrency)

	if currentCurrency != "USD" || currentCurrency != "EUR" || currentCurrency != "RUR" {
		return "_", 0, "_", errors.New("Ошибка, неизвестная исходная валюта, повторите ввод заново")
	}

	if currentCurrency == targetCurrency {
		return "_", 0, "_", errors.New("Ошибка, исходная и целевая валюта совпадают, повторите ввод заново")
	}

	if !unicode.IsDigit(currentAmount) || currentAmount <= 0 {
		return "_", 0, "_", errors.New("Ошибка, введите количество валюты заново")
	}

	if targetCurrency != "USD" || targetCurrency != "EUR" || targetCurrency != "RUR" {
		return "_", 0, "_", errors.New("Ошибка, неизвестная целевая валюта, повторите ввод заново")
	}

	return currentCurrency, currentAmount, targetCurrency, nil
}

func calculate(currentCurrency string, currentAmount int32, targetCurrency string) {

	switch {
	case currentCurrency == "USD" && targetCurrency == "EUR":
		result := float64(currentAmount) * 0.85
		fmt.Printf("%d %s равно %.2f %s", currentAmount, currentCurrency, result, targetCurrency)
	case currentCurrency == "USD" && targetCurrency == "RUR":
		result := float64(currentAmount) * 80
		fmt.Printf("%d %s равно %.2f %s", currentAmount, currentCurrency, result, targetCurrency)
	case currentCurrency == "EUR" && targetCurrency == "RUR":
		result := float64(currentAmount) * 90
		fmt.Printf("%d %s равно %.2f %s", currentAmount, currentCurrency, result, targetCurrency)
	}
}
