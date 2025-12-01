import re

with open("test.txt", "r") as f:
    inp = str(f.read())

pattern = r"do\(\)|don't\(\)|mul\(\d+,\d+\)"
instructions = re.findall(pattern, inp)

print(instructions)

total = 0
enabled = True

for instr in instructions:
    if instr == "do()":
        enabled = True
    elif instr == "don't()":
        enabled = False
    elif instr.startswith("mul"):
        if enabled:
            x, y = map(int, re.findall(r"\d+", instr))
            total += x * y

print(total)