package bubbleSort

import "fmt"

func Demo(array []int) {
	//array := []int{1, 22, 7, 19, 54, 222, 5}
	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	fmt.Println(array)
}
