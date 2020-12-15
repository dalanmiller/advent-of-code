from typing import List, Tuple
import math
import sys
import re
from functools import reduce

test_case = """0,3,6"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows[0])
    print(result)
    assert result == 436

    numbers = "1,3,2"
    result = main(numbers)
    print(result)
    assert result == 1

    numbers = "2,1,3"
    result = main(numbers)
    print(result)
    assert result == 10

    numbers = "1,2,3"
    result = main(numbers)
    print(result)
    assert result == 27

    numbers = "2,3,1"
    result = main(numbers)
    print(result)
    assert result == 78

    numbers = "3,2,1"
    result = main(numbers)
    print(result)
    assert result == 438

    numbers = "3,1,2"
    result = main(numbers)
    print(result)
    assert result == 1836
    
    print("Tests complete.")

instruction_pattern = re.compile(r"mem\[(\d*)\] = (\d*)")
def parse_line(row: str) -> List[str]:
    return [int(x) for x in row.split(",")]

def main(rows: List[List[str]]) -> int:
    
    numbers = parse_line(rows)
    spoken = {n: [i] for i, n in enumerate(numbers, start=1)}

    number = numbers[-1]
    print("starting number ", number)
    for turn in range(len(numbers)+1, 2021):
    
        # spoken for first time
        # => say 0
        # spoken for sexond time
        # => say difference between turn and orig turn

        if number in spoken and len(spoken[number]) == 1:
            # print(turn, number, spoken[number])
            number = 0    
        elif number in spoken and len(spoken[number]) > 1:
            # print(turn, number, spoken[number])
            number = spoken[number][-1] - spoken[number][-2]

        if number in spoken:
            spoken[number] = [spoken[number][-1], turn]
        else:
            spoken[number] = [turn]
        
        if turn == 2020:
            return number

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows[0]))
    # 1939887
