package lab6_5

import (
	"fmt"
	"sync"
)

// Структура для запроса вычисления
type CalculationRequest struct {
	A, B     float64
	Operator string
	Result   chan float64
}

// Функция обработки запроса
func calculator(wg *sync.WaitGroup, requests <-chan CalculationRequest) {
	defer wg.Done()
	for req := range requests {
		var result float64
		switch req.Operator {
		case "+":
			result = req.A + req.B
		case "-":
			result = req.A - req.B
		case "*":
			result = req.A * req.B
		case "/":
			if req.B != 0 {
				result = req.A / req.B
			} else {
				fmt.Println("Ошибка: деление на ноль")
				req.Result <- 0
				continue
			}
		default:
			fmt.Println("Ошибка: неизвестная операция")
			req.Result <- 0
			continue
		}
		req.Result <- result
	}
}

func Run() {
	// Канал для отправки запросов на вычисление
	requests := make(chan CalculationRequest)

	// Используем группу ожидания для завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутин, каждая будет выполнять вычисления
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go calculator(&wg, requests)
	}

	// Функция для отправки запросов
	sendRequest := func(a, b float64, operator string) {
		resultChan := make(chan float64)
		request := CalculationRequest{A: a, B: b, Operator: operator, Result: resultChan}
		requests <- request
		result := <-resultChan
		fmt.Printf("Результат операции: %.2f %s %.2f = %.2f\n", a, operator, b, result)
	}

	// Отправляем несколько запросов
	go sendRequest(10, 5, "+")
	go sendRequest(20, 4, "-")
	go sendRequest(6, 3, "*")
	go sendRequest(15, 3, "/")
	go sendRequest(10, 0, "/") // Деление на ноль

	// Закрываем канал после отправки всех запросов
	go func() {
		wg.Wait()
		close(requests)
	}()

	// Небольшая пауза, чтобы все запросы обработались (можно заменить на wg.Wait() в main)
	fmt.Scanln()
}
