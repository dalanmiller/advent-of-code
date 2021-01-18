import re
import functools
import itertools
import math
import collections
from typing import List

test_case = """389125467"""

def test():
    rows = test_case.split("\n")

    matches = main(rows[0])
    print(matches)
    assert matches == "67384529"


def parse_cups(row) -> List[int]:
    return [int(x) for x in row]


def main(row) -> str:
    cups = collections.deque()
    cups.extend(parse_cups(row))
    cups.reverse()

    current_cup = cups.pop()
    for move in range(1, 101):

        # grab three cups clockwise from current cup
        pick_up_cups = [cups.pop(), cups.pop(), cups.pop()]

        i = current_cup - 1
        destination: int = -1
        while destination == -1:
            if i in cups:
                destination = cups.index(i)
            else:
                i -= 1

            if i < 1:
                i = max(cups)
        
        for cup in pick_up_cups:
            cups.insert(destination, cup)

        temp_cup = current_cup
        current_cup = cups.pop()
        cups.append(temp_cup)
        cups.rotate(1)

    cups.append(current_cup)
    one = cups.index(1)
    cups.rotate(-one)
    cups.reverse()
    cups.pop()
        
    return "".join([str(x) for x in cups])

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows[0]))
