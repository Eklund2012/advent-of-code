
from collections import Counter

left_list = []
right_list = []
total_sum = 0

with open("input.txt", "r") as f:
    inp = f.readlines()

for pair in inp:
    a, b = map(int, pair.split())
    left_list.append(a)
    right_list.append(b)

left_list.sort()
right_list.sort()


right_count = Counter(right_list)

for n in left_list:
    if n in right_count: 
        total_sum += n * right_count[n]

print(total_sum)
