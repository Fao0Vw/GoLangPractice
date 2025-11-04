package main

import "fmt"

func main() {
	//Классический for
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
	// Идиоматический для GO
	i := 0
	for k := true; k; k = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	//Имитация цикла while в GO
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	//Использование range для перебора элементов среза или массиваW
	aSlice := []int{-1, 2, 1, -1, 2, -4}
	for i, v := range aSlice {
		fmt.Printf("Index: %d Value: %d\n", i, v)
	}
}
