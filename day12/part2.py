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
directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]


class Plot:
    def __init__(self, c, i, j):
        self.char = c
        self.nodes = set()
        self.nodes.add((i, j))
        self.sides = 4

    def add_node(self, node):
        self.nodes.add(node)
        for x, y in directions:
            r, c = node[0] + x, node[1] + y
            if r < 0 or c < 0 or r >= m or c >= n:
                self.sides += 1
            else:
                if (r, c) in self.nodes:
                    self.sides -= 1
                else:
                    self.sides += 1

    def __repr__(self):
        return f"Plot(char='{self.char}', perimeter={self.perimeter}, nodes={self.nodes})"


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
            plot_list.append(plot)

price = 0
for plot in plot_list:
    price += plot.sides * len(plot.nodes)
# print(plot_list)
print(price)
