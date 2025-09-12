//Работа с передачей в функции неограниченного кол-ва аргументов

package myPackage

func FindMin(nums ...int) int {
	min := nums[0]
	for _, i := range nums {
		if i < min {
			min = i
		}
	}
	return min
}

func FindMax(nums ...int) int {
	max := nums[0]
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max
}
