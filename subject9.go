package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// sub 1
	// var numbers []int
	// number := 0
	// for {
	// 	fmt.Scan(&number)
	// 	if number == 0 {
	// 		break
	// 	}
	// 	numbers = append(numbers, number)
	// }
	// fmt.Println(numbers)
	// summ := 0
	// for _, num := range numbers {
	// 	summ += num
	// }
	// fmt.Println(summ)

	// // sub 2

	// var values []int
	// var values2 []int
	// value := 0
	// for {
	// 	fmt.Scan(&value)

	// 	if value == 0 {
	// 		break
	// 	}
	// 	values = append(values, value)
	// }
	// fmt.Println(values)

	// for _, num := range values {
	// 	if num%2 == 0 {
	// 		values2 = append(values2, num)
	// 	}
	// }
	// fmt.Println(values2)
	// // sub 3

	// var data []int
	// data2 := 0
	// for {
	// 	fmt.Scan(&data2)

	// 	if data2 == 0 {
	// 		break
	// 	}
	// 	data = append(data, data2)
	// }
	// fmt.Println(data)

	// data = append(data[:2], data[2+1:]...)
	// fmt.Println(data)

	// sub 4

	// var temps []int
	// temp := 0
	// for {
	// 	fmt.Scan(&temp)
	// 	if temp == 0 {
	// 		break
	// 	}
	// 	temps = append(temps, temp)
	// }
	// fmt.Println(temps)
	// min := temps[0]
	// max := temps[0]
	// for _, num := range temps {
	// 	if num < min {
	// 		min = num
	// 	}
	// 	if num > max {
	// 		max = num
	// 	}
	// }
	// fmt.Println(min)
	// fmt.Println(max)

	// sub 5

	// var words []string
	// word := ""
	// for {
	// 	fmt.Scan(&word)
	// 	if word == "stop" {
	// 		break
	// 	}
	// 	words = append(words, word)
	// }
	// fmt.Println(words)

	// for i := len(words) - 1; i >= 0; i-- {
	// 	fmt.Print(words[i], " ")

	// }

	// sub 6

	var nums []int
	num := 0
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		nums = append(nums, num)
	}
	fmt.Println(nums)
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	sort.Ints(nums)

	fmt.Println(reflect.DeepEqual(nums, nums2))

	// sub 7

	myWishList := []string{"aa", "bb", "cc"}
	friendsWishList := []string{"gg", "ff", "dd"}

	var resultList []string
	for _, registrationList := range friendsWishList {
		myWishList = append(myWishList, registrationList)
		resultList = myWishList
	}
	fmt.Println(resultList)

	// sub 8

	toyList := []string{"Car", "Doll", "Ball"}
	testToyList := []string{"Car", "Doll", "Ball"}
	testToyList[1] = "boat"
	fmt.Println(toyList)
	fmt.Println(testToyList)
}
