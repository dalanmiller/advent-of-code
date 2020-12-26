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
    assert result == 112
    print("Tests complete.")


def test_num_adjacent_active():

    coord = (1, 1, 0)
    new_space: Set[Tuple[int, int, int]] = {(0, 0, 0), (0, 1, 0), (1, 0, 0), (2, 2, 0)}
    result = num_adjacent_active(new_space, coord)
    assert result == 4

    coord = (1, 1, 0)
    new_space = {(0, 0, 0), (0, 1, 0), (1, 0, 0), (2, 2, 0), (1, 1, -1), (1, 1, 1)}
    result = num_adjacent_active(new_space, coord)
    assert result == 6


def test_iterate_cube_space():

    test = """#.#
...
#.."""

    space = parse_input(test.split("\n"))
    r_space = iterate_cube_space(space)

    assert (1, 1, 0) in r_space
    # assert r_values[(1, 1, 0)] == "#"

    test = """#.#
.#.
#.#"""

    space = parse_input(test.split("\n"))
    r_space = iterate_cube_space(space)

    assert (1, 0, 0) in r_space
    assert (0, 1, 0) in r_space
    assert (2, 1, 0) in r_space
    assert (1, 2, 0) in r_space

    test = """.#.
..#
###"""

    # import pdb; pdb.set_trace()
    space = parse_input(test.split("\n"))
    result = num_adjacent_active(space, (2, 1, 0))
    assert result == 3
    result = num_adjacent_active(space, (2, 1, 1))
    assert result == 4
    # import pdb; pdb.set_trace()
    result = num_adjacent_active(space, (2, 0, 1))
    print(space, result)
    assert result == 3

    r_space = iterate_cube_space(space)
    print(r_space)

    assert (1, 0, 0) in r_space
    assert (2, 1, 0) in r_space
    assert (2, 0, 0) in r_space
    assert (0, 1, 1) in r_space
    assert (0, 1, -1) in r_space


def parse_input(rows) -> Set[Tuple[int, int, int]]:
    coords: Set[Tuple[int, int, int]] = set()

    rows.reverse()
    for j_space, x in enumerate(rows): 
        for i_space, y in enumerate(x):
            if y == "#":
                coords.add((i_space, j_space, 0))
    return coords


adjacent_shifts = {x for x in itertools.product([-1, 0, 1], repeat=3)}
adjacent_shifts.remove((0, 0, 0))

def num_adjacent_active(
    space: Set[Tuple[int, int, int]],
    coord: Tuple[int, int, int],
) -> int:
    cx, cy, cz = coord
    adjacent: List[Tuple[int, int, int]] = [
        (x + cx, y + cy, z + cz) for x, y, z in adjacent_shifts
    ]

    return len([x for x in adjacent if x in space])


def adjacent_coordinates(coord: Tuple[int, int, int]) -> List[Tuple[int, int, int]]:
    x, y, z = coord
    return [(x + a, y + b, z + c) for a, b, c in adjacent_shifts]


def iterate_cube_space(space: Set[Tuple[int, int, int]]) -> Set[Tuple[int, int, int]]:
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
    space: Set[Tuple[int, int, int]] = parse_input(rows)

    for i in range(6):
        print("I:", i)
        space = iterate_cube_space(space)

    return len(space)


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test_num_adjacent_active()
    test_iterate_cube_space()
    test()
    print("Main result: ", main(rows))
