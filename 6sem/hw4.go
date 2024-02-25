// Перепишите задачи 1 и 2, используя пакет ioutil.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "toDoList.txt"
	
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Вот что в файле")
	fmt.Println(string(resultBytes))
	}
