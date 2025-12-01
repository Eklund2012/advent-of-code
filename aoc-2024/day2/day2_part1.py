
safe_count = 0

with open("input.txt", "r") as f:
    puzzle_input = f.readlines()

for report in puzzle_input:
    levels = list(map(int, report.split()))
    
    is_increasing = all(1 <= b - a <= 3 for a, b in zip(levels, levels[1:]))
    is_decreasing = all(1 <= a - b <= 3 for a, b in zip(levels, levels[1:]))
    
    if is_increasing or is_decreasing:
        safe_count += 1 

print(safe_count)