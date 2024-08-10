/*
Первая программа
на языке Go
*/
package main // определение пакета для текущего файла
import (
	"fmt"
	"strconv"
)

// подключение пакета fmt

// определение функции main
func main() {

	//newFunction()
	var deleg func(int, int) int = test

	var res = delegate(10, 25, deleg)
	fmt.Println(res)
}

func newFunction() {
	firstFunc()
	var variant string
	fmt.Println("Enter variants: 1 - one cicle, 2 - inner cicle")
	fmt.Scanln(&variant)

	value, err := strconv.Atoi(variant)

	if err != nil {
		fmt.Printf("You enter incorrect value")
	}
	ciclies(value)

	result := Add(value, value)
	fmt.Println(result)

	prev, curr, next := Tuple(value)

	fmt.Printf("Result %d %d %d", prev, curr, next)
}

func Add(a int, b int) int {
	return a + b
}

func Tuple(value int) (previous int, current int, next int) {
	current = value
	previous = value - 1
	next = value + 1
	return
}

func test(x, y int) int {
	return x + y
}

func delegate(first, second int, del func(int, int) int) (sum int) {
	sum = del(first, second)
	return
}
