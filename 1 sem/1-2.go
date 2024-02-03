package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wordCount, letterCount, byteCount int
	var isWord bool

	for _, char := range text {
		byteCount++ 
		if unicode.IsSpace(char) {
			isWord = false
		} else if unicode.IsLetter(char) {
			letterCount++
			if !isWord {
				wordCount++
				isWord = true
			}
		}
	}

	fmt.Printf("Количество слов: %d\n", wordCount)
	fmt.Printf("Количество букв: %d\n", letterCount)
	fmt.Printf("Количество байт: %d\n", byteCount)
}
