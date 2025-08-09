package main

import (
	"errors"
	"fmt"
)

func main() {

	for {
		currentCurrency, currentAmount, targetCurrency := getUserCurrentCurrencyAndAmount()

		err := check(currentCurrency, currentAmount, targetCurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}

		calculate(currentCurrency, currentAmount, targetCurrency)
	}

}

func getUserCurrentCurrencyAndAmount() (string, int32, string) {
	var currentCurrency string
	var currentAmount int32
	var targetCurrency string
	fmt.Print("Введите исходную валюту (USD, EUR, RUR): ")
	fmt.Scan(&currentCurrency)
	fmt.Print("Сколько исходной валюты хотите сконвертировать: ")
	fmt.Scan(&currentAmount)
	fmt.Print("Введите целевую валюту (USD, EUR, RUR): ")
	fmt.Scan(&targetCurrency)
	return currentCurrency, currentAmount, targetCurrency
}

func check(currentCurrency string, currentAmount int32, targetCurrency string) error {

	if currentCurrency != "USD" && currentCurrency != "EUR" && currentCurrency != "RUR" {
		return errors.New("Ошибка, неизвестная исходная валюта, повторите ввод заново")
	}

	if currentCurrency == targetCurrency {
		return errors.New("Ошибка, исходная и целевая валюта совпадают, повторите ввод заново")
	}

	if currentAmount <= 0 {
		return errors.New("Ошибка, введите количество валюты заново")
	}

	if targetCurrency != "USD" && targetCurrency != "EUR" && targetCurrency != "RUR" {
		return errors.New("Ошибка, неизвестная целевая валюта, повторите ввод заново")
	}

	return nil
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
	case currentCurrency == "EUR" && targetCurrency == "USD":
		result := float64(currentAmount) * 1.17
		fmt.Printf("%d %s равно %.2f %s", currentAmount, currentCurrency, result, targetCurrency)
	case currentCurrency == "RUR" && targetCurrency == "USD":
		result := float64(currentAmount) * 0.013
		fmt.Printf("%d %s равно %.2f %s", currentAmount, currentCurrency, result, targetCurrency)
	case currentCurrency == "RUR" && targetCurrency == "EUR":
		result := float64(currentAmount) * 0.011
		fmt.Printf("%d %s равно %.2f %s", currentAmount, currentCurrency, result, targetCurrency)
	}
}
