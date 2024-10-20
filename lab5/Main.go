package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	height float32
	width  float32
}

type Shape interface {
	Area() float32
}

type Book struct {
	title  string
	author string
}

type Stringer interface {
	String() string
}

func main() {

	//Объявление элементов круга и прямоугольника
	person1 := Person{name: "Igor", age: 16}

	fmt.Println("Информация о человеке")
	person1.PrintInfo()

	fmt.Println()
	fmt.Println("Добавление года к возрасту")
	person1.Birthday()
	person1.PrintInfo()

	//Создание среза интерфейсов Shape
	fmt.Println()
	fmt.Println("ShapeArea")
	ShapeSlise := make([]Shape, 0)
	ShapeSlise = append(ShapeSlise, Circle{radius: 45.9}, Circle{radius: 16}, Rectangle{width: 4, height: 3}, Rectangle{width: 6.4, height: 7.9})
	fmt.Println(ShapeSlise)

	//Вывод Area для Shape
	PrintAreas(ShapeSlise)

	fmt.Println()
	fmt.Println("Информация о книгах")
	book := make([]Stringer, 0)
	book = append(book, Book{title: "Война и мир", author: "Толстой Л. Н"}, Book{title: "Преступление и наказание", author: "Достоевский Ф. М"})
	for i := 0; i < len(book); i++ {
		fmt.Println(Stringer(book[i]))
	}

}

func (p Person) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

// Метод birthday для увеличения возраста на 1 год
func (p *Person) Birthday() {
	p.age += 1
}

// Метод для вычисления площади круга
func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float32 {
	return r.width * r.height
}

// Функция для вывода площадей
func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

// Книги
func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.title, b.author)
}
