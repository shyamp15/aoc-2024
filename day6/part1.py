def getDirection(direction):
    if direction == (-1, 0):
        return (0, 1)
    elif direction == (0, 1):
        return (1, 0)
    elif direction == (1, 0):
        return (0, -1)
    elif direction == (0, -1):
        return (-1, 0)


f = open('day6.txt', 'r')

matrix = []
start = None
i = 0
for line in f:
    matrix.append([])
    j = 0
    for c in line:
        if c == '.':
            matrix[i].append(True)
        elif c == '#':
            matrix[i].append(False)
        elif c == '^':
            matrix[i].append(True)
            start = (i, j)
        j += 1
    i += 1

x, y = start
direction = (-1, 0)
visited = set()
visited.add((x, y))
while True:
    newX, newY = x + direction[0], y + direction[1]
    if not (newX >= 0 and newX < len(matrix) and newY >= 0 and newY < len(matrix[0])):
        break
    if matrix[newX][newY]:
        x = newX
        y = newY
        visited.add((x, y))
    else:
        direction = getDirection(direction)

print(len(visited))
