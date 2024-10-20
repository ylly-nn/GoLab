package lab6_6

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Воркер, который получает задачи через канал tasks и отправляет результаты в канал results
func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Здесь воркер просто отправляет строку без изменений
		fmt.Printf("Worker %d processed: %s\n", id, task)
		results <- task
	}
}

// Функция реверсирования порядка строк
func reverseStrings(lines []string) []string {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}

func Run() {
	// Ввод файла с задачами
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Считывание строк из файла
	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Реверсируем порядок строк
	tasks = reverseStrings(tasks)

	// Количество воркеров задается пользователем
	var numWorkers int
	fmt.Print("Enter the number of workers: ")
	fmt.Scanf("%d", &numWorkers)

	// Создание каналов для задач и результатов
	taskChan := make(chan string, len(tasks))
	resultChan := make(chan string, len(tasks))

	// WaitGroup для ожидания завершения работы всех воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskChan, resultChan, &wg)
	}

	// Отправляем задачи (уже реверсированные) в канал
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	// Ожидаем завершения работы всех воркеров
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Запись результатов в файл output.txt
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Получаем результаты и выводим их в файл
	for result := range resultChan {
		outputFile.WriteString(result + "\n")
	}

	fmt.Println("Processing complete. Results saved to output.txt")
}
