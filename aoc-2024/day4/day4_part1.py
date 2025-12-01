import re

with open('test.txt', 'r') as f:
    inp = [line.strip() for line in f.readlines()]

pattern = "XMAS"
pattern_rev = "SAMX"

def horizontal(inp):
    hori_xmas_count = 0
    for line in inp:
        hori_xmas_count += len(re.findall(pattern, line))
        hori_xmas_count += len(re.findall(pattern_rev, line))
    return hori_xmas_count

def vertical(inp):
    vert_xmas_count = 0
    combined_columns = zip(*inp)
    new_list = [''.join(column) for column in combined_columns]
    for col in new_list:
        vert_xmas_count += len(re.findall(pattern, col))
        vert_xmas_count += len(re.findall(pattern_rev, col))
    return vert_xmas_count

def diagonal(inp):
    diag_xmas_count = 0
    diagonal_list = extract_diagonals(inp)
    for set in diagonal_list:
        diag_xmas_count += len(re.findall(pattern, set))
        diag_xmas_count += len(re.findall(pattern_rev, set))
    return diag_xmas_count

def extract_diagonals(inp):
    rows, cols = len(inp), len(inp[0])
    diagonals, locations = [], []

    for d in range(rows + cols - 1):
        diagonal, locs = [], []
        for row in range(rows):
            col = d - row
            if 0 <= col < cols:
                diagonal.append(inp[row][col])
                locs.append((row, col))
        diagonals.append(''.join(diagonal))
        locations.append(locs)

    for d in range(rows + cols - 1):
        diagonal, locs = [], []
        for row in range(rows):
            col = d - (rows - 1 - row)
            if 0 <= col < cols:
                diagonal.append(inp[row][col])
                locs.append((row, col))
        diagonals.append(''.join(diagonal))
        locations.append(locs)

    return diagonals

xmas_counter = horizontal(inp) + vertical(inp) + diagonal(inp)

print(xmas_counter)