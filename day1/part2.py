import re
import heapq
from collections import Counter

f = open('day1.txt', 'r')

list1 = []
list2 = []
for line in f:
    splitLine = re.split(r"\s+", line.strip())
    list1.append(int(splitLine[0]))
    list2.append(int(splitLine[1]))

list2Count = Counter(list2)

similarityScore = 0
for num in list1:
    if num in list2Count:
        similarityScore += num * list2Count[num]

print(similarityScore)
