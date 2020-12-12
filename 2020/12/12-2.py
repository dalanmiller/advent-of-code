from typing import List, Tuple
import re

test_case = """F10
N3
F7
R90
F11
L270
F10
R270"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 226

    rows = ["F10", "F10", "F10", "R180", "F10", "F10", "F10"]
    result = main(rows)
    print(result)
    assert result == 0

    rows = ["F10", "F10", "F10", "L180", "F10", "F10", "F10"]
    result = main(rows)
    print(result)
    assert result == 0

    rows = ["F10", "R90", "F10", "R90", "F10", "R90", "F10"]
    result = main(rows)
    print(result)
    assert result == 0

    rows = ["F10", "L90", "F10", "L90", "F10", "L90", "F10"]
    result = main(rows)
    print(result)
    assert result == 0

    rows = ["S1", "L90", "F10", "L90", "F100", "L180", "F200", "L180", "F100", "L90", "F10"]
    result = main(rows)
    print(result)
    assert result == 0

    print("Tests complete.")

line_pattern = re.compile(r"(\w)(\d*)")


def parse_line(row: str) -> Tuple[str, int]:
    action, value = re.match(line_pattern, row.strip()).groups()
    return action, int(value)


def test_turn():
    assert turn("L", 1, 1, 1) == (-1, 1)
    assert turn("L", -1, 1, 1) == (-1, -1)
    assert turn("L", -1, -1, 1) == (1, -1)
    assert turn("L", 1, -1, 1) == (1, 1)

    assert turn("R", 1, 1, 1) == (1, -1)
    assert turn("R", 1, -1, 1) == (-1, -1)
    assert turn("R", -1, -1, 1) == (-1, 1)
    assert turn("R", -1, 1, 1) == (1, 1)

    assert turn("L", -1, 1, 2) == (1, -1)
    assert turn("L", -1, 1, 3) == (1, 1)
    assert turn("L", -1, 1, 4) == (-1, 1)

    assert turn("R", 1, -1, 2) == (-1, 1)
    assert turn("R", 1, -1, 3) == (1, 1)
    assert turn("R", 1, -1, 4) == (1, -1)


def turn(dir: str, current_x, current_y, turns: int) -> Tuple[int, int]:
    x = current_x
    y = current_y

    for _ in range(turns):
        if dir == "L":
            x, y = -y, x
        elif dir == "R":
            x, y = y, -x

    return x, y


def main(rows: List[List[str]]) -> int:

    rows = [parse_line(x) for x in rows]

    waypoint_x = 10
    waypoint_y = 1

    current_x = 0
    current_y = 0

    for action, value in rows:
        if action == "N":
            waypoint_y += value
        elif action == "S":
            waypoint_y -= value
        elif action == "E":
            waypoint_x += value
        elif action == "W":
            waypoint_x -= value

        elif action in ("L", "R"):
            waypoint_x, waypoint_y = turn(action, waypoint_x, waypoint_y, value // 90)

        elif action == "F":
            current_x += waypoint_x * value
            current_y += waypoint_y * value

    return abs(current_x) + abs(current_y)


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    test_turn()
    print("Main result: ", main(rows))

