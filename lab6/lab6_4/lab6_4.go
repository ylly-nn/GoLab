package lab6_4

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex // мьютекс для синхронизации
)

func increment(wg *sync.WaitGroup, useMutex bool) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		if useMutex {
			mutex.Lock()   // блокируем доступ
			counter++      // увеличиваем счетчик
			mutex.Unlock() // разблокируем доступ
		} else {
			// если не используем мьютекс
			counter++
		}
	}
}

func Run() {
	var wg sync.WaitGroup

	// Пример с использованием мьютексов
	counter = 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go increment(&wg, true) // запускаем горутины с мьютексами
	}
	wg.Wait()
	fmt.Println("С использованием мьютексов:", counter)

	// Пример без использования мьютексов
	counter = 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go increment(&wg, false) // запускаем горутины без мьютексов
	}
	wg.Wait()
	fmt.Println("Без использования мьютексов:", counter)
}
