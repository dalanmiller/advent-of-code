from collections import Counter
from typing import Tuple, List
import functools
import sys
import itertools
import pprint
import re
import math
import operator

test_case = """2 * 3 + (4 * 5)
1 + (2 * 3) + (4 * (5 + 6))
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"""

# test_case = """1 + (2 * 3) + (4 * (5 + 6))"""


def test():
    rows = [x for x in test_case.split("\n")]

    results = [resolve(parse_line(row)) for row in rows]
    print(results)

    assert results[0] == 46
    assert results[1] == 51
    assert results[2] == 1445
    assert results[3] == 669060
    assert results[4] == 23340
    print("Tests complete.")

def parse_line(line):
    stack = []
    prev = []
    pointer = stack
    for char in line:
        if char == " ":
            continue
        elif re.match(r"\d+", char):
            pointer.append(int(char))
            
        elif char == "+":
            pointer.append(operator.add)
        elif char == "*":
            pointer.append(operator.mul)
        elif char == "(":
            new_frame = []
            pointer.append(new_frame)
            stack.append(new_frame)
            prev.append(pointer)
            pointer = new_frame
        elif char == ")":
            stack.pop()
            pointer = prev[-1]
            prev.pop()

    def parse_dive(frame):
        new_frame = []
        skip = False
        for i, token in enumerate(frame):
            if skip:
                skip = False
                continue

            if type(token) == list:
                new_frame.append(parse_dive(token))
            elif token == operator.add:
                prev = new_frame.pop()
                if type(frame[i+1]) == list:
                    new_frame.append([prev, token, parse_dive(frame[i+1])])
                else:
                    new_frame.append([prev, token, frame[i+1]])
                skip = True
            else:
                new_frame.append(token)

        return new_frame

    
    ordered_stack = parse_dive(stack)
                
    return ordered_stack


def resolve(frame):

    operator = None
    operands = []
    print(frame, type(frame))
    for token in frame:
        # print(token)
        if type(token) == int:
            operands.append(token)
        elif type(token) == list:
            operands.append(resolve(token))
        else:
            operator = token

        if len(operands) == 2:
            x, y = operands
            operands = [operator(x,y)]
            # print(x,y, operator, operands)

    return operands[0]


def main(rows):

    result = 0
    for row in rows:
        stack = parse_line(row)
        print(stack)
        result += resolve(stack)

    return result


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
