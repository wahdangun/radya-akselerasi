package main

import "fmt"

func intGrouping(num int) map[int]int {
	var hasil map[int]int
	hasil = make(map[int]int)
	for num > 0 {
		hasil[num%10]++
		num = num / 10
	}
	return hasil

}

func main() {

	fmt.Println(intGrouping(1112244))

}
