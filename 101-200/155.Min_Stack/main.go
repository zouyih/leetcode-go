package main

import "fmt"

type MinStack struct {
	nums    []int
	minNums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	nums := make([]int, 0)
	minNums := make([]int, 0)
	return MinStack{nums: nums, minNums: minNums}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	if len(this.minNums) == 0 || x <= this.minNums[len(this.minNums)-1] {
		this.minNums = append(this.minNums, x)
	}
}

func (this *MinStack) Pop() {
	popNum := this.Top()
	this.nums = this.nums[:len(this.nums)-1]
	if popNum == this.GetMin() {
		this.minNums = this.minNums[:len(this.minNums)-1]
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.minNums[len(this.minNums)-1]
}

func main() {
	ms := Constructor()
	ms.Push(1)
	ms.Push(2)
	fmt.Printf("top: %d, min: %d\n", ms.Top(), ms.GetMin())
	ms.Push(1)
	fmt.Printf("top: %d, min: %d\n", ms.Top(), ms.GetMin())
	ms.Pop()
	fmt.Printf("top: %d, min: %d\n", ms.Top(), ms.GetMin())
}
