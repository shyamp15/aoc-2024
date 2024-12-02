import re
import heapq

f = open('day1.txt', 'r')

heap1 = []
heap2 = []
for line in f:
    splitLine = re.split(r"\s+", line.strip())
    heap1.append(int(splitLine[0]))
    heap2.append(int(splitLine[1]))

heapq.heapify(heap1)
heapq.heapify(heap2)

total = 0
for _ in range(len(heap1)):
    val1 = heapq.heappop(heap1)
    val2 = heapq.heappop(heap2)
    total += abs(val1 - val2)

print(total)
