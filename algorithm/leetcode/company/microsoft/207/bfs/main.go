package main

// bfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var degree = make([]int, numCourses)
	var edges = make([][]int, numCourses)
	var q []int
	for i := 0; i < len(prerequisites); i++ {
		node1 := prerequisites[i][1]
		node2 := prerequisites[i][0]
		degree[node2]++
		edges[node1] = append(edges[node1], node2)
	}
	for i := 0; i < len(degree); i++ {
		if degree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		numCourses--
		node := q[0]
		for i := 0; i < len(edges[node]); i++ {
			node2 := edges[node][i]
			degree[node2]--
			if degree[node2] == 0 {
				q = append(q, node2)
			}
		}
		q = q[1:]
	}
	if numCourses == 0 {
		return true
	} else {
		return false
	}

}
