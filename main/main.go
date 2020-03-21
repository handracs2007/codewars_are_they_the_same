package main

import "fmt"

func Comp(array1 []int, array2 []int) bool {
	if array1 != nil && array2 != nil && len(array1) == len(array2) {
		var array2Map = make(map[int]int)

		for _, val := range array2 {
			array2Map[val] = array2Map[val] + 1
		}

		for _, val := range array1 {
			var _, ok = array2Map[val*val]

			if !ok {
				return false
			} else {
				array2Map[val*val] = array2Map[val*val] - 1
			}
		}

		for _, val := range array2Map {
			if val != 0 {
				return false
			}
		}

		return true
	}

	return false
}

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))

	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))

	a1 = nil
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))
}
