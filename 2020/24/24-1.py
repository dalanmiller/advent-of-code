import re
import functools
import itertools
import math
import collections
from dataclasses import dataclass
from typing import List, Optional, Tuple, Literal, Dict, Callable

test_case = """sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew"""


def test():
    rows = test_case.split("\n")

    current_location = (0.0, 0.0)
    current_location = move("ne", current_location)
    current_location = move("nw", current_location)
    current_location = move("sw", current_location)
    current_location = move("se", current_location)
    assert current_location == (0.0, 0.0)

    current_location = (0.0, 0.0)
    current_location = move("se", current_location)
    current_location = move("sw", current_location)
    current_location = move("nw", current_location)
    current_location = move("ne", current_location)
    assert current_location == (0.0, 0.0)

    current_location = (0.0, 0.0)
    current_location = move("e", current_location)
    current_location = move("se", current_location)
    current_location = move("w", current_location)
    current_location = move("w", current_location)
    current_location = move("w", current_location)
    current_location = move("ne", current_location)
    current_location = move("e", current_location)
    assert current_location == (0.0, 0.0)

    matches = main(rows)
    print(matches)
    assert matches == 10

dir_pattern = re.compile(r"(se|ne|nw|sw|[we])")
def parse_hops(rows) -> List[List[str]]:
    all_dirs = []
    for row in rows:
        m = re.findall(dir_pattern, row)
        all_dirs.append(m)
    
    return all_dirs

dir_map: Dict[str, Callable] = {
        "ne": lambda x: (x[0] + 0.5, x[1] + 0.5),
        "e": lambda x: (x[0] + 1.0, x[1]),
        "se": lambda x: (x[0] + 0.5, x[1] - 0.5),
        
        "sw": lambda x: (x[0] - 0.5, x[1] - 0.5),
        "w": lambda x: (x[0] - 1.0, x[1] ),
        "nw": lambda x: (x[0] - 0.5, x[1] + 0.5),
    }
def move(dir: str, current_location: Tuple[float, float]) -> Tuple[float, float]:
    return dir_map[dir](current_location)

def main(rows) -> int:
    hops: List[List[str]] = parse_hops(rows)
    
    visited_tiles = {}

    for hop in hops:
        current_location: Tuple[float, float] = (0.0, 0.0) 
        for dir in hop:
            current_location = move(dir, current_location)

        if current_location in visited_tiles:
            del visited_tiles[current_location]
        else:
            visited_tiles[current_location] = 1

        print(hop, current_location)

    print(visited_tiles)
    return sum(visited_tiles.values())


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
