package main

import "fmt"

func canFinish(num int, preRequires [][]int) bool {
	inDegrees := make([]int, num)
	adjacency := make([][]int, num)
	for i := 0; i < num; i++ {
		adjacency[i] = make([]int, 0)
	}
	for _, preRequire := range preRequires {
		course, require := preRequire[0], preRequire[1]
		inDegrees[course]++
		adjacency[require] = append(adjacency[require], course)
	}

	queue := make([]int, 0)
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, course)
		}
	}

	numCourse := 0
	for len(queue) > 0 {
		numCourse++

		preCourse := queue[0]
		queue = queue[1:]

		for _, course := range adjacency[preCourse] {
			inDegrees[course]--
			if inDegrees[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	return numCourse == num
}

func canFinish1(num int, preRequires [][]int) bool {
	flags := make([]int, num)
	adjacency := make([][]int, num)
	for i := 0; i < num; i++ {
		adjacency[i] = make([]int, 0)
	}

	for _, preRequire := range preRequires {
		course, require := preRequire[0], preRequire[1]
		adjacency[require] = append(adjacency[require], course)
	}

	for i := 0; i < num; i++ {
		if !dfs(i, adjacency, flags) {
			return false
		}
	}

	return true
}

func dfs(i int, adjacency [][]int, flags []int) bool {
	if flags[i] == 1 {
		return true
	}
	if flags[i] == -1 {
		return false
	}

	flags[i] = -1
	for _, course := range adjacency[i] {
		if !dfs(course, adjacency, flags) {
			return false
		}
	}
	flags[i] = 1

	return true
}

func main() {
	num := 2
	preRequires := [][]int{
		{1, 0},
	}
	fmt.Println(canFinish(num, preRequires))

	num = 2
	preRequires = [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(canFinish(num, preRequires))
}
