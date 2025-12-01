import re
from itertools import product

with open('input.txt', 'r') as f:
    inp = [line.strip() for line in f.readlines()]

extracted_numbers = [re.findall(r'\d+', s) for s in inp]

total = 0
result = 0

for lines in extracted_numbers:
    target = int(lines[0])
    operands = list(map(int, lines[1:]))
    combinations = list(product(["+", "*"], repeat=len(operands)-1))
    for operator in combinations:
        result = operands[0] 
        for i in range(len(operator)):  #
            if operator[i] == "+":
                result += operands[i + 1]
            elif operator[i] == "*":
                result *= operands[i + 1]
        if result == target: 
            total += target
            break

print(total)







