package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

	type Item struct {
		ID   int
		Name string
		Date time.Time
		Tags string
		Link string
	}

	items := make(map[int]Item)
	var lastID int = 0

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате url описание теги")
				continue OuterLoop
			}

			// Создаем новый объект Item
			newItem := Item{
				ID:   lastID + 1,
				Name: args[1],
				Date: time.Now(),
				Tags: args[2],
				Link: args[0],
			}

			// Добавляем новый объект в мапу с уникальным ID
			lastID++
			items[newItem.ID] = newItem

			fmt.Println("Ссылка добавлена успешно!")


			// Напишите свой код здесь
			
		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>
			fmt.Println("Список добавленных URL:")
			for id, item := range items {
				fmt.Printf("ID: %d\n", id)
				fmt.Printf("Имя: %s\n", item.Name)
				fmt.Printf("URL: %s\n", item.Link)
				fmt.Printf("Теги: %s\n", item.Tags)
				fmt.Printf("Дата: %v\n\n", item.Date)
			}
			// Напишите свой код здесь
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите ID ссылки, которую нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			id := strings.TrimSpace(text)

			// Преобразование строки в число
			deleteID := 0
			fmt.Sscanf(id, "%d", &deleteID)

			// Удаляем ссылку из мапы
			if _, exists := items[deleteID]; exists {
				delete(items, deleteID)
				fmt.Println("Ссылка удалена успешно!")
			} else {
				fmt.Println("Ссылка с указанным ID не найдена.")
			}
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
