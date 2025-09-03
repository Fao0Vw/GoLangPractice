package main

import (
	"fmt"
	"os"
)

func main() {
	var str, sep string
	for i := 0; i < len(os.Args); i++ {
		str += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(str)
}

// func main() {
// 	for index, arg := range os.Args[0:] {
// 		fmt.Printf("%d: %s\n", index, arg)
// 	}

// }

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		if input.Text() == "" {
// 			break
// 		}
// 		counts[input.Text()]++
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Println(n, line)
// 		}
// 	}
// }
