package main

import (
	"fmt"
	"strconv"
)

func runeFactory(s string) []rune {

	chars := []rune(s)
	return chars

}

func reverse(r []rune) string {
	var hasil string
	for i := (len(r) - 1); i >= 0; i-- {
		hasil = hasil + string(r[i])
	}
	return hasil
}

func joinReversed(exclude []int, nums ...int) (int, error) {
	var excludeStr string
	var numsStr string
	var hasil string
	var counter int
	var final []rune
	//var validate bool
	for _, exclude := range exclude {
		excludeStr += fmt.Sprintf("%d", exclude)
	}
	for _, num := range nums {
		numsStr += fmt.Sprintf("%d", num)

	}
	runeNums := runeFactory(numsStr)
	runeExclude := runeFactory(excludeStr)
	for _, num := range runeNums {
		for x := 0; x < len(runeExclude); x++ {
			if num == runeExclude[x] {
				num = 0
				counter++
			}
		}
		if num > 0 {
			final = append(final, num)
		}
	}

	if counter == len(runeNums) {
		err := fmt.Errorf("Not found")
		return 0, err
	} else {
		hasil = reverse(final)
		angka, err := strconv.Atoi(hasil)
		return angka, err
	}

}

func main() {
	fmt.Println(joinReversed([]int{0, 1, 0}, 1000, 1010))
}
