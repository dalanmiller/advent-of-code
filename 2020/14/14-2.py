from typing import List, Tuple
import math
import sys
import re
from functools import reduce
import itertools

test_case = """mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 208
    print("Tests complete.")

mask_pattern = re.compile(r"^mask \= ([01X]*)")
def parse_mask_line(row: str) -> List[str]:
    result = re.match(mask_pattern, row.strip()).groups()
    mask = list(result[0])
    return [(i, x) for i, x in enumerate(mask)]

instruction_pattern = re.compile(r"mem\[(\d*)\] = (\d*)")
def parse_instruction_line(row: str) -> Tuple[int, int]:
    result = re.match(instruction_pattern, row.strip()).groups()
    mem, value = int(result[0]), int(result[1])
    return mem, value

def apply_bitmask(memory_address: int, mask: List[Tuple[int, int]]) -> List[int]:
    address_bit_string = list(format(memory_address, 'b').zfill(36))
    
    x_locations = []
    for mem, value in mask:
        if value == "X":
            x_locations.append(mem)
        elif value == "1":
            address_bit_string[mem] = value

    possible_addresses = [] 
    combos = [x for x in itertools.product('01', repeat=len(x_locations))]
    
    for combo in combos:
        temp = address_bit_string.copy()
        for val, loc in zip(combo, x_locations):
            temp[loc] = val
        
        possible_addresses.append(temp)
    
    return [int("".join(x), base=2) for x in possible_addresses]

def main(rows: List[List[str]]) -> int:
    mask = None
    memory = {}

    for row in rows: 
        if re.match(mask_pattern, row):
            mask = parse_mask_line(row)
        elif re.match(instruction_pattern, row):
            mem, value = parse_instruction_line(row)
            addresses = apply_bitmask(mem, mask)
            for address in addresses:
                memory[address] = value
    
    return sum(memory.values())

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
