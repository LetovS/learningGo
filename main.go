/*
Первая программа
на языке Go
*/
package main // определение пакета для текущего файла
import (
	"fmt" // подключение пакета fmt
	"strconv"
)

// определение функции main
func main() {

	var a string
	fmt.Println("Enter first value")
	fmt.Scanln(&a)

	var b string
	fmt.Println("Enter second value")
	fmt.Scanln(&b)

	fmt.Printf("You entered %s and %s\n", a, b)

	intA, errA := strconv.Atoi(a)
	if errA != nil {
		fmt.Println("Ошибка преобразования a:", errA)
		return
	}

	intB, errB := strconv.Atoi(b)
	if errB != nil {
		fmt.Println("Ошибка преобразования b:", errB)
		return
	}
	fmt.Println(intA + intB)

}
