package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

// создание структуры
type rectangle struct {
	length float64
	width  float64
}

func main() {
	var key int
	fmt.Println("\033[2J")
	fmt.Println("1 - Чётное/Нечётное число\n" +
		"2 - Положительное/Отрицательное/ноль\n" +
		"3 - for от 1 до 10\n" +
		"4 - Длина вводимой строки\n" +
		"5 - площадь прямоугольника\n" +
		"6 - Среднее знячение для 2-x целых чисел\n\n" +
		"для выхода введите 0")
	fmt.Println("\n\nВведите число:")
	fmt.Fscan(os.Stdin, &key)
	for key != 0 {
		switch key {
		case 1:
			fmt.Println("\033[2J")
			EvenOdd()
			fmt.Scanln()

		case 2:
			fmt.Println("\033[2J")
			fmt.Print("Введите целое число  >>> ")
			number := checkInt()
			PosNegZero(number)
			fmt.Scanln()

		case 3:
			fmt.Println("\033[2J")
			my_for()
			fmt.Scanln()

		case 4:
			str := ""
			fmt.Println("\033[2J")
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Строка >>> ")
			str, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			StrLen(str)
			fmt.Scanln()

		case 5:
			var inLength float64
			var inWidth float64
			fmt.Println("\033[2J")
			fmt.Print("Введите длину >>> ")
			fmt.Fscan(os.Stdin, &inLength)
			fmt.Print("Введите ширину >>> ")
			fmt.Fscan(os.Stdin, &inWidth)

			//создание экземляра структуры
			r := rectangle{length: inLength, width: inWidth}
			fmt.Printf("Площадь: %.2f\n", r.area())

			fmt.Scanln()

		case 6:
			fmt.Println("\033[2J")
			fmt.Print("Введите целое число 1 >>> ")
			number1 := checkInt()
			fmt.Print("Введите целое число 2 >>> ")
			number2 := checkInt()
			avg := avgInt(number1, number2)
			fmt.Println("Среднее арифметическое: ", avg)
			fmt.Scanln()
		}

		fmt.Println("\033[2J")
		fmt.Println("1 - Чётное/Нечётное число\n" +
			"2 - Положительное/Отрицательное/ноль\n" +
			"3 - for от 1 до 10\n" +
			"4 - Длина вводимой строки\n" +
			"5 - площадь прямоугольника\n" +
			"6 - Среднее знячение для 2-x целых чисел\n\n" +
			"для выхода введите 0")
		fmt.Println("\n\nВведите число:")
		fmt.Fscan(os.Stdin, &key)

	}
}

func EvenOdd() {
	fmt.Print("\n\nВведите целое число >>> ")
	value := checkInt()
	value = value % 2
	if value == 1 {
		fmt.Println("Число нечётное")
	} else {
		fmt.Println("Число чётное")
	}
}

func PosNegZero(number int) {

	if number < 0 {
		fmt.Println("Negative")
	} else if number == 0 {
		fmt.Println("Zero")
	} else if number > 0 {
		fmt.Println("Positive")
	}
}

func my_for() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("")
}

func StrLen(str string) {
	fmt.Println(str)
	runeCount := utf8.RuneCountInString(str)
	fmt.Println("Длина строки:", runeCount)
}

// метод для расчёта площади через структуру
func (r rectangle) area() float64 {
	return r.length * r.width
}

func avgInt(number1 int, number2 int) int {
	return (number1 + number2) / 2
}

// проверка что вводится целое число
func checkInt() int {
	var number int

	for {
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Ошибка: введите целое число.")
			// Очистка ввода, если возникла ошибка
			var input string
			fmt.Scan(&input) // Читаем неверный ввод, чтобы избежать бесконечного цикла
			continue
		}
		break // Ввод успешен, выходим из цикла
	}

	return number
}
