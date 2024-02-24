// Отсортируйте массив длиной шесть пузырьком.

// Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
package main

import (
	"fmt"

)

func sort (input[]int) {

	n:= len(input)

	for i:=0; i< n-1; i++ {
		for j:=0; j< n-1-i; j++ {
			if input[j] > input[j+1] {
                input[j], input[j+1] = input[j+1], input[j]
            } 
		}

	}
}

func main() {
	arr := make([]int, 6)

	// Вводим значения с консоли и сразу преобразуем их в целые числа
	fmt.Print("Введите шесть целых числе через пробел для массива: ")

	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%d", &arr[i])
	}
	
	fmt.Println("Массив:", arr)
	sort(arr)
	fmt.Println("Отсортированный массив #1:", arr)
	

}