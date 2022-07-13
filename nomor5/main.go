package main

import (
	"fmt"
	"math"
)

func multiPlications(nums []int, hasil chan int) {

	hasil <- nums[0] * nums[1]

}

func sum(nums []int) int {
	var hasil int
	data := make(chan int)
	var hasilF float64
	hasilF = math.Ceil(float64(len(nums)) / 2)

	for i := 0; i < len(nums); i = i + 2 {
		if i == len(nums)-1 {
			fmt.Println(nums[i], "*", nums[i])
			go multiPlications([]int{nums[i], nums[i]}, data)
		} else {
			fmt.Println(nums[i], "*", nums[i+1])
			go multiPlications([]int{nums[i], nums[i+1]}, data)
		}
	}
	i := 0
	for n := range data {
		hasil += n
		if i >= int(hasilF)-1 {
			close(data)
		}
		i++
	}

	return hasil
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6}))
}
