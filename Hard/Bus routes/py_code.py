from collections import deque, defaultdict

class Solution:
    def numBusesToDestination(self, routes: list[list[int]], source: int, target: int) -> int:
        if source == target:
            return 0
            
        stop_to_buses = defaultdict(list)
        for bus_id, route in enumerate(routes):
            for stop in route:
                stop_to_buses[stop].append(bus_id)
                
        queue = deque([source])
        visited_stops = {source}
        visited_buses = [False] * len(routes)
        buses_taken = 0
        
        while queue:
            buses_taken += 1
            for _ in range(len(queue)):
                stop = queue.popleft()
                
                for bus_id in stop_to_buses[stop]:
                    if visited_buses[bus_id]:
                        continue
                    visited_buses[bus_id] = True
                    
                    for next_stop in routes[bus_id]:
                        if next_stop == target:
                            return buses_taken
                        if next_stop not in visited_stops:
                            visited_stops.add(next_stop)
                            queue.append(next_stop)
        return -1
