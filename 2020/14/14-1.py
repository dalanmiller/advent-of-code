from typing import List, Tuple
import math
import sys
import re
from functools import reduce

test_case = """mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 165
    print("Tests complete.")


mask_pattern = re.compile(r"^mask \= ([01X]*)")
def parse_mask_line(row: str) -> List[str]:
    result = re.match(mask_pattern, row.strip()).groups()
    mask = list(result[0])
    return [(i, int(x)) for i, x in enumerate(mask) if x != "X"]

instruction_pattern = re.compile(r"mem\[(\d*)\] = (\d*)")
def parse_instruction_line(row: str) -> Tuple[int, int]:
    result = re.match(instruction_pattern, row.strip()).groups()
    mem, value = int(result[0]), int(result[1])
    return mem, value

def apply_bitmask(value: int, mask: List[Tuple[int, int]]) -> str: 
    byte_string = list(format(value, 'b').zfill(36))
    for pos, value in mask:
        byte_string[pos] = str(value)
    return "".join(byte_string)

def main(rows: List[List[str]]) -> int:
    mask = None
    memory = {}

    for row in rows: 
        if re.match(mask_pattern, row):
            mask = parse_mask_line(row)
        elif re.match(instruction_pattern, row):
            mem, value = parse_instruction_line(row)
            memory[mem] = apply_bitmask(value, mask)
    return sum(int(x, base=2) for x in memory.values())

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
    # 1939887
