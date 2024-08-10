package main

import (
	"fmt"
	"strconv"
)

func firstFunc() {
	var input string
	fmt.Println("Enter something...")
	fmt.Scanln(&input)

	parseIntInput, error := strconv.Atoi(input)

	if error != nil {
		fmt.Println("Ошибка преобразования a:", error)
		return
	}
	fmt.Printf("You enterd number: %d \n", parseIntInput)
}

func ciclies(variant int) {

	switch variant {
	case 1:
		for i := 1; i < 10; i++ {
			fmt.Println(i * i)
		}
	case 2:
		fmt.Println("Inner cicle")
		for i := 1; i < 10; i++ {
			for j := 1; j < 10; j++ {
				fmt.Print(i*j, "\t")
			}
			fmt.Println()
		}
	default:
		fmt.Println("I don't know what I should do")
	}
}
