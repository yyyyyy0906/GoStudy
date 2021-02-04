/*
 * Slice(切片) -> 其本质是一个动态数组的封装，类似于C++的std::array和std::vector
*/
package main

import "fmt"

func BaseOperator() {
	// define:
	var sliceArray = []int{2, 0, 2, 1}
	for index, value := range sliceArray {
		fmt.Printf("current slice index = %v, value = %v \n", index, value)
	}
}
