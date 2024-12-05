def parseLine(line):
    if len(line) < 2:
        return True

    increasing = True
    if line[0] - line[1] < 0:
        increasing = False

    prevNum = line[0]
    for num in line[1:]:
        diff = prevNum - num
        if not increasing:
            diff = -diff
        if diff < 1 or diff > 3:
            return False
        prevNum = num
    return True


f = open('day2.txt', 'r')

safeReports = 0
for line in f:
    splitLine = line.split(' ')
    intLine = [int(item) for item in splitLine]
    if parseLine(intLine):
        safeReports += 1
print(safeReports)
