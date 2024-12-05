def parseLine(line, increasing, first, tolerate):
    if len(line) < 2:
        return True

    prevNum = line[0]
    for num in line[1:]:
        diff = prevNum - num
        if not increasing:
            diff = -diff
        if diff < 1 or diff > 3:
            if tolerate > 0:
                tolerate -= 1
                continue
            if first:
                return parseLine(line[1:], increasing, False, tolerate)
            return False
        prevNum = num
    return True


f = open('day1.txt', 'r')

safeReports = 0
for line in f:
    splitLine = line.split(' ')
    intLine = [int(item) for item in splitLine]
    if parseLine(intLine, True, True, 1) or parseLine(intLine, False, True, 1):
        safeReports += 1
print(safeReports)
