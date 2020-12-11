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

test_case_two = """LLLLL.LLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLLL.LLLL
LLLLLLLLLL.LLL.LLLLL.LLLLLL.LLLLLLL.LLLLL.LLLLL
LL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLL.LLLLLLL.LLL
LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLL.LLLLLLLLLLL.LLL
LLLLL.LLLLLLLLLLLLLL.LLL.LL.LLLL.LLL.LLLLLL.LLL
LLLLL.LLLL..LLLLLLLL.LLLLLLLLLLL.LLLLLLL.LLLLLL
.......L...L.L...L......L..LLL..L.............L
LLLLL..LLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLL..LL
LLLLLLLLLL.LLLLLLLLL.LLLLLL..LLLLLL.LLLLLLL.LLL
L.LLL.LLLLLLLLLLLLLLLLLLLLL.LLLLLL..LLLLLLL.LLL
LLLLL.LLLL.LLLLLLLLL.LLLLLLLLLLLLLL.LLL.LLLLLLL
LLLLLLLLLL.LLLLLLLLLLLLLLLL.LLL.LLLLLLLLLLL.LLL
LLLLL.LLLL.LLLLLLLLL.LLLLLL.LLLLLLL.LLLLL.L.LLL
.....LL..L....LL.L....L....L.....L.L...L......L"""


def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 26
    print("Tests complete.")


def test2():
    rows = [x for x in test_case_two.split("\n")]
    result = main(rows)
    print(result)
    assert result == 160
    print("Tests complete.")


def test_num_adjacent_seats():
    rows = [["#", "#", "#"], ["#", "L", "#"], ["#", "#", "#"]]
    adjacent_taken_seats = num_visible_seats(rows, 1, 1)
    assert adjacent_taken_seats == 8

    rows = [["L", "#", "#"], ["#", "#", "#"], ["#", "#", "#"]]
    adjacent_taken_seats = num_visible_seats(rows, 0, 0)
    assert adjacent_taken_seats == 3

    rows = [["#", "#", "#"], ["L", "#", "#"], ["#", "#", "#"]]
    adjacent_taken_seats = num_visible_seats(rows, 1, 0)
    assert adjacent_taken_seats == 5

    rows = [["#", "#", "#"], ["#", "#", "#"], ["#", "#", "L"]]
    adjacent_taken_seats = num_visible_seats(rows, 2, 2)
    assert adjacent_taken_seats == 3

    rows = [
        ["#", "#", "#", "#", "#"],
        ["#", ".", ".", ".", "#"],
        ["#", "#", "L", "#", "#"],
        ["#", ".", ".", ".", "#"],
        ["#", ".", ".", ".", "#"],
    ]
    adjacent_taken_seats = num_visible_seats(rows, 2, 2)
    assert adjacent_taken_seats == 7

    rows = [
        ["#", "#", "#", "#", "#"],
        ["#", ".", ".", ".", "#"],
        ["#", "#", "#", "#", "#"],
        ["#", ".", ".", ".", "#"],
        ["#", ".", ".", ".", "L"],
    ]
    adjacent_taken_seats = num_visible_seats(rows, 4, 4)
    assert adjacent_taken_seats == 3

    rows = [
        [".", "#", ".", "#", "."],
        ["#", ".", ".", ".", "#"],
        [".", ".", "L", ".", "."],
        ["#", ".", ".", ".", "#"],
        [".", "#", ".", "#", "."],
    ]
    adjacent_taken_seats = num_visible_seats(rows, 2, 2)
    assert adjacent_taken_seats == 0

    example_case = """.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#....."""
    rows = [list(x) for x in example_case.split("\n")]
    adjacent_taken_seats = num_visible_seats(rows, 4, 3)
    assert adjacent_taken_seats == 8

    example_case_2 = """.............
.L.L.#.#.#.#.
............."""

    rows = [list(x) for x in example_case_2.split("\n")]
    adjacent_taken_seats = num_visible_seats(rows, 1, 1)
    assert adjacent_taken_seats == 0


def test_seating_changed():
    orig = [[1], [2], [3]]
    new = [[1], [2], [3]]
    result = seating_changed(orig, new)
    assert result == False

    orig = [[1], [2], [3]]
    new = [[1], ["Y"], [3]]
    result = seating_changed(orig, new)
    assert result == True


