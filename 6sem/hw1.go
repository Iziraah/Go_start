/* Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:
 2020-02-10 15:00:00 продам гараж.
При вводе слова exit программа завершает работу.
**/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer
	fmt.Println("Каков план действий на сегодня?")
	b.WriteString("Каков план действий на сегодня? \n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var answer string
		scanner.Scan()
		answer = scanner.Text()

		if answer == "exit" {
			fmt.Println("Выход из программы.")
			break
		}

		currentTime := time.Now()
		fmt.Println("Запись успешно добавлена")
		fmt.Fprintf(&b, "%s %s \n", currentTime.Format("2006-01-02 15:04:05"), answer)
	}

	fileName := "toDoList.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Вот что в логах")
	fmt.Println(string(resultBytes))
}
