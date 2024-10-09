package main

import (
	"bufio"
	"fmt"
	mathutils "lab3/mathutils"
	"lab3/stringutils"
	"math/rand"
	"os"
)

func main() {
	var key int
	arr := createRandomArray()
	fmt.Println("\033[2J")
	fmt.Println("1 - Факториал числа\n" +
		"2 - Переворот строки\n" +
		"3 - Массив случайных чисел\n" +
		"4 - Работа со срезом (добавление и удаление элементов)\n" +
		"5 - Поиск самой длинной строки\n\n" +
		"Для выхода введите 0")
	fmt.Println("\n\nВведите номер задачи:")
	fmt.Fscan(os.Stdin, &key)

	for key != 0 {
		switch key {
		case 1:
			fmt.Println("\033[2J")
			num := 0
			fmt.Print("Введите число для расчёта факториала >>>")
			fmt.Fscan(os.Stdin, &num)
			fmt.Println(mathutils.Factorial(num))
			fmt.Scanln()

		case 2:
			fmt.Println("\033[2J")
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Введите строку >>> ")
			str, _ := reader.ReadString('\n')
			fmt.Println(stringutils.Str_Reverse(str))
			fmt.Scanln()

		case 3:
			fmt.Println("\033[2J")
			arr := createRandomArray()
			fmt.Println("Массив случайных чисел:", arr)
			fmt.Scanln()

		case 4:
			fmt.Println("\033[2J")
			modifySlice(arr)
			fmt.Scanln()

		case 5:
			reader := bufio.NewReader(os.Stdin)
			strings := make([]string, 0)
			for i := 0; i < 3; i++ {
				fmt.Println("Введите строку >>> ")
				st, _ := reader.ReadString('\n')
				strings = append(strings, st)
			}
			longest := findLongestString(strings)
			fmt.Println("Самая длинная строка:", longest)
			fmt.Scanln()
		}

		fmt.Println("\033[2J")
		fmt.Println("1 - Факториал числа\n" +
			"2 - Переворот строки\n" +
			"3 - Массив случайных чисел\n" +
			"4 - Работа со срезом (добавление и удаление элементов)\n" +
			"5 - Поиск самой длинной строки\n\n" +
			"Для выхода введите 0")
		fmt.Println("\n\nВведите номер задачи:")
		fmt.Fscan(os.Stdin, &key)
	}
}

// Функция для создания массива из 5 случайных целых чисел
func createRandomArray() [5]int {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) // Случайные числа от 0 до 99
	}
	return arr
}

// Функция для работы со срезом: добавление и удаление элементов
func modifySlice(arr [5]int) []int {
	slice := make([]int, 0) // Создание среза из массива
	for i := 0; i < len(arr); i++ {
		slice = append(slice, arr[i])
	}
	fmt.Println("Исходный срез:", slice)

	// Добавление случайного элемента в срез
	slice = append(slice, rand.Intn(100))
	fmt.Println("Срез после добавления элемента:", slice)

	// Удаление второго элемента (индекс 1)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Срез после удаления второго элемента:", slice)

	return slice
}

// Функция для нахождения самой длинной строки в срезе строк
func findLongestString(strings []string) string {
	longest := strings[0]
	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}
