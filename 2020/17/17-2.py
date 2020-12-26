from collections import Counter
from typing import Tuple, List, Any, Dict, Set
import functools
import operator
import sys
import itertools
import pprint

test_case = """.#.
..#
###"""


def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 848
    print("Tests complete.")


def parse_input(rows) -> Set[Tuple[int, int, int, int]]:
    coords: Set[Tuple[int, int, int, int]] = set()

    rows.reverse()
    for j_space, x in enumerate(rows): 
        for i_space, y in enumerate(x):
            if y == "#":
                coords.add((i_space, j_space, 0, 0))
    return coords


adjacent_shifts = {x for x in itertools.product([-1, 0, 1], repeat=4)}
adjacent_shifts.remove((0, 0, 0, 0))

def num_adjacent_active(
    space: Set[Tuple[int, int, int, int]],
    coord: Tuple[int, int, int, int],
) -> int:
    dx, dy, dz, dw  = coord
    adjacent: List[Tuple[int, int, int, int]] = [
        (x + dx, y + dy, z + dz, w + dw) for x, y, z, w in adjacent_shifts
    ]

    return len([x for x in adjacent if x in space])


def adjacent_coordinates(coord: Tuple[int, int, int, int]) -> List[Tuple[int, int, int, int]]:
    x, y, z, w = coord
    return [(x + a, y + b, z + c, w + d) for a, b, c, d in adjacent_shifts]


def iterate_cube_space(space: Set[Tuple[int, int, int, int]]) -> Set[Tuple[int, int, int, int]]:
    new_space = space.copy()

    adjacent_coords = set(functools.reduce(
        operator.iconcat, [adjacent_coordinates(x) for x in space]
    ))

    for coord in set(space).union(adjacent_coords):
        adjacent = num_adjacent_active(space, coord)

        if coord in space:
            if coord in new_space and adjacent != 2 and adjacent != 3:
                new_space.remove(coord)
        elif adjacent == 3:
            new_space.add(coord)

    return new_space


def main(rows: List[str]):
    print(rows)
    space: Set[Tuple[int, int, int, int]] = parse_input(rows)

    for i in range(6):
        print("I:", i)
        space = iterate_cube_space(space)

    return len(space)


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]
        
    test()
    print("Main result: ", main(rows))
