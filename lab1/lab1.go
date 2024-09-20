package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var key int
	fmt.Println("\033[2J")
	fmt.Println("1 - реальное время\n" +
		"2 - переменные различных типов\n" +
		"3 - арифметические операции\n" +
		"4 - Сумма разность для числа с плавающей точкой\n" +
		"5 - Среднее знячение для 3-х\n\n" +
		"для выхода введите 0")
	fmt.Println("\n\nВведите число:")
	fmt.Fscan(os.Stdin, &key)
	for key != 0 {
		switch key {
		case 1:
			fmt.Println("\033[2J")
			my_time()
			fmt.Scanln()

		case 2:
			fmt.Println("\033[2J")
			variable()
			fmt.Scanln()

		case 3:
			fmt.Println("\033[2J")
			arifm()
			fmt.Scanln()

		case 4:
			fmt.Println("\033[2J")
			SumDif()
			fmt.Scanln()

		case 5:
			fmt.Println("\033[2J")
			avg()
			fmt.Scanln()
		}
		fmt.Println("\033[2J")
		fmt.Println("1 - реальное время\n" +
			"2 - переменные различных типов\n" +
			"3 - арифметические операции\n" +
			"4 - Сумма разность для числа с плавающей точкой\n" +
			"5 - Среднее знячение для 3-х\n\n" +
			"для выхода введите 0")
		fmt.Println("\n\nВведите число:")
		fmt.Fscan(os.Stdin, &key)
	}
}

func my_time() {
	Now_time := time.Now()
	fmt.Println("Реальное время:")
	fmt.Println(Now_time)
}

func variable() {
	Int := 3
	Float := 4.55
	Str := "String"
	Bool := true

	fmt.Println("int: ", Int)
	fmt.Println("float64: ", Float)
	fmt.Println("string: ", Str)
	fmt.Println("bool: ", Bool)
}

func arifm() {
	a := 10
	b := 3
	fmt.Println("a = ", a, "; b = ", b, ";")
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", float64(a)/float64(b))
	fmt.Println("a //b = ", a/b)
	fmt.Println("a \\%b = ", a%b)

}

func SumDif() {
	var a float64
	var b float64
	fmt.Print("a >>> ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Print("b >>> ")
	fmt.Fscan(os.Stdin, &b)
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
}

func avg() {
	a := 0
	b := 0
	c := 0
	fmt.Print("a >>> ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Print("b >>> ")
	fmt.Fscan(os.Stdin, &b)
	fmt.Print("c >>> ")
	fmt.Fscan(os.Stdin, &c)
	fmt.Println("avg(a, b, c): ", float64(a+b+c)/3)

}
