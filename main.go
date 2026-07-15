package main

import (
	"fmt"
	"math"
)

func main(){
	var (
		num1 float32
		num2 float32
		oper string
	)
	
	fmt.Print("Masukkan Angka Pertama: ")
	fmt.Scanln(&num1)
	fmt.Print("Masukkan Angka Kedua: ")
	fmt.Scanln(&num2)
	fmt.Print("Masukkan Operator: ")
	fmt.Scanln(&oper)

	calculate(num1, num2, oper)
}

func calculate(num1 float32, num2 float32, oper string) float32 {
	var hasil float32
	
	switch oper {
		case "+":
			hasil = num1 + num2
			return hasil 
		case "-":
			hasil = num1 - num2
			return hasil 
		case "*":
			hasil = num1 * num2
			return hasil 
		case "/":
			hasil = num1 / num2
			return hasil 
		case "%":
			hasil = float32(math.Mod(float64(num1), float64(num2)))
			return hasil
		default:
			fmt.Println("Salah Operator Coba Lagi")
			return 1
	}
}