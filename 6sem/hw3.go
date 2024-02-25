/* Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
Рекомендация:
Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.
**/

package main

import (
	"fmt"
	"os"
)

func main() {

	text := "This file you can only read"
	fileName := "onlyRead.txt"

	file, err := os.Create(fileName)
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
    }
    defer file.Close() 
    file.WriteString(text)     
    fmt.Println("Done.")

	if err = os.Chmod(fileName, 0400); err != nil {
		fmt.Println("Error changing file permissions:", err)
	}

	// Попытка добавить текст после установки разрешений
	_, err = file.WriteString("Test")
	if err != nil {
		fmt.Println("Error writing to read-only file:", err)
	}
}