/*
Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.
**/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("\nПрограмма завершает выполнение.")
				close(ch)
				return
			default:
				var input string
				fmt.Print("Введите число: ")
				fmt.Scanln(&input)

				num, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Ошибка ввода числа:", err)
					continue
				}

				square := num * num
				fmt.Printf("%d^2 = %d\n", num, square)
			}
		}
	}()

	wg.Wait()
}
