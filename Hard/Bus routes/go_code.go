func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stopToBuses := make(map[int][]int)
	for busId, route := range routes {
		for _, stop := range route {
			stopToBuses[stop] = append(stopToBuses[stop], busId)
		}
	}

	queue := []int{source}
	visitedStops := map[int]bool{source: true}
	visitedBuses := make([]bool, len(routes))
	busesTaken := 0

	for len(queue) > 0 {
		busesTaken++
		size := len(queue)

		for i := 0; i < size; i++ {
			stop := queue[0]
			queue = queue[1:]

			for _, busId := range stopToBuses[stop] {
				if visitedBuses[busId] {
					continue
				}
				visitedBuses[busId] = true

				for _, nextStop := range routes[busId] {
					if nextStop == target {
						return busesTaken
					}
					if !visitedStops[nextStop] {
						visitedStops[nextStop] = true
						queue = append(queue, nextStop)
					}
				}
			}
		}
	}
	return -1
}
