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
    return [int(x) for x in row.strip().split(",") if x != "x"]

import math
def lcm(a, b):
    return abs(a*b) // math.gcd(a, b)

def main(rows: List[List[str]]) -> int:

    start_minute = parse_first_line(rows[0])
    buses = parse_second_line(rows[1])

    min_bus = 0
    min_delta = sys.maxsize
    buses.sort()
    buses.reverse()
    print(buses)

    for i in range(start_minute, sys.maxsize):        
        for bus in buses:
            if lcm(bus, i) == i:
                print(i, bus)
                min_bus = bus
                min_delta = i - start_minute
                break
        
        if min_bus:
            break

    return min_bus * min_delta


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
    # 1939887

