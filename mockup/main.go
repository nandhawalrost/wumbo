package main

import "fmt"

//math operation mockup
func multiplyArray(arr []int) int {
	product := 1
	for _, num := range arr {
		product *= num
	}
	return product
}

func divideArray(arr []int) float64 {
	quotient := float64(arr[0])
	for _, num := range arr[1:] {
		quotient /= float64(num)
	}
	return quotient
}

func main() {
	arr := []int{1, 2, 3}
	multiply := multiplyArray(arr)
	fmt.Println(multiply)
	divide := divideArray(arr)
	fmt.Println(divide)
}
