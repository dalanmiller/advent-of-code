from typing import List, Tuple, Any
import itertools
import sys
import math

def test():
    rows = [str(x) for x in [17, "x", 13, 19]]
    result = main(rows)
    print(result)
    assert result == 3417

    rows = [str(x) for x in [67, 7, 59, 61]]
    result = main(rows)
    print(result)
    assert result == 754018

    rows = [str(x) for x in [67, "x", 7, 59, 61]]
    result = main(rows)
    print(result)
    assert result == 779210

    rows = [str(x) for x in [67, 7, "x", 59, 61]]
    result = main(rows)
    print(result)
    assert result == 1261476

    print("Tests complete.")


def parse_second_line(row: str) -> List[Any]:
    return [int(x) if x != "x" else x for x in row]

def main(rows: List[List[str]], start_timestamp:int = 0) -> int:
    buses = parse_second_line(rows)

    bus_ids = len([x for x in buses if x != "x"])

    step = 1
    found = 0
    start = start_timestamp

    while found < bus_ids:
        for i in itertools.count(start, step):
            # print(i)
            matches = tuple(bus for offset, bus in enumerate(buses, start=i) if bus != 'x' and offset % bus == 0)

            if len(matches) > found:
                start = i
                step = math.lcm(*matches)
                found = len(matches)
                break

    return start


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    print("Main result: ", main(rows[1].split(","), start_timestamp=100000000000000))

