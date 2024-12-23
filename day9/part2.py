from collections import defaultdict, deque
import heapq


file = open("day9.txt", "r")

file_list = []
empty_map = {}


class File:
    def __init__(self, start, length, id):
        self.start = start
        self.length = length
        self.id = id

    def __lt__(self, other):
        return self.start < other.start

    def __str__(self):
        return f"File(start={self.start}, length={self.length}, id={self.id})"


for line in file:
    line = line.strip()
    index = 0
    id = 0
    for i, c in enumerate(line):
        num = int(c)
        if i % 2 == 0:
            file_list.append(File(index, num, id))
            index += num
        else:
            if num != 0:
                if num not in empty_map:
                    empty_map[num] = []
                heapq.heappush(empty_map[num], index)

                index += num
            id += 1


def search(length):
    min = float('inf')
    ret = -1
    for key in empty_map.keys():
        curr = empty_map[key]
        if key >= length and curr[0] < min:
            min = curr[0]
            ret = key
    return ret


new_file_list = []
for i in range(len(file_list) - 1, -1, -1):
    file = file_list[i]
    key = search(file.length)
    if key != -1 and file.start > empty_map[key][0]:
        new_start = heapq.heappop(empty_map[key])
        if len(empty_map[key]) == 0:
            del empty_map[key]
        heapq.heappush(new_file_list, File(new_start, file.length, file.id))
        new_length = key - file.length
        if new_length > 0:
            if new_length not in empty_map:
                empty_map[new_length] = []
            heapq.heappush(empty_map[new_length], new_start + file.length)

    else:
        heapq.heappush(new_file_list, file)

# arr = [-1] * 52000
total = 0
# new_file = open("wrong.txt", "w")
while new_file_list:
    file = heapq.heappop(new_file_list)
    for i in range(file.length):
        total += (file.start + i) * file.id
        # new_file.write(str(file.start + i) + "\t" + str(file.id) + "\n")
        # arr[file.start + i] = file.id

# print(arr)
# print(empty_map)

print(total)
