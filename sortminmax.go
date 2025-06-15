/* Написать программу на ЯП Go. Условие: На вход программе подаются 5 натуральных чисел.
Задача программы - считать эти числа, отсортировать по убыванию и вывести результат сортировки на экран.
Также вывести на экран другие параметры считанной последовательности в формате:
"Самое большое число: {число}",
"Самое маленькое число: {число}",
"Среднее арифметическое: {число}".

Пример работы программы:
выход: 5 2 4 3 1
выход:
"Отсортированые элементы: 5 4 3 2 1
Самое большое число: 5
Самое маленькое число: 1
Среднее-арифметическое: 3
"

Выложить программу на github и отправить ссылку для просмотра или отправить файл программы.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2, num3, num4, num5 float64

	fmt.Println("Введите первое число: ")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число: ")
	fmt.Scanln(&num2)

	fmt.Println("Введите третье число: ")
	fmt.Scanln(&num3)

	fmt.Println("Введите четвертое число: ")
	fmt.Scanln(&num4)

	fmt.Println("Введите пятое число: ")
	fmt.Scanln(&num5)

	a, b, c, d, e := num1, num2, num3, num4, num5

	for i := 0; i < 5; i++ {
		if a < b {
			a, b = b, a
		}
		if b < c {
			b, c = c, b
		}
		if c < d {
			c, d = d, c
		}
		if d < e {
			d, e = e, d
		}
	}

	fmt.Println("Отсортированные элементы (по убыванию):", a, b, c, d, e)

	maxNum1 := math.Max(num1, num2)
	maxNum2 := math.Max(maxNum1, num3)
	maxNum3 := math.Max(maxNum2, num4)
	maxNum4 := math.Max(maxNum3, num5)
	fmt.Println("Самое большое число: ", maxNum4)

	minNum1 := math.Min(num1, num2)
	minNum2 := math.Min(minNum1, num3)
	minNum3 := math.Min(minNum2, num4)
	minNum4 := math.Min(minNum3, num5)
	fmt.Println("Самое маленькое число: ", minNum4)

	midNum := (num1 + num2 + num3 + num4 + num5) / 5
	fmt.Println("Среднее арифметическое: ", midNum)

}
