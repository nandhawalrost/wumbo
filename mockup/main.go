package main

import (
	"fmt"
	"regexp"
	"strings"
)

// math operation mockup
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

func extractNumber(str string) []string {
	trim_operator := regexp.MustCompile("[*/+-]") // contains kalibataku
	nums := trim_operator.ReplaceAllString(str, "-")

	array := strings.Split(nums, "-")

	return array
}

func main() {
	arr := []int{1, 2, 3}
	multiply := multiplyArray(arr)
	fmt.Println(multiply)
	divide := divideArray(arr)
	fmt.Println(divide)

	fmt.Println(extractNumber("1+22*3"))
	/*
		todo step math operation:
		contoh: 1 + 3 * 4
		-cari prioritasnya, misal kali dulu baru tambah
		-untuk bagi dan kali (prioritas) gunakan
		 function multiplyArray dan divideArray lalu dikurang atau ditambahkan,
		 gunakan sorting untuk menentukan prioritas
		-gunakan paralelisme pada operasi prioritas (kali dan bagi),
		 gunakan serialisme pada operasi non-prioritas (tambah dan kurang)
	*/
}
