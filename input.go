package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Program required one of more arguments!")
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range os.Args[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		invalid = append(invalid, k)
	}
	fmt.Printf("Всего проверено: %d, Int: %d, Float: %d\n", total, nInts, nFloats)
	if len(invalid) >= total {
		fmt.Println("Too many invalid arguments!")
		return
	}
	var max, min float64
	for i := 1; i < len(os.Args); i++ {
		val, err := strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			fmt.Printf("Error in parse strToFloat: \"%s\"\n", os.Args[i])
		}
		if i == 1 {
			min = val
			max = val
			continue
		}
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Printf("Min: %.1f, Max: %.1f", min, max)
}
