package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	people := map[string]int{
		"Геворг": 59,
		"Марина": 25,
		"Игорь":  37,
	}
	var key int
	fmt.Println("\033[2J")
	fmt.Println("1 - Добавление нового человека\n" +
		"2 - Средний возраст людей\n" +
		"3 - Удаление записи по имени\n" +
		"4 - Вывод строки в верхнем регистре\n" +
		"5 - Сумма введённых чисел\n" +
		"6 - Переворот массива чисел\n\n" +
		"Для выхода введите 0")
	fmt.Println("\n\nВведите номер задачи:")
	fmt.Fscan(os.Stdin, &key)

	for key != 0 {
		switch key {
		case 1:
			fmt.Println("\033[2J")
			fmt.Println("Карта до добавления")
			fmt.Println(people)
			mapAdd(people)
			fmt.Println("Карта после добавления")
			fmt.Println(people)
			fmt.Scanln()

		case 2:
			fmt.Println("\033[2J")
			fmt.Println(avgAge(people))
			fmt.Scanln()

		case 3:
			fmt.Println("\033[2J")
			fmt.Println(people)
			name := ""
			fmt.Print("Введите имя >>> ")
			fmt.Fscan(os.Stdin, &name)
			delName(people, name)
			fmt.Println(people)
			fmt.Scanln()

		case 4:
			fmt.Println("\033[2J")
			fmt.Println("Введите строку:")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			upperCase(input)
			fmt.Scanln()

		case 5:
			fmt.Println("\033[2J")
			SumInput()
			fmt.Scanln()
		case 6:
			fmt.Println("\033[2J")
			IntArray()
			fmt.Scanln()
		}

		fmt.Println("\033[2J")
		fmt.Println("1 - Добавление нового человека\n" +
			"2 - Средний возраст людей\n" +
			"3 - Удаление записи по имени\n" +
			"4 - Вывод строки в верхнем регистре\n" +
			"5 - Сумма введённых чисел\n" +
			"6 - Переворот массива чисел\n\n" +
			"Для выхода введите 0")
		fmt.Println("\n\nВведите номер задачи:")
		fmt.Fscan(os.Stdin, &key)
	}

}

func mapAdd(mp map[string]int) {
	mp["Ольга"] = 34
}

func avgAge(mp map[string]int) float64 {
	if len(mp) == 0 {
		return 0
	}
	Age := 0
	for _, age := range mp {
		Age += age
	}
	return float64(Age) / float64(len(mp))
}

func delName(mp map[string]int, name string) {
	_, exists := mp[name]
	if exists {
		delete(mp, name)
	} else {
		fmt.Printf("Запись с именем %s не найдена.\n", name)
	}
}

func upperCase(str string) {
	upperStr := strings.ToUpper(str)
	fmt.Println("Строка в верхнем регистре:", upperStr)
}

func SumInput() {
	fmt.Println("Введите несколько чисел, разделённых пробелами:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")

	sum := 0

	for _, num := range numbers {
		// Преобразуем строку в целое число
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Ошибка преобразования числа:", num)
			return
		}

		sum += n
	}

	fmt.Println("Сумма введённых чисел:", sum)

}

// Обратный порядок
func IntArray() {
	fmt.Println("Введите несколько целых чисел, разделённых пробелами:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	strNumbers := strings.Split(input, " ")

	intArray := make([]int, 0)

	// Преобразуем каждую строку в целое число и добавляем в массив
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка преобразования числа:", str)
			return
		}
		intArray = append(intArray, num)
	}

	//Обратный порядок
	fmt.Println("Числа в обратном порядке:")
	for i := len(intArray) - 1; i >= 0; i-- {
		fmt.Print(intArray[i], " ")
	}
	fmt.Println()
}
