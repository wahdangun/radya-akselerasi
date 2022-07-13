package main

import "fmt"

func Desending(arr []int) []int {
	var hasil []int
	for i := len(arr) - 1; i >= 0; i-- {
		hasil = append(hasil, arr[i])
	}
	return hasil
}

func removeDuplicates(arr []int) []int {
	var hasil []int
	for _, num := range arr {
		var exist bool
		for _, num2 := range hasil {
			if num == num2 {
				exist = true
			}
		}
		if exist == false {
			hasil = append(hasil, num)
		}
	}
	return hasil
}

func existInt(numsA, numsB []int) []int {
	var hasil []int
	for _, numb := range numsB {
		for _, numa := range numsA {
			if numa == numb {
				hasil = append(hasil, numa)
			}

		}

	}
	hasil = Desending(hasil)
	hasil = removeDuplicates(hasil)
	return hasil
}

func main() {
	fmt.Println(existInt([]int{1, 2, 2}, []int{1, 2, 4}))

}
