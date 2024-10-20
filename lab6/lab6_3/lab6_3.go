package lab6_3

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan<- int) {
	for i := 0; i < 15; i++ {
		num := rand.Intn(100) // Генерация случайного числа от 0 до 99
		ch <- num
		time.Sleep(time.Second) // Пауза между отправкой чисел
	}
	close(ch) // Закрытие канала после завершения генерации чисел
}

func checkEvenOdd(numbers <-chan int, results chan<- string) {
	for num := range numbers { // Обработка чисел из канала до его закрытия
		if num%2 == 0 {
			results <- fmt.Sprintf("Число %d чётное", num)
		} else {
			results <- fmt.Sprintf("Число %d нечётное", num)
		}
	}
	close(results) // Закрытие канала после завершения обработки чисел
}

func Run() {
	numberCh := make(chan int)
	resultCh := make(chan string)

	// Запуск горутин
	go generateNumbers(numberCh)
	go checkEvenOdd(numberCh, resultCh)

	// Основной цикл с использованием select
	for i := 0; i < 15; i++ {
		select {
		case num := <-numberCh:
			fmt.Printf("Сгенерировано число: %d\n", num)
		case result := <-resultCh: // Получаем результат чётности/нечётности
			fmt.Println(result)
		}
	}

	fmt.Println("Завершение программы.")
}
