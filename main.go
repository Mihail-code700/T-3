package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Введите первое число:")
	fmt.Scanln(&num1)

	fmt.Println("Введите оператор (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	default:
		fmt.Println("Ошибка: неверный оператор!")
	}
}