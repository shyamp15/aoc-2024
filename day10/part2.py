file = open("day10.txt", "r")


starts = []
matrix = []
i = 0
for line in file:
    curr = []
    matrix.append(curr)
    line = line.strip()
    j = 0
    for c in line:
        num = -1
        if c != '.':
            num = int(c)
        curr.append(num)
        if num == 0:
            starts.append((i, j))

        j += 1
    i += 1

m = len(matrix)
n = len(matrix[0])

directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]


def dfs(start, num):
    if num == 9:
        return 1
    i, j = start
    total = 0
    for x, y in directions:
        newX, newY = i + x, j + y
        if newX >= 0 and newX < m and newY >= 0 and newY < n and matrix[newX][newY] == num + 1:
            total += dfs((newX, newY), num + 1)
    return total


trailheads = 0
for start in starts:
    trailheads += dfs(start, 0)

print(trailheads)
