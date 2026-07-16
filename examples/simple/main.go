// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main(){
// 	var (
// 		num1 float32
// 		num2 float32
// 		oper string
// 		hasil float32
// 	)
	
// 	fmt.Print("Masukkan Angka Pertama: ")
// 	fmt.Scanln(&num1)
// 	fmt.Print("Masukkan Angka Kedua: ")
// 	fmt.Scanln(&num2)
// 	fmt.Print("Masukkan Operator: ")
// 	fmt.Scanln(&oper)

// 	hasil = calculate(num1, num2, oper)
// 	fmt.Printf("Hasil: %f\n", hasil)
// }

// func calculate(num1 float32, num2 float32, oper string) float32 {
// 	var hasil float32
	
// 	switch oper {
// 		case "+":
// 			hasil = num1 + num2
// 			return hasil 
// 		case "-":
// 			hasil = num1 - num2
// 			return hasil 
// 		case "*":
// 			hasil = num1 * num2
// 			return hasil 
// 		case "/":
// 			hasil = num1 / num2
// 			return hasil 
// 		case "%":
// 			hasil = float32(math.Mod(float64(num1), float64(num2)))
// 			return hasil
// 		default:
// 			fmt.Println("Salah Operator Coba Lagi")
// 			return 1
// 	}
// }
package main

import (
	"fmt"
	"log"

	"github.com/aexionn/belajar_go"
)

func main() {
	val, err := calculator.Calculate("(2.5 - 1.35) * 2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 2.3

	val, err = calculator.Calculate("-sin((-1+2.5)*pi)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 1

	val, err = calculator.Calculate("180*atan2(log(e), log10(10))/pi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 45
}