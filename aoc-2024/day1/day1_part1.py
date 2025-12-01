
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

for a, b in zip(left_list, right_list):
    total_sum += abs(a-b)

print(total_sum)


