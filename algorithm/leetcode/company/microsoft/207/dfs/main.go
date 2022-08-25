package dfs

var visit []int
var edges [][]int
var res []int
var valid bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges = make([][]int, numCourses)
	visit = make([]int, numCourses)
	res = nil
	valid = true
	for i := 0; i < len(prerequisites); i++ {
		first := prerequisites[i][1] //必须条件
		second := prerequisites[i][0]
		edges[first] = append(edges[first], second)
	}

	for i := 0; i < numCourses; i++ {
		if !valid {
			break
		} else {
			if visit[i] == 0 {
				dfs(i)
			}
		}
	}
	return valid
}

func dfs(node int) {
	visit[node] = 1
	for i := 0; i < len(edges[node]); i++ {
		node1 := edges[node][i]
		if visit[node1] == 0 {
			dfs(node1)
			if !valid {
				return
			}
		} else if visit[node1] == 1 {
			valid = false
			return
		}
	}
	visit[node] = 2
}
