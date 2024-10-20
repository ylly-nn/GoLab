package lab6_1

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Факториал числа n
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Генерация массива случайных чисел
func generateRandomNumbers(n, min, max int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}
	return numbers
}

// Сумма числового ряда
func sumOfSeries() (int, []int) {
	numbers := make([]int, 0)
	for i := 0; i <= 5; i++ {
		numbers = append(numbers, rand.Intn(10-1+1)+1)
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum, numbers
}

func Run() {
	// Используем WaitGroup для синхронизации завершения горутин
	var wg sync.WaitGroup
	wg.Add(3) // Так как у нас три функции, добавляем 3

	// Каналы для получения результатов

	// Горутине для расчёта факториала
	switch {
	}
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second) // Имитация задержки
		result := factorial(5)
		fmt.Println("Factorial (executed in goroutine):", result)
	}()

	// Горутине для генерации случайных чисел
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second) // Имитация задержки
		numbers := generateRandomNumbers(5, 1, 10)
		fmt.Println("Random Numbers (executed in goroutine):", numbers)
	}()

	// Горутине для вычисления суммы числового ряда
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Имитация задержки
		sum, numbers := sumOfSeries()
		fmt.Println("Sum of Series (executed in goroutine):", sum, numbers)
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()

}
