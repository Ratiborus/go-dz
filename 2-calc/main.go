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

var operations = map[string]func([]int) float64{
	AVG: avg,
	SUM: sum,
	MED: med,
}

func main() {
	for {
		op := readOperation()
		numbers := read("Введите числовой набор: ", "числовой набор в формате (1, 2, 3)")
		arr := toArray(numbers)
		fn, ok := operations[normalize(op)]
		if !ok {
			fmt.Printf("Operation %v is not found\n", op)
			continue
		}
		res := fn(arr[:])
		fmt.Printf("Результат вычисления: %.2f\n", res)
		if !proceed() {
			break
		}
	}
}

func normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func readOperation() string {
	for {
		operation := read("Введите тип операции(AVG, SUM, MED): ", "AVG, SUM, MED")
		operation = normalize(operation)
		if isIncorrect(operation) {
			fmt.Printf("Неизвестная команда %v, введите AVG, SUM, MED\n", operation)
			continue
		}
		return operation
	}

}

func isIncorrect(operation string) bool {
	_, ok := operations[operation]
	return !ok
}

func read(inputMessage string, errMessage string) string {
	for {
		fmt.Print(inputMessage)
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка ввода %v введите %v\n", err, errMessage)
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
