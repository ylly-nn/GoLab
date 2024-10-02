package mathutils

func Factorial(n int) (string, int) {
	if n < 0 {
		return "Ошибка: отрицательное число", 0
	} else if n == 0 {
		return "Функция успешно завершена c значением:", 1
	} else {
		fact := 1
		for i := 1; i <= n; i++ {
			fact *= i
		}
		return "Функция успешно завершена с значением:", fact
	}

}
