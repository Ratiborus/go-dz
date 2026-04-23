package main

import "fmt"

func main() {
	const UsdToEur = 0.92
	const UsdToRub = 75.0
	const EurToRub = UsdToRub / UsdToEur
	fmt.Println(UsdToEur, UsdToRub, EurToRub)
}