def test_out_of_bounds():
    rows = [
        ["#", "#", "#", "#", "#"],
        ["#", ".", ".", ".", "#"],
        ["#", "#", "L", "#", "#"],
        ["#", ".", ".", ".", "#"],
        ["#", ".", ".", ".", "#"],
    ]
    assert out_of_bounds(rows, -1, -1) == True
    assert out_of_bounds(rows, 5, 0) == True
    assert out_of_bounds(rows, 0, 0) == False
    assert out_of_bounds(rows, 0, 5) == True


def test_iterate_seating_chart():
    frame1 = """L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL"""

    frame2 = """#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##"""

    frame3 = """#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#"""

    frame4 = """#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#"""

    expected_rows_2 = [list(x) for x in frame2.split("\n")]
    expected_rows_3 = [list(x) for x in frame3.split("\n")]
    expected_rows_4 = [list(x) for x in frame4.split("\n")]

    old_rows = [list(x) for x in frame1.split("\n")]
    new_rows = [list(x) for x in frame1.split("\n")]

    iterate_seating_chart(old_rows, new_rows)
    assert new_rows == expected_rows_2
    old_rows = [[y for y in x] for x in new_rows]
    iterate_seating_chart(old_rows, new_rows)
    assert new_rows == expected_rows_3
    old_rows = [[y for y in x] for x in new_rows]
    iterate_seating_chart(old_rows, new_rows)
    assert new_rows == expected_rows_4


def parse_line(row) -> str:
    return list(row.strip())


def out_of_bounds(rows: List[List[str]], x: int, y: int) -> bool:
    return x < 0 or y < 0 or x >= len(rows) or y >= len(rows[0])


def num_visible_seats(rows: List[List[str]], x: int, y: int) -> int:
    direction_ops = {
        "n": lambda x, y: (x, y - 1),
        "ne": lambda x, y: (x + 1, y - 1),
        "e": lambda x, y: (x + 1, y),
        "se": lambda x, y: (x + 1, y + 1),
        "s": lambda x, y: (x, y + 1),
        "sw": lambda x, y: (x - 1, y + 1),
        "w": lambda x, y: (x - 1, y),
        "nw": lambda x, y: (x - 1, y - 1),
    }

    taken_seats = 0
    for direction in direction_ops.keys():
        current_x, current_y = direction_ops[direction](x, y)
        found_empty = False
        while not out_of_bounds(rows, current_x, current_y) and not found_empty:
            target_location = rows[current_x][current_y]
            if target_location == "#":
                taken_seats += 1
                break
            elif target_location == "L":
                found_empty = True
                break

            current_x, current_y = direction_ops[direction](current_x, current_y)

    return taken_seats


def seating_changed(orig: List[List[str]], new: List[List[str]]) -> bool:
    for x, y in zip(orig, new):
        if x != y:
            return True
    return False


def iterate_seating_chart(prev_seating_chart, new_seating_chart):
    for i, row in enumerate(prev_seating_chart):
        for j, location in enumerate(row):
            if location == "L":
                adjacent_seats = num_visible_seats(prev_seating_chart, i, j)
                if adjacent_seats == 0:
                    new_seating_chart[i][j] = "#"

            elif location == "#":
                adjacent_seats = num_visible_seats(prev_seating_chart, i, j)
                if adjacent_seats >= 5:
                    new_seating_chart[i][j] = "L"


def main(rows: List[List[str]]) -> int:
    prev_seating_chart = [parse_line(x) for x in rows]
    new_seating_chart = [parse_line(x) for x in rows]
    changed = True
    while changed:
        iterate_seating_chart(prev_seating_chart, new_seating_chart)

        changed = seating_changed(prev_seating_chart, new_seating_chart)
        prev_seating_chart = [[y for y in x] for x in new_seating_chart]

        with open("output", "a") as o:
            for row in prev_seating_chart:
                o.write("".join(row) + "\n")
            o.write("\n")

    count = Counter()
    for row in new_seating_chart:
        count.update(row)

    return count["#"]


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test_out_of_bounds()
    test_num_adjacent_seats()
    test_seating_changed()
    test_iterate_seating_chart()
    test2()
    test()
    print("Main result: ", main(rows))