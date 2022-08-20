package min_stops

// A Priority Queue implementation is copied from Go standart library documentation
// https://pkg.go.dev/container/heap#example-package-PriorityQueue

import (
	"container/heap"
)

type StationFuel = int

type PriorityQueue []*StationFuel

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return *pq[i] > *pq[j] }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*StationFuel)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func minRefuelStops(target int, fuel int, stations [][]int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	stationIdx := 0
	stops := 0
	dist := 0

	for {
		dist += fuel
		fuel = 0
		if dist >= target { // success!
			return stops
		}

		// check passed stations and add them to the pq
		for ; stationIdx < len(stations); stationIdx++ {
			st := stations[stationIdx]
			if st[0] > dist {
				break
			}
			heap.Push(&pq, &st[1])
		}

		if pq.Len() == 0 { // fail
			return -1
		}

		// if we have stations in the pq use the one with the most fuel and travel futher
		stFuel := heap.Pop(&pq).(*StationFuel)
		fuel += *stFuel
		stops += 1
	}
}
