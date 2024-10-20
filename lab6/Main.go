package main

import (
	"fmt"
	lab6_1 "lab6/lab6_1"
	lab6_2 "lab6/lab6_2"
	lab6_3 "lab6/lab6_3"
	lab6_4 "lab6/lab6_4"
	lab6_5 "lab6/lab6_5"
	lab6_6 "lab6/lab6_6"
	"os"
)

func main() {
	var key int
	fmt.Println("\033[2J")
	fmt.Println("1 - 3 Функции в разных потоках\n" +
		"2 - Отправка сичтывание\n" +
		"3 - Select\n" +
		"4 - Мьютексы\n" +
		"5 - Калькулятор\n" +
		"6 - Пул воркеров\n\n" +
		"для выхода введите 0")
	fmt.Println("\n\nВведите число:")
	fmt.Fscan(os.Stdin, &key)
	for key != 0 {
		switch key {
		case 1:
			fmt.Println("\033[2J")
			fmt.Println("3 Функции в разных потоках")
			lab6_1.Run()
			fmt.Print("Для продолжения нажмите Enter >>>")
			fmt.Scanln()

		case 2:
			fmt.Println("\033[2J")
			fmt.Println("Отправка Считывание")
			lab6_2.Run()
			fmt.Print("Для продолжения нажмите Enter >>>")
			fmt.Scanln()
		case 3:
			fmt.Println("\033[2J")
			fmt.Println("Select")
			lab6_3.Run()
			fmt.Print("Для продолжения нажмите Enter >>>")
			fmt.Scanln()
		case 4:
			fmt.Println("\033[2J")
			fmt.Println("Мьютексы")
			lab6_4.Run()
			fmt.Print("Для продолжения нажмите Enter >>>")
			fmt.Scanln()
		case 5:
			fmt.Println("\033[2J")
			fmt.Println("Калькулятор")
			lab6_5.Run()
			fmt.Print("Для продолжения нажмите Enter >>>")
			fmt.Scanln()
		case 6:
			fmt.Println("\033[2J")
			fmt.Println("Пул воркеров")
			lab6_6.Run()
			fmt.Print("Для продолжения нажмите Enter >>>")
			fmt.Scanln()
		}

		fmt.Println("\033[2J")
		fmt.Println("1 - 3 Функции в разных потоках\n" +
			"2 - Отправка сичтывание\n" +
			"3 - Select\n" +
			"4 - Мьютексы\n" +
			"5 - Калькулято\n" +
			"6 - Пул воркеров\n\n" +
			"для выхода введите 0")
		fmt.Println("\n\nВведите число:")
		fmt.Fscan(os.Stdin, &key)
	}

}
