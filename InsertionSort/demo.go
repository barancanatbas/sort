package InsertionSort

func Demo(array []int) []int {
	//array := []int{1, 22, 7, 19, 54, 222, 5}
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				temp := array[j]
				array[j] = array[j-1]
				array[j-1] = temp
			}
		}
	}

	return array
}
