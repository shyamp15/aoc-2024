from collections import defaultdict

file = open("day8.txt", "r")

char_to_coord_map = defaultdict(list)

i = 0
for line in file:
    line = line.strip()
    j = 0
    for c in line:
        if c != '.' and not c.isspace():
            char_to_coord_map[c].append((i, j))
        j += 1
    i += 1

m = i
n = j


def check_bounds(coord, n, m):
    x, y = coord
    if x >= 0 and x < m and y >= 0 and y < n:
        return True
    return False


def find_antennas(top, bottom, antinodes):
    if top[0] > bottom[0] or (top[0] == bottom[0] and top[1] < bottom[1]):
        top, bottom = bottom, top

    x1, y1 = top
    x2, y2 = bottom
    x_diff = x2 - x1
    y_diff = y1 - y2

    slope = -1
    if not y_diff == 0:
        slope = x_diff / y_diff

    x_diff = abs(x_diff)
    y_diff = abs(y_diff)
    coord1 = None
    coord2 = None
    if slope >= 0:
        coord1 = (x1 - x_diff, y1 + y_diff)
        while check_bounds(coord1, m, n):
            antinodes.add(coord1)
            coord1 = (coord1[0] - x_diff, coord1[1] + y_diff)
        coord2 = (x2 + x_diff, y2 - y_diff)
        while check_bounds(coord2, m, n):
            antinodes.add(coord2)
            coord2 = (coord2[0] + x_diff, coord2[1] - y_diff)
    else:
        coord1 = (x1 - x_diff, y1 - y_diff)
        while check_bounds(coord1, m, n):
            antinodes.add(coord1)
            coord1 = (coord1[0] - x_diff, coord1[1] - y_diff)
        coord2 = (x2 + x_diff, y2 + y_diff)
        while check_bounds(coord2, m, n):
            antinodes.add(coord2)
            coord2 = (coord2[0] + x_diff, coord2[1] + y_diff)


antinodes = set()
for key in char_to_coord_map.keys():
    coord_list = char_to_coord_map[key]
    length = len(coord_list)
    if length > 1:
        antinodes.update(coord_list)
    for i in range(length):
        for j in range(i + 1, length):
            find_antennas(coord_list[i], coord_list[j], antinodes)


print(len(antinodes))
# print(antinodes)
