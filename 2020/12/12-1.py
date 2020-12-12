from typing import List, Tuple
from itertools import cycle
import re

test_case = """F10
N3
F7
R90
F11"""

test_case_2 = """F12
W1
N3
E3
W3
F93
N2
R90
N4
L180
F13
E2
R270"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 25
    print("Tests complete.")

def test2():
    rows = [x for x in test_case_2.split("\n")]
    result = main(rows)
    print(result)
    assert result == 128
    print("Tests complete.")

line_pattern = re.compile(r"(\w)(\d*)")
def parse_line(row: str) -> Tuple[str, int]:
    action, value = re.match(line_pattern, row.strip()).groups()
    return action, int(value)

def main(rows: List[List[str]]) -> int:

    # Action N means to move north by the given value.
    # Action S means to move south by the given value.
    # Action E means to move east by the given value.
    # Action W means to move west by the given value.
    # Action L means to turn left the given number of degrees.
    # Action R means to turn right the given number of degrees.
    # Action F means to move forward by the given value in the direction the ship is currently facing.

    rows = [parse_line(x) for x in rows]

    direction = "E"

    direction_ops = {
        "N": lambda x, y, d: (x, y + d),
        "S": lambda x, y, d: (x, y - d),
        "E": lambda x, y, d: (x + d, y),
        "W": lambda x, y, d: (x - d, y),
    }

    dirs = ["N", "E", "S", "W"]

    current_x = 0
    current_y = 0

    for action, value in rows:
        if action == "N":
            current_y += value
        elif action == "S":
            current_y -= value
        elif action == "E":
            current_x += value
        elif action == "W":
            current_x -= value

        elif action == "L":
            turns = value // 90
            i = dirs.index(direction)
            direction = dirs[(i - turns) % 4]
            
        elif action == "R":
            turns = value // 90
            i = dirs.index(direction)
            direction = dirs[(i + turns) % 4]

        elif action == "F":
            current_x, current_y = direction_ops[direction](current_x, current_y, value)

    return abs(current_x) + abs(current_y)


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    test2()
    print("Main result: ", main(rows))
