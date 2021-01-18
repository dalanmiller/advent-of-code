from collections import Counter
from typing import Tuple, List
import functools
import sys
import itertools
import pprint
import re
import math

test_case = """2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"""

# test_case = """((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"""


def test():
    rows = [x for x in test_case.split("\n")]

    results = [resolve(parse_line(row)) for row in rows]
    print(results)

    assert results[0] == 26
    assert results[1] == 437
    assert results[2] == 12240
    assert results[3] == 13632
    print("Tests complete.")

    parse_line("5 * (4 * 3 * 4 * 2) + 8")


def parse_line(line):
    buffer = []
    stack = [buffer]
    pointer = buffer
    for char in line:
        if char == " ":
            continue
        elif re.match(r"\d", char):
            pointer.append(int(char))
        elif char == "+":
            pointer.append(sum)
        elif char == "*":
            pointer.append(math.prod)
        elif char == "(":
            new_frame = []
            pointer.append(new_frame)
            stack.append(new_frame)
            pointer = new_frame
        elif char == ")":
            stack.pop()
            pointer = stack[-1]

    return stack


def resolve(frame):

    operator = None
    operands = []
    for token in frame:
        if type(token) == int:
            operands.append(token)
        elif type(token) == list:
            operands.append(resolve(token))
        else:
            operator = token

        if len(operands) == 2:
            operands = [operator(operands)]

    return operands[0]


def main(rows):

    result = 0
    for row in rows:
        stack = parse_line(row)
        result += resolve(stack[0])

    return result


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
