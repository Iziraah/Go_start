package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func main() {

	fmt.Println("Введите адрес")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	// filePth := os.Args[1]

	// var fileName, fileExt string
	// C:/Users/Mariia/Desktop/учеба/go/1 sem/1-3.go

	elements := strings.Split(text, "/")
	lastElement := elements[len(elements)-1]
	parts := strings.Split(lastElement, ".")
	fileName := parts[0]
	fileExt := parts[1]


	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)

}

