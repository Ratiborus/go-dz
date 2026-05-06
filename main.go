package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

const Usd = "usd"
const Eur = "eur"
const Rub = "rub"
const UsdToEur = 0.92
const UsdToRub = 75.0
const EurToRub = UsdToRub / UsdToEur

type ConvertMap = map[string]ConvertData
type ConvertData = map[string]float64

func main() {

	fmt.Println("___Конвертер валюты___")
	var convertMap = ConvertMap{
		"usd": ConvertData{"eur": UsdToEur, "rub": UsdToRub},
		"eur": ConvertData{"usd": 1 / UsdToEur, "rub": EurToRub},
		"rub": ConvertData{"usd": 1 / UsdToRub, "eur": 1 / EurToRub},
	}
	for {
		source := getSourceInput(&convertMap)
		amount := getAmountInput()
		target := getTragetInput(source, &convertMap)
		result := calculate(amount, source, target, &convertMap)
		fmt.Printf("Результат конвертации %v (%v) = %.2f (%v)\n", amount, source, result, target)
		if !proceed() {
			break
		}
	}
}

func getSourceInput(convertMap *ConvertMap) string {
	var source string
	for {
		fmt.Printf("Введите исходную валюту для конвертации (%v, %v, %v): ", Usd, Eur, Rub)
		_, err := fmt.Scan(&source)
		if isCorrectInput(err, source, convertMap) {
			break
		}
	}
	return strings.ToLower(source)
}

func getTragetInput(source string, convertMap *ConvertMap) string {
	var target string
	for {
		listCurrencies := maps.Keys((*convertMap)[source])
		fmt.Printf("Введите целевую валюту для конвертации %v: ", slices.Collect(listCurrencies))
		_, err := fmt.Scan(&target)
		target = strings.ToLower(target)
		if isCorrectInput(err, target, convertMap) && isCorrectTargetInput(source, target, convertMap) {
			break
		}
	}
	return target
}

func isCorrectInput(err error, val string, convertmap *ConvertMap) bool {
	if err != nil {
		fmt.Printf("Произошла ошибка: %v,  повторите ввод\n", err)
		return false
	} else if unknownCurrency(val, convertmap) {
		fmt.Printf("Вы ввели неизвестную валюту '%v', повторите ввод \n", val)
		return false
	}
	return true
}

func isCorrectTargetInput(source, target string, convertMap *ConvertMap) bool {
	if source == target {
		fmt.Println("Исходная и целевая валюта должны отличаться")
		return false
	}
	nestedMap, ok := (*convertMap)[source]
	if !ok {
		panic(fmt.Sprintf("Unknown currency: %v", source))
	}
	_, ok = nestedMap[target]
	return ok
}

func unknownCurrency(val string, convertMap *ConvertMap) bool {
	normalized := strings.ToLower(val)
	_, ok := (*convertMap)[normalized]
	return !ok
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

func calculate(amount int, source, target string, convertMap *ConvertMap) float64 {
	return float64(amount) * (*convertMap)[source][target]
}

func proceed() bool {
	fmt.Print("Хотите выполнить новый расчет (y/n)? ")
	var answer string
	fmt.Scan(&answer)
	return strings.EqualFold("y", answer)
}
