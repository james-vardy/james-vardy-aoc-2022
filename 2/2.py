f = open('input.txt', 'r')
inputData = f.readlines()

# split into dict with key [round], linking to a list of your opponents play, and yours
round = 0
strategyDict = {} # e.g. {0: ['A', 'Y'], 1: ['B', 'Z']}

for element in inputData:
    choiceList = []
    choiceList.append(element[0])
    choiceList.append(element[2])
    strategyDict[round] = choiceList
    round += 1

'''
The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors.
The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors.

Your total score is the sum of your scores for each round.
The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

'''

def decideWin(opp, you): # returns 0 if loss, 1 if win, 2 if draw
    if opp == 'A': # if opp rock
        if you == 'X':
            return 2
        if you == 'Y':
            return 1
        if you == 'Z':
            return 0
        
    if opp == 'B': # if opp paper
        if you == 'X':
            return 0
        if you == 'Y':
            return 2
        if you == 'Z':
            return 1
        
    if opp == 'C': # if opp scissors
        if you == 'X':
            return 1
        if you == 'Y':
            return 0
        if you == 'Z':
            return 2

def outcomePoints(result): # returns points based on outcome
    if result == 0:
        return 0
    if result == 1:
        return 6
    if result == 2:
        return 3

def choicePoints(choice): # returns points based on player choice
    if choice == 'X':
        return 1
    if choice == 'Y':
        return 2
    if choice == 'Z':
        return 3

# calculate total points
totalPoints = 0
for i in range(round):
    opp = (strategyDict[i])[0]
    you = (strategyDict[i])[1]
    totalPoints += outcomePoints(decideWin(opp, you)) + choicePoints(you)

print(totalPoints) # part 1 answer

#part 2

'''
"Anyway, the second column says how the round needs to end:
X means you need to lose,
Y means you need to end the round in a draw,
and Z means you need to win. Good luck!"

'''

def calcChoice(opp, res): #add in a function that calculates what choice you pick based on col 2
    if res == 'X': #loss
        if opp == 'A':
            return 'Z'
        if opp == 'B':
            return 'X'
        if opp == 'C':
            return 'Y'

    if res == 'Y': #draw
        if opp == 'A':
            return 'X'
        if opp == 'B':
            return 'Y'
        if opp == 'C':
            return 'Z'

    if res == 'Z': #win
        if opp == 'A':
            return 'Y'
        if opp == 'B':
            return 'Z'
        if opp == 'C':
            return 'X'

# calculate total points same as before but calc your choice from res first
totalPoints = 0
for i in range(round):
    opp = (strategyDict[i])[0]
    res = (strategyDict[i])[1]
    you = calcChoice(opp, res)
    totalPoints += (outcomePoints(decideWin(opp, you)) + choicePoints(you))

print(totalPoints) # part 2 answer

