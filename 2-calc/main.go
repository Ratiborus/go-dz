package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const AVG string = "avg"
const SUM string = "sum"
const MED string = "med"

func main() {
	for {
		operation := readOperation()
		numbers := readNumbers()
		arr := toArray(numbers)
		res := calculate(operation, arr[:])
		fmt.Printf("Результат вычисления: %.2f\n", res)
		if !proceed() {
			break
		}
	}
}

func readOperation() string {
	for {
		fmt.Printf("Введите тип операции(AVG, SUM, MED): ")
		var operation string
		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Printf("Ошибка ввода %v введите AVG, SUM, MED\n", err)
			continue
		}
		if isIncorrect(operation) {
			fmt.Printf("Неизвестная команда %v, введите AVG, SUM, MED\n", operation)
			continue
		}
		return operation
	}

}

func isIncorrect(operation string) bool {
	return !(strings.EqualFold(AVG, operation) ||
		strings.EqualFold(SUM, operation) ||
		strings.EqualFold(MED, operation))
}

func readNumbers() string {
	for {
		fmt.Printf("Введите числовой набор: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка ввода %v введите числовой набор в формате (1, 2, 3)\n", err)
			continue
		}
		return line
	}
}

func toArray(nums string) []int {
	split := strings.Split(strings.TrimSpace(nums), ",")
	arr := make([]int, 0, len(split))
	for _, val := range split {
		val := strings.TrimSpace(val)
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Невозможно преобразовать символ %v в число ошибка: %v", val, err))
		}
		arr = append(arr, num)
	}
	return arr
}

func calculate(op string, arr []int) float64 {
	switch {
	case strings.EqualFold(AVG, op):
		return avg(arr)
	case strings.EqualFold(SUM, op):
		return sum(arr)
	case strings.EqualFold(MED, op):
		return med(arr)
	default:
		panic(fmt.Sprintf("Unknown operation %v", op))
	}
}

func avg(arr []int) float64 {
	return sum(arr) / float64(len(arr))
}

func sum(arr []int) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += float64(val)
	}
	return sum
}

func med(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	if n == 0 {
		return 0.0
	} else if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	} else {
		return float64(arr[n/2])
	}
}

func proceed() bool {
	for {
		fmt.Print("Вы хотите продолжить вычисления (y/n)? ")
		var res string
		_, err := fmt.Scan(&res)
		if err != nil {
			fmt.Printf("Неизвестная команда: %v\n", err)
			continue
		}
		return strings.EqualFold("y", res)
	}
}
