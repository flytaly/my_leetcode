package courseschedule

type Status int

const (
	NotVisited Status = iota
	Visiting
	Visited
)

// Topological Sort
// https://www.youtube.com/watch?v=eL-KzMXSXXI

// returns false if graph is not DAG
func dfs(id int, statuses *map[int]Status, graph *map[int][]int) bool {
	if (*statuses)[id] == Visiting { // cycle
		return false
	}

	if (*statuses)[id] == Visited {
		return true
	}

	(*statuses)[id] = Visiting
	for _, next := range (*graph)[id] {
		if !dfs(next, statuses, graph) {
			return false
		}
	}
	(*statuses)[id] = Visited

	return true // the end => no more dependencies
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	statuses := map[int]Status{}
	for at := 0; at < numCourses; at++ {
		if statuses[at] == NotVisited {
			if !dfs(at, &statuses, &graph) {
				return false
			}
		}
	}

	return true
}
