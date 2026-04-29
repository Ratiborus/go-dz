package main

import (
	"fmt"
	"strings"
)

const Usd = "usd"
const Eur = "eur"
const Rub = "rub"
const UsdToEur = 0.92
const UsdToRub = 75.0
const EurToRub = UsdToRub / UsdToEur

func main() {
	fmt.Println("___Конвертер валюты___")
	for {
		source := getSourceInput()
		amount := getAmountInput()
		target := getTragetInput(source)
		result := calculate(amount, source, target)
		fmt.Printf("Результат конвертации %v (%v) = %.2f (%v)\n", amount, source, result, target)
		if !proceed() {
			break
		}
	}
}

func getSourceInput() string {
	var source string
	for {
		fmt.Printf("Введите исходную валюту для конвертации (%v, %v, %v): ", Usd, Eur, Rub)
		_, err := fmt.Scan(&source)
		if isCorrectInput(err, source) {
			break
		}
	}
	return source
}

func getTragetInput(source string) string {
	var target string
	for {
		listCurrencies := getListAvailableCurrencies(source)
		fmt.Printf("Введите целевую валюту для конвертации (%v): ", listCurrencies)
		_, err := fmt.Scan(&target)
		if isCorrectInput(err, target) && isCorrectTargetInput(source, target) {
			break
		}
	}
	return target
}

func getListAvailableCurrencies(source string) string {
	pattern := "%v, %v"
	switch source {
	case Usd:
		return fmt.Sprintf(pattern, Eur, Rub)
	case Eur:
		return fmt.Sprintf(pattern, Usd, Rub)
	case Rub:
		return fmt.Sprintf(pattern, Usd, Eur)
	default:
		panic(fmt.Sprintf("Incorrect currency '%v'", source))
	}
}

func isCorrectInput(err error, val string) bool {
	if err != nil {
		fmt.Printf("Произошла ошибка: %v,  повторите ввод\n", err)
		return false
	} else if unknownCurrency(val) {
		fmt.Printf("Вы ввели неизвестную валюту '%v', повторите ввод \n", val)
		return false
	}
	return true
}

func isCorrectTargetInput(source, target string) bool {
	if source == target {
		fmt.Println("Исходная и целевая валюта должны отличаться")
		return false
	}
	switch source {
	case Usd:
		return target == Eur || target == Rub
	case Eur:
		return target == Usd || target == Rub
	case Rub:
		return target == Usd || target == Eur
	default:
		panic(fmt.Sprintf("Unknown currency: %v", source))
	}
}

func unknownCurrency(val string) bool {
	normalized := strings.ToLower(val)
	return !(normalized == Usd || normalized == Eur || normalized == Rub)
}

func getAmountInput() int {
	var amount int
	for {
		fmt.Print("Введите количество :  ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Printf("Произошла ошибка при вводе количества: %v,  повторите ввод\n", err)
		} else if amount < 1 {
			fmt.Println("Необходимо целое положительное число, повторите ввод")
		} else {
			return amount
		}
	}
}

func calculate(amount int, source, target string) float64 {
	switch source {
	case Usd:
		switch target {
		case Eur:
			return float64(amount) * UsdToEur
		case Rub:
			return float64(amount) * UsdToRub
		default:
			panic(fmt.Sprintf("Incorrect target input (source = %v, target = %v)", source, target))
		}
	case Eur:
		switch target {
		case Usd:
			return float64(amount) / UsdToEur
		case Rub:
			return float64(amount) * EurToRub
		default:
			panic(fmt.Sprintf("Incorrect target input (source = %v, target = %v)", source, target))
		}
	case Rub:
		switch target {
		case Usd:
			return float64(amount) / UsdToRub
		case Eur:
			return float64(amount) / EurToRub
		default:
			panic(fmt.Sprintf("Incorrect target input (source = %v, target = %v)", source, target))
		}
	default:
		panic(fmt.Sprintf("Unknown currency '%v'", source))
	}
}

func proceed() bool {
	fmt.Print("Хотите выполнить новый расчет (y/n)? ")
	var answer string
	fmt.Scan(&answer)
	return strings.EqualFold("y", answer)
}
