package main

import (
	"unicode"
	"strings"
)

/* 
Релизовать функцию, которая принимает на вход slice из 
чисел и возвращает slice в котором удалены все четные числа. 
Имя функции: RemoveEven
Пример использования:
input := []int{0, 3, 2, 5}
result := RemoveEven(input)
fmt.Println(result) // Должно напечататься [3 5]
*/

func RemoveEven(numbers []int) (without_even []int) {
	for _, value := range(numbers) {
		if value % 2 == 1 {
			without_even = append(without_even, value)
		}
	}
	return 
}

/*
Написать генератор, который будет принимать число и при вызове 
возвращать очередную степень этого числа. Имя генератора: PowerGenerator
Пример использования:

gen := PowerGenerator(2)
fmt.Println(gen()) // Должно напечатать 2
fmt.Println(gen()) // Должно напечатать 4
fmt.Println(gen()) // Должно напечатать 8
*/

func PowerGenerator(base int) (func() int) {
	power := 1
	return func() (res int) {
		res = 1
		for i := 0; i < power; i++ {
			res *= base
		}
		power++
		return
 	}
}


/*Реализуйте функцию, которая подсчитывает число различных слов в тексте (строчка). 
“Словом” считается непустая последовательность букв. 
Слова, отличающиеся только регистром символов, считаются одинаковыми.
Для решения этой задачи вам могут быть полезны функции из модуля unicode.
Имя функции: DifferentWordsCount
Пример использования:

fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
// Должно напечатать 2
*/

func DifferentWordsCount(text string) int {
	words := make(map[string]bool)
	word := ""
	for _, char := range text {
		if unicode.IsLetter(char) {
			word = word + string(char)
		} else {
			word = strings.ToLower(word)
			if _, ok := words[word]; !ok {
				if (word != "") {
					words[word] = true
				}
			}
			word = ""
		}
	}
	if _, ok := words[word]; !ok {
		if (word != "") {
			words[word] = true
		}
	}
	return len(words)
}

// func main() {
// 	fmt.Println(DifferentWordsCount("Hello, world!gHELLO  wOrlD...12"))
// }

