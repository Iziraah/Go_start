// Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
package main

import (
	"fmt"
	"os"
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
	arr1 := make([]int, 4)
	arr2 := make([]int, 5)

	// Вводим значения с консоли и сразу преобразуем их в целые числа
	fmt.Print("Введите четыре целых числа через пробел для первого массива: ")

	for i := 0; i < len(arr1); i++ {
		fmt.Scanf("%d", &arr1[i])
	}
	fmt.Print("Введите пять целых чисел через пробел для второго массива: ")
	for j := 0; j < len(arr2); j++ {
		fmt.Fscan(os.Stdin,  &arr2[j])
	}

	fmt.Println("Массив #1:", arr1)
	sort(arr1)
	fmt.Println("Отсортированный массив #1:", arr1)
	fmt.Println("Массив #2:", arr2)
	sort(arr2)
	fmt.Println("Отсортированный массив #2:", arr2)

	mergedArray := append(arr1, arr2...)
	fmt.Println("Объединенный массив:", mergedArray)
	sort(mergedArray)
	fmt.Println("Отсортированный объединенный массив:", mergedArray)

}