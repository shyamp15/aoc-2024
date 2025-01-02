import re

file = open("day13.txt", "r")


def parseButton(line):
    x_pattern = r"X\+(\d+)"
    match = re.search(x_pattern, line)
    x = 0
    if match:
        x = int(match.group(1))

    y_pattern = r"Y\+(\d+)"
    match = re.search(y_pattern, line)
    y = 0
    if match:
        y = int(match.group(1))

    return (x, y)


def parsePrize(line):
    x_pattern = r"X\=(\d+)"
    match = re.search(x_pattern, line)
    x = 0
    if match:
        x = int(match.group(1))

    y_pattern = r"Y\=(\d+)"
    match = re.search(y_pattern, line)
    y = 0
    if match:
        y = int(match.group(1))

    return (x, y)


class Game:
    def __init__(self, a, b, prize):
        self.a = a
        self.b = b
        self.prize = prize
        self.memo = {}

    def compute(self):
        return self.recurse(0, 0, 0)

    def recurse(self, x, y, count):
        if count > 200 or x > self.prize[0] or y > self.prize[1]:
            return float('inf')
        if x == self.prize[0] and y == self.prize[1]:
            return 0
        if (x, y) in self.memo:
            return self.memo[(x, y)]

        # Choose A
        a = 3 + self.recurse(x + self.a[0], y + self.a[1], count + 1)
        # Choose B
        b = 1 + self.recurse(x + self.b[0], y + self.b[1], count + 1)

        self.memo[(x, y)] = min(a, b)
        return self.memo[(x, y)]

    def __repr__(self):
        return f"Game(a={self.a}, b={self.b}, prize={self.prize})"


lines = file.readlines()
i = 0
game_list = []
while i < len(lines):
    a = parseButton(lines[i].strip())
    b = parseButton(lines[i + 1].strip())
    prize = parsePrize(lines[i + 2].strip())
    game_list.append(Game(a, b, prize))
    i += 4
# print(game_list)

min_tokens = 0
for game in game_list:
    tokens = game.compute()
    if tokens != float('inf'):
        min_tokens += tokens
print(min_tokens)
