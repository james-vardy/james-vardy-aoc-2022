f = open('input.txt', 'r')
inputData = f.readlines()

elfDict = {}
elfNum = 0
elfDict[elfNum] = []

# creates dict of elves, key of elf number, relates to an array of their snacks
for element in inputData:
    if element == "\n":
        elfNum += 1
        elfDict[elfNum] = []
    else:
        elfDict[elfNum].append(int(element))


#order elf Id's by their total calories
elfMaxCalories = 0
maxElf = 0
elfOrderList = []
tempElfDict = elfDict.copy()

# orders elf by calories
for orderElf in range(elfNum):
    elfMaxCalories = 0
    for elf in range(elfNum):
        elftotalCalories = 0
        for foodItem in tempElfDict[elf]:
            elftotalCalories += foodItem
        if elftotalCalories > elfMaxCalories:
            elfMaxCalories = elftotalCalories
            maxElf = elf
    tempElfDict[maxElf] = [0]
    elfOrderList.append(maxElf)

# return total calories of top x elves
x = 3 # answer for part 1: x=1 / answer for part 2: x=3
totalCals = 0

for i in range(x):
    elf = elfOrderList[i]
    for foodItem in elfDict[elf]:
            totalCals += foodItem

print(totalCals)