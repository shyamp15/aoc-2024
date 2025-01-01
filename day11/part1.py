import sys

if len(sys.argv) < 2:
    exit()

blink = int(sys.argv[1])

file = open("day11.txt", "r")

for line in file:
    split_line = line.strip().split(" ")
    nums = [int(x) for x in split_line]

for _ in range(blink):
    temp = []
    for num in nums:
        str_num = str(num)
        if num == 0:
            temp.append(1)
        elif len(str_num) % 2 == 0:
            length = len(str_num) // 2
            first = int(str_num[:length])
            second = int(str_num[length:])
            temp.append(first)
            temp.append(second)
        else:
            temp.append(num * 2024)
    nums = temp


print(len(nums))
