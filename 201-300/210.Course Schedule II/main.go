package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	inDegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for i := range adjacency {
		adjacency[i] = make([]int, 0)
	}

	for _, prerequisite := range prerequisites {
		course, require := prerequisite[0], prerequisite[1]
		inDegrees[course]++
		adjacency[require] = append(adjacency[require], course)
	}

	queue := make([]int, 0)
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, course)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		res = append(res, course)

		for _, nextCourse := range adjacency[course] {
			inDegrees[nextCourse]--
			if inDegrees[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}
	return res
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		{1, 0},
	}
	fmt.Println(findOrder(numCourses, prerequisites))

	numCourses = 4
	prerequisites = [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}
	fmt.Println(findOrder(numCourses, prerequisites))
}
