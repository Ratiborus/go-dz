package main

import "fmt"

func main() {
	const UsdToEur = 0.92
	const UsdToRub = 75.0
	const EurToRub = UsdToRub / UsdToEur
	readInput()
}

func readInput() {
	fmt.Print("Введите строку: ")
	var input string
	fmt.Scan(&input)
}

func calculate(amount float64, source, target string) float64 {
	return 1
}
