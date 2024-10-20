package lab6_2

import (
	"fmt"
)

// Генерация чисел Фибоначчи и отправка их в канал
func GenerateFibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a // Отправляем число в канал
		a, b = b, a+b
	}
	close(ch) // Закрываем канал, чтобы другая горутина знала, что данных больше не будет
}

// Чтение чисел из канала и вывод на экран
func PrintFibonacci(ch chan int) {
	// Чтение из канала до тех пор, пока он не будет закрыт
	for num := range ch {
		fmt.Println(num) // Выводим число на экран
	}
}

func Run() {
	fibChan := make(chan int)

	// Запускаем горутину для генерации чисел Фибоначчи
	go GenerateFibonacci(10, fibChan)

	// Запускаем горутину для чтения и вывода чисел Фибоначчи
	PrintFibonacci(fibChan)
}
