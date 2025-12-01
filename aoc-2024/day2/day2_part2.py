
safe_count = 0

file = open("input.txt", "r")
puzzle_input = file.readlines()

for report in puzzle_input:
    levels = list(map(int, report.split()))
    
    is_increasing = all(1 <= b - a <= 3 for a, b in zip(levels, levels[1:]))
    is_decreasing = all(1 <= a - b <= 3 for a, b in zip(levels, levels[1:]))
    
    if is_increasing or is_decreasing:
        safe_count += 1 
    else:
        for n in range(len(levels)):
            new_levels = levels[:n] + levels[n+1:]
            is_increasing = all(1 <= b - a <= 3 for a, b in zip(new_levels, new_levels[1:]))
            is_decreasing = all(1 <= a - b <= 3 for a, b in zip(new_levels, new_levels[1:]))
            if is_increasing or is_decreasing:
                safe_count += 1 
                break

print(safe_count)