from typing import List, Tuple
import sys

test_case = """939
7,13,x,x,59,x,31,19"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 295
    print("Tests complete.")


def parse_first_line(row: str) -> int:
    start_minute = int(row.strip())
    return start_minute

def parse_second_line(row: str) -> Tuple[int]:
    return map(int, filter(lambda x: x != "x", row.split(",")))

def main(rows: List[List[str]]) -> int:

    start_minute = rows[0]
    buses = rows[1]

    min_delta = sys.maxsize()
    buses.sort()

    for bus in buses: 
        temp_min = bus
        while temp_min < start_minute: 
            temp_min * 2

        if temp_min < min_delta:
            min_delta = temp_min

    return min_delta * (min_delta - start_minute)


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
