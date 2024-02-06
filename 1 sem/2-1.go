// На основе шаблона напишите программу, подсчитывающую сколько раз буква встречается в предложении, а также частоту встречаемости в процентах. То есть в предложении hello world результатом будет
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func countLetters(word string, letterCount map[rune]int) {
	for _, char := range word {
		// Игнорируем пробелы и другие символы, не являющиеся буквами
		if unicode.IsLetter(char) {
			letterCount[unicode.ToLower(char)]++
		}
	}
}

func main() {
	var text string
	fmt.Print("Введите предложение: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	
	words := strings.Split(text, "/")

	letterCount := make(map[rune]int)
	var countL = 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			countL++			
		}
	}

	for _, word := range words {
		// Подсчет букв в слове и добавление в словарь
		countLetters(word, letterCount)
	}
	
	// Вывод результатов
	fmt.Println("Количество повторений каждой буквы:")
	for letter, count := range letterCount {
		percent := float64(count) / float64(countL)
    	fmt.Printf("%c: %d %.1f\n", letter, count, percent)
	}
}