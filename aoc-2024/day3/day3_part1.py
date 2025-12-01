import re

with open("input.txt", "r") as f:
    inp = str(f.readlines())

result = str(re.findall('mul\(\d+,\s*\d+\)', inp))

numbers = list(map(int, re.findall('\d+', result)))

answer = 0

for a, b in zip(numbers[::2], numbers[1::2]):
    answer += a * b

print(answer)