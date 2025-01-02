file = open("day12.txt", "r")

grid = []

for line in file:
    line = line.strip()
    row = []
    for c in line:
        row.append(c)
    grid.append(row)

m = len(grid)
n = len(grid[0])

visited = set()
directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]


class Plot:
    def __init__(self, c, i, j):
        self.char = c
        self.nodes = set()
        self.nodes.add((i, j))
        self.sides = 0

    def add_node(self, node):
        self.nodes.add(node)

    def calc_sides(self):
        for node in self.nodes:
            for i in range(4):
                dir1 = directions[i]
                dir2 = directions[(i + 1) % 4]

                if (((node[0] + dir1[0], node[1] + dir1[1]) not in self.nodes) and
                        ((node[0] + dir2[0], node[1] + dir2[1]) not in self.nodes)):
                    self.sides += 1
                if (((node[0] + dir1[0], node[1] + dir1[1]) in self.nodes) and
                        ((node[0] + dir2[0], node[1] + dir2[1]) in self.nodes) and
                        ((node[0] + dir1[0] + dir2[0], node[1] + dir1[1] + dir2[1]) not in self.nodes)):
                    self.sides += 1

    def __repr__(self):
        return f"Plot(char='{self.char}', sides={self.sides}, nodes={self.nodes})"


def dfs(i, j, plot):
    plot.add_node((i, j))
    visited.add((i, j))

    for x, y in directions:
        r, c = i + x, y + j
        if r >= 0 and c >= 0 and r < m and c < n and (r, c) not in visited and grid[r][c] == grid[i][j]:
            dfs(r, c, plot)


plot_list = []
for i in range(m):
    for j in range(n):
        if (i, j) not in visited:
            plot = Plot(grid[i][j], i, j)
            dfs(i, j, plot)
            plot.calc_sides()
            plot_list.append(plot)

price = 0
for plot in plot_list:
    price += plot.sides * len(plot.nodes)
print(price)
