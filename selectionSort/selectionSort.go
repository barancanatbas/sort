package selectionSort

import "fmt"

func Demo(array []int) {
	//array := []int{1, 22, 7, 19, 54, 222, 5}

	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		temp := array[i]
		array[i] = array[minIndex]
		array[minIndex] = temp
	}
}

type Node struct {
	Value int
	Next  *Node
}

func LinkedList(value int) *Node {
	return &Node{Value: value, Next: nil}
}

func (n *Node) Push(value int) {
	temp := n.FindLast()
	fmt.Println("last value ", temp.Value)
	node := LinkedList(value)
	temp.Next = node
}

func (n *Node) FindLast() *Node {
	temp := n
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	return temp
}

func (n *Node) List() {
	temp := n
	for {
		fmt.Println(temp.Value)
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
}

func (n *Node) Length() int {
	temp := n
	count := 1
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		count++
	}
	return count
}

func (n *Node) Get(index int) *Node {
	temp := n
	count := 0
	for {
		if count == index {
			return temp
		}
		temp = temp.Next
		count++
		if temp.Next == nil {
			break
		}
	}

	return temp
}

func (n *Node) Set(index, value int) {
	// temp := n
	// count := 0
	// for {
	// 	if count == index {
	// 		break
	// 	}
	// 	temp = temp.Next
	// 	count++
	// 	if temp.Next == nil {
	// 		break
	// 	}
	// }

	// temp.Value = value
}

// func (n *Node) SelectionSort() {
// 	length := n.Length()
// 	// 12 -> 13 -> 24 -> 1 -> 6 -> 16 -> 8
// 	// 1 -> 6
// 	for i := 0; i < length; i++ {
// 		iItem := n.Get(i)
// 		tempJ := iItem
// 		for j := i + 1; j < length; j++ {
// 			jItem := n.Get(j)
// 			if tempJ.Value < jItem.Value {
// 				tempJ = jItem
// 			}
// 		}
// 		tempNext := tempJ.Next
// 		tempJ.Next = iItem.Next
// 		n = tempJ
// 		iItem.Next = tempJ.Next
// 	}
// }
