// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	arr := []int32{7, 69, 2, 221, 8974}
	//	min, max := 0, 0
	min := 0

	// for min
	for index := 1; index < (len(arr)); index++ {
		min += int(arr[index])
		fmt.Println(int(arr[index]) + min)
	}
	fmt.Println(min)

	// for max
	/*for index := 2; index < (len(arr) + 1); index++ {
		max = int(arr[index]) + max
	}

	fmt.Println(min, max)*/
}
