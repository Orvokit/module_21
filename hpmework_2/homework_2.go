/*Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).
Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.*/

package main

import (
	"fmt"
)

func F(A func(x, y int) int) string {
	defer fmt.Println(A(2, 3), "вывод функции через defer")
	return "..."
}

func main() {
	fmt.Println(F(func(x, y int) int { return x * y }), "вывод функции 1")
	fmt.Println(F(func(x, y int) int { return x + y }), "вывод функции 2")
	fmt.Println(F(func(x, y int) int { return x - y }), "вывод функции 3")
}
