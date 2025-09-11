// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func incr(p *int) {
// 	*p++
// }

// type Parent struct{}

// func (c *Parent) Print() {
// 	fmt.Println("parent")
// }

// type Child struct {
// 	Parent
// }

// func (p *Child) Print() {
// 	fmt.Println("child")
// }

// func main() {
// 	var x Child

// 	x.Print()
// 	fvar := 1.1
// 	fmt.Println(reflect.TypeOf(fvar))
// 	var bvar byte = 105
// 	var ivar int = 105
// 	fmt.Printf("%c - %c", bvar, ivar)
// }

// func main() {
// 	var str, sep, val = "", "", 0
// 	var arr [3]int = [3]int{1, 2, 3}
// 	arr[0], arr[1], arr[2] = 2, 4, 6
// 	val = arr[2]
// 	arrCurs := &arr
// 	varCurs := &val
// 	fmt.Println(arrCurs)
// 	fmt.Println(varCurs)
// 	*varCurs = 7
// 	fmt.Println(*varCurs)
// 	incr(varCurs)
// 	fmt.Println(*varCurs)
// 	varCurs = nil
// 	fmt.Println(varCurs, "\n", varCurs == nil)

// 	fmt.Println(val)
// 	for i := 1; i < len(os.Args); i++ {
// 		str += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(str)
// 	fmt.Println(*&str, str)
// 	for i := 0; i < 10; i++ {
// 		str := new(string)
// 		fmt.Println(str)
// 	}
// }

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