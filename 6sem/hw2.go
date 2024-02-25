/*
Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
Рекомендация:
Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.
**/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "toDoList.txt"
	
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := make([]byte,64)

	for{
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}