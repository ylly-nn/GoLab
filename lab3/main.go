package main

import (
	"bufio"
	"fmt"
	mathutils "lab3/mathutils"
	"lab3/stringutils"
	"os"
)

func main() {
	num := 0
	fmt.Print("Введите число для расчёта факториала >>>")
	fmt.Fscan(os.Stdin, &num)
	fmt.Println(mathutils.Factorial(num))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку >>> ")
	str, _ := reader.ReadString('\n')
	fmt.Println(stringutils.Str_Reverse(str))
}
