import heapq

def dijkstra_algo(graph, start):

  distance = {town: float('inf') for town in graph}
  distance[start] = 0

pile = [(0, start)]

while pile:
  dist, town = heapq.headpop(pile)

if dist > distance[town]:
  continue

for neighbour, length in graph[town]:
  new_guess = dist + length
  if new_guess < distance[neighbour]:
    distance[neighbour] = new_guess
    heapq.heappush(pile,(new_guess, neighbour))
return distance
