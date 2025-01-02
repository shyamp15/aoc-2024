import numpy as np
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
        self.prize = (10000000000000 + prize[0], 10000000000000 + prize[1])
        self.memo = {}

    def compute(self):
        det = self.a[0] * self.b[1] - self.a[1] * self.b[0]
        if det == 0:
            return 0
        pdet = self.b[1] * self.prize[0] - self.b[0] * self.prize[1]
        qdet = self.a[0] * self.prize[1] - self.a[1] * self.prize[0]
        if pdet % det != 0 or qdet % det != 0:
            return 0
        return 3 * pdet // det + qdet // det

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

min_tokens = 0
for game in game_list:
    min_tokens += game.compute()
print(min_tokens)
