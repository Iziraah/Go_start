/* Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
-Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
-Произведение: следующая горутина умножает квадрат числа на 2.
-При вводе «стоп» выполнение программы останавливается.**/

package main

import (
	"fmt"
	"sync"
	"strconv"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			num := <-ch
			if num < 0 {
				fmt.Println("Горутина 1 завершает выполнение.")
				return
			}
			square := num * num
			fmt.Printf("Горутина 1: %d^2 = %d\n", num, square)

			ch <- square
		}
	}()

	go func() {
		defer wg.Done()
		for {
			num := <-ch
			if num < 0 {
				fmt.Println("Горутина 2 завершает выполнение.")
				return
			}
			result := num * 2
			fmt.Printf("Горутина 2: %d * 2 = %d\n", num, result)
		}
	}()

	fmt.Print("Введите число (или 'стоп' для завершения): ")
	for {
		var input string
		fmt.Scanln(&input)

		if input == "стоп" {
			ch <- -1
			ch <- -1
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка ввода числа:", err)
			continue
		}

		ch <- num
	}

	wg.Wait()
}
