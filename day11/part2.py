import sys
import functools

if len(sys.argv) < 2:
    exit()

blink = int(sys.argv[1])

file = open("day11.txt", "r")

for line in file:
    split_line = line.strip().split(" ")
    nums = [int(x) for x in split_line]

memo = {}
memo2 = {}


def count_single_blink(num):
    if num in memo2:
        return memo2[num]

    str_num = str(num)
    res = ()
    if num == 0:
        res = (1, None)
    elif len(str_num) % 2 == 0:
        mid = len(str_num) // 2
        first = int(str_num[:mid])
        second = int(str_num[mid:])
        res = (first, second)
    else:
        res = (num * 2024, None)
    memo2[num] = res
    return res


def count_stone_blinks(num, depth):
    left_stone, right_stone = count_single_blink(num)

    if (num, depth) in memo:
        return memo[(num, depth)]

    if depth == 1:
        if right_stone is None:
            return 1
        else:
            return 2

    total = count_stone_blinks(left_stone, depth - 1)
    if right_stone is not None:
        total += count_stone_blinks(right_stone, depth - 1)
    memo[(num, depth)] = total
    return total


count = 0
for num in nums:
    count += count_stone_blinks(num, blink)

print(count)
