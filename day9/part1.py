import heapq

file = open("day9.txt", "r")

array = []
empty = []


for line in file:
    line = line.strip()
    index = 0
    id = 0
    for i, c in enumerate(line):
        num = int(c)
        if i % 2 == 0:
            for _ in range(num):
                array.append(id)
                index += 1
        else:
            for _ in range(num):
                array.append(-1)
                heapq.heappush(empty, index)
                index += 1
            id += 1

for i in range(len(array) - 1, -1, -1):
    if not empty:
        break
    if array[i] != -1:
        empty_idx = heapq.heappop(empty)
        if i <= empty_idx:
            break
        array[empty_idx] = array[i]
        array[i] = -1
        heapq.heappush(empty, i)

total = 0
for i in range(len(array)):
    if array[i] == -1:
        break

    total += array[i] * i

print(total)
