package InsertionSort

import "testing"

func BenchmarkSelectionSort(b *testing.B) {
	array := []int{1, 22, 7, 19, 54, 222, 343, 45, 56, 23, 5, 57, 7896, 78, 234, 23, 246, 456879, 345, 213, 64, 234, 123, 3546, 234, 123, 4537, 234, 123, 346, 79, 780, 678, 567, 0, 234, 4357, 568, 234, 123, 57, 9, 808, 23, 28, 594, 4, 22234, 23, 4, 687, 0, 5}
	for i := 0; i < b.N; i++ {
		Demo(array)
	}
}
