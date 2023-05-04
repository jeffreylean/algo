package algo

import (
	"sort"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

// BFS method which will eventually list out all the delayed project
func BFS(delayProjects []string, dependencies [][]string) []string {
	// Queue that store all the delayed project
	delayQueue := make([]string, 0)
	// List of project that is delayed
	res := make([]string, 0)
	// Hashmap that store the dependencies with key as the project and value as the dependant
	dependance := make(map[string][]string)
	// Cache checked project
	cache := make(map[string]bool)

	// Initialize the dependencies into the hashmap
	for _, each := range dependencies {
		// A,B == A --> B == A is depend on B
		if _, ok := dependance[each[1]]; !ok {
			dependance[each[1]] = make([]string, 0)
			dependance[each[1]] = append(dependance[each[1]], each[0])
		} else {
			dependance[each[1]] = append(dependance[each[1]], each[0])
		}
	}

	// Find the delay project's dependant
	for _, each := range delayProjects {
		delayed := dependance[each]
		cache[each] = true
		res = append(res, each)
		for _, depend := range delayed {
			if _, ok := cache[depend]; !ok {
				cache[depend] = true
				res = append(res, depend)
				Enqueue(&delayQueue, depend)
			}
		}

	}

	for len(delayQueue) > 0 {
		p := Dequeue(&delayQueue)
		if _, ok := cache[p]; !ok {
			delayed := dependance[p]
			cache[p] = true
			res = append(res, p)
			for _, each := range delayed {
				Enqueue(&delayQueue, each)
			}
		}

	}
	sort.Strings(res)
	return res
}

func BinaryTreeBFS(root *Node) [][]int {
	if root == nil {
		return nil
	}

	q := make([]*Node, 0)
	q = append(q, root)
	res := [][]int{}
	for len(q) > 0 {
		levelVal := make([]int, 0)
		size := len(q)
		for x := 0; x < size; x++ {
			curr := q[0]
			q = q[1:]

			levelVal = append(levelVal, curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		res = append(res, levelVal)
	}
	return res
}

// Enqueue: Add delay project into queue
func Enqueue(queue *[]string, input string) {
	(*queue) = append((*queue), input)
}

// Dequeue: Pop delay project from queue
func Dequeue(queue *[]string) string {
	out := (*queue)[0]
	(*queue)[0] = ""
	(*queue) = (*queue)[1:]
	return out
}
