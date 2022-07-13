package main

import (
	"fmt"
)

func minus(arr []int) int {
	var negative bool
	var positive bool
	var validate bool
	var z int
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative = true
		}
		if arr[i] > 0 {
			positive = true
		}

	}
	if positive == false && negative == true {
		for {

			for x := 0; x < len(arr); x++ {
				if arr[x] == z && z != 0 {
					validate = true
				}
			}

			if validate == false && z != 0 {
				return z
			}
			validate = false
			z--

		}
	}
	if positive == true && negative == false {
		for {

			for x := 0; x < len(arr); x++ {
				if arr[x] == z && z != 0 {
					validate = true
				}
			}

			if validate == false && z != 0 {
				return z
			}
			validate = false
			z++

		}
	}

	return 0

}

func main() {
	arr := []int{2, 3, 1, 6}
	fmt.Println(minus(arr))
}
