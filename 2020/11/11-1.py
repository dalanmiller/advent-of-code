from collections import Counter
from typing import Tuple, List

test_case = """L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 37
    print("Tests complete.")

def test_num_adjacent_seats():
    rows = [["#", "#", "#"],["#", "L", "#"],["#", "#", "#"]]
    adjacent_taken_seats = num_adjacent_seats(rows, 1, 1)
    assert adjacent_taken_seats == 8

    rows = [["L", "#", "#"],["#", "#", "#"],["#", "#", "#"]]
    adjacent_taken_seats = num_adjacent_seats(rows, 0, 0)
    assert adjacent_taken_seats == 3

    rows = [["#", "#", "#"],["L", "#", "#"],["#", "#", "#"]]
    adjacent_taken_seats = num_adjacent_seats(rows, 1, 0)
    assert adjacent_taken_seats == 5

    rows = [["#", "#", "#"],["#", "#", "#"],["#", "#", "L"]]
    adjacent_taken_seats = num_adjacent_seats(rows, 2, 2)
    assert adjacent_taken_seats == 3

def test_seating_changed():
    orig = [[1],[2],[3]]
    new = [[1], [2], [3]]
    result = seating_changed(orig, new)
    assert result == False

    orig = [[1],[2],[3]]
    new = [[1], ["Y"], [3]]
    result = seating_changed(orig, new)
    assert result == True


def parse_line(row) -> str:
    return list(row.strip())
    
def num_adjacent_seats(rows: List[List[str]], x: int, y: int) -> int: 
    taken_seats = 0
    for x_pos in range(-1, 2):
        for y_pos in range(-1, 2):
            if x_pos == 0 and y_pos == 0:
                continue
            if all([y_pos + y != -1, y_pos + y < len(rows[x]), x_pos + x != -1, x_pos + x < len(rows)]):
                if rows[x_pos + x][y_pos + y] == "#":
                    taken_seats += 1
    
    return taken_seats
        
def seating_changed(orig: List[List[str]], new: List[List[str]]) -> bool:
    for x, y in zip(orig, new):
        if x != y:
            return True
    return False

def main(rows: List[List[str]]):
    # If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
    # If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
    # Otherwise, the seat's state does not change.

    prev_seating_chart = [parse_line(x) for x in rows]
    new_seating_chart = [parse_line(x) for x in rows]
    changed = True

    while changed:
        for i, row in enumerate(prev_seating_chart):
            for j, location in enumerate(row):
                if location == "L": 
                    adjacent_seats = num_adjacent_seats(prev_seating_chart, i, j)
                    if adjacent_seats == 0:
                        new_seating_chart[i][j] = '#'
                
                elif location == "#":
                    adjacent_seats = num_adjacent_seats(prev_seating_chart, i, j)
                    if adjacent_seats >= 4:
                        new_seating_chart[i][j] = 'L'

        changed = seating_changed(prev_seating_chart, new_seating_chart)
        prev_seating_chart = [[y for y in x] for x in new_seating_chart]

    count = Counter() 
    for row in new_seating_chart:
        count.update(row)

    return count["#"]

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]
    
    test_num_adjacent_seats()
    test_seating_changed()
    test()
    print("Main result: ", main(rows))
