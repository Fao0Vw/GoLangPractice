// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Please provide a command line argument")
// 		return
// 	}
// 	argument := os.Args[1]
// 	//switch можно использовать в двух сценариях:
// 	// 1) Производить вычисление выражения переменной в начале цикла
// 	// 2) Вычислять выражение внутри каждого case отдельно
// 	switch argument {
// 	case "0":
// 		fmt.Println("Zero")
// 	case "1":
// 		fmt.Println("One")
// 	case "2", "3", "4":
// 		fmt.Println("Two, three or four")
// 	default:
// 		fmt.Println("Value: ", argument)
// 	}

// 	value, err := strconv.Atoi(argument)
// 	if err != nil {
// 		fmt.Println("Cannot convert ASCII to integer: ", argument)
// 		return
// 	}

// 	switch {
// 	case value == 0:
// 		fmt.Println("0")
// 	case value > 0:
// 		fmt.Println("Positive Integer")
// 	case value < 0:
// 		fmt.Println("Negative Integer")
// 	default:
// 		fmt.Println("This should not happen: ", value)
// 	}
// }
