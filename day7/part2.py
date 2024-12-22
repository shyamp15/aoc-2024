f = open("day7.txt")


def findEquation(numbers, curr_value, curr_index, test_value):
    if curr_value == test_value and curr_index == len(numbers):
        return True
    if curr_index >= len(numbers):
        return False

    first = findEquation(numbers, curr_value +
                         numbers[curr_index], curr_index + 1, test_value)
    second = False
    third = False
    if curr_index != 0:
        second = findEquation(numbers, curr_value *
                              numbers[curr_index], curr_index + 1, test_value)
        concat = int(str(curr_value) + str(numbers[curr_index]))
        third = findEquation(numbers, concat, curr_index + 1, test_value)

    return first or second or third


total = 0
for line in f:
    split_line = line.split(":")
    test_value = int(split_line[0])

    numbers = [int(item) for item in split_line[1].strip().split(" ")]
    if findEquation(numbers, 0, 0, test_value):
        total += test_value


print(total)
