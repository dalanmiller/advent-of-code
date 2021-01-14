from functools import lru_cache
import itertools
from dataclasses import *
import math
import sys
import re
import collections
from typing import List, Dict, Any, Tuple, Set
import numpy as np
from numpy.core.defchararray import join
from numpy.core.numeric import full
from numpy.lib.npyio import recfromtxt

test_case = """Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###..."""


def test():
    lines = test_case.split("\n")
    result = main(lines)
    print("Test result: ", result)
    assert result == 273


@dataclass
class T:
    id: str
    tile: Any

    def __init__(self, id, tile):
        self.id = id
        self.tile = tile
        self.trim()

    def __repr__(self):
        return f"<T id={self.id} size={self.tile.size}"

    def __hash__(self):
        return hash(self.id)

    def __eq__(self, o):
        return self.id == o.id

    def __gt__(self, o):
        return self.id > o.id

    def generate_text(self):
        rows = []
        for i in range(len(self.tile[0]), 0):
            rows.append("".join([x[i] for x in self.tile]))
        return rows

    def rotate(self, r: int):
        self.tile = np.rot90(self.tile, k=r, axes=(1, 0))
        return self

    def flip_horiz(self):
        self.tile = np.fliplr(self.tile)
        return self

    def flip_vert(self):
        self.tile = np.flip(self.tile, axis=0)
        return self

    def copy(self):
        return T(self.id, self.tile)

    def trim(self):    
        l = len(self.tile) - 1
        # remove top and bottom
        trimmed_tile = np.delete(self.tile, [0, l], axis=1)
        # remove left and right
        trimmed_tile = np.delete(trimmed_tile, [0, l], axis=0)

        self.trimmed = trimmed_tile
        return self

    def side(self, side):
        edges: Tuple[str] = self.edges()
        m: Dict[str, int] = {"n": 0, "e": 1, "s": 2, "w": 3}
        return edges[m[side]]

    def edges(self):
        edges: Dict[str, List[str]] = {}
        l: int = len(self.tile)

        # starting from the bottom up
        edges["s"] = [x[0] for x in self.tile]  # bottom edge
        edges["n"] = [x[-1] for x in self.tile]  # top edge
        for i, column in enumerate(self.tile):
            if i == 0:
                edges["w"] = column.tolist()  # left edge

            if i == l - 1:
                edges["e"] = column.tolist()  # right edge

        # edges['s'] = [x for x in reversed(list(edges['s']))]
        # edges['e'] = [x for x in reversed(edges['e'])]
        joined_edges: Dict[str, str] = {k: "".join(v) for k, v in edges.items()}
        return tuple(
            [joined_edges["n"], joined_edges["e"], joined_edges["s"], joined_edges["w"]]
        )

    def generate_arrs(self):
        # start
        # flip horiz
        # flip vert
        # 1 cw
        # 1 cw, flip horiz
        # 1 cw, flip vert
        # 2 cw
        # 3 cw

        return [
            self,  # 0
            # np.fliplr(tile),  # 1 -- Technically flipping along 'columns'
            self.copy().flip_horiz(),
            # np.flip(tile, axis=0),  # 2 -- flipping along rows
            self.copy().flip_vert(),
            # np.rot90(tile, axes=(1, 0)),  # 3
            self.copy().rotate(1),
            # np.fliplr(np.rot90(tile, axes=(1, 0))),  # 4
            self.copy().flip_horiz().rotate(1),
            # np.flip(np.rot90(tile, axes=(1, 0)), axis=0),  # 5
            self.copy().flip_vert().rotate(1),
            # np.rot90(tile, k=2, axes=(1, 0)),  # 6
            self.copy().rotate(2),
            # np.rot90(tile, k=3, axes=(1, 0)),  # 7
            self.copy().rotate(3),
        ]


# def compute_edges(image_data) -> Dict[str, Any]:
#     edges: Dict[str, Any] = {}
#     left: List[str] = []
#     core: List[List[str]] = []
#     right: List[str] = []
#     core_full: List[List[str]] = []

#     for i, row in enumerate(image_data):
#         if i == 0:
#             edges["n"] = row  # top edge

#         if i == len(image_data) - 1:
#             edges["s"] = row  # bottom edge

#         if i not in (0, len(image_data) - 1):
#             core.append(list(row[1:-1]))

#         left.append(row[0])
#         right.append(row[-1])
#         core_full.append(list(row))

#     # core.reverse()
#     edges["w"] = "".join(left)
#     edges["e"] = "".join(right)
#     edges["core"] = np.array(core)
#     edges["core_full"] = np.array(core_full)

#     full_edges = edges.copy()
#     for k, v in edges.items():
#         if k.startswith("core"):
#             continue
#         inv = list(v)
#         inv.reverse()
#         full_edges[f"{k}_i"] = "".join(inv)

#     return full_edges


def test_rotations():
    tile = np.full([3, 3], [("1", "2", "3"), ("4", "5", "6"), ("7", "8", "9")])
    t = T("1", tile)
    edges: Set[str] = set(t.edges())

    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["7", "4", "1"], ["8", "5", "2"], ["9", "6", "3"]])
    ).all()
    assert t.edges() == ("123", "963", "789", "741")
    assert t.side("n") == "123"
    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["9", "8", "7"], ["6", "5", "4"], ["3", "2", "1"]])
    ).all()
    assert t.edges() == ("741", "321", "963", "987")
    assert t.side("n") == "741"
    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["3", "6", "9"], ["2", "5", "8"], ["1", "4", "7"]])
    ).all()
    assert t.edges() == ("987", "147", "321", "369")
    assert t.side("n") == "987"
    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["1", "2", "3"], ["4", "5", "6"], ["7", "8", "9"]])
    ).all()
    assert t.edges() == ("369", "789", "147", "123")
    assert t.side("n") == "369"


def parse_input(input: List[str]) -> Dict[str, T]:
    tiles: Dict[str, T] = {}
    tile_number: str = ""
    image_data: List[str] = []

    for i in input:
        if "Tile" in i:
            tile_number = i.split(" ")[-1][:-1]
        elif i == "\n" or i == "":
            tiles[tile_number] = T(tile_number, np.array([list(x) for x in image_data]))

            tile_number, image_data = "", []
        else:
            image_data.append(i)

    tiles[tile_number] = T(tile_number, np.array([list(x) for x in image_data]))

    return tiles


# def test_transform_core():
#     test_core = np.full([3, 3], [(7, 4, 1), (8, 5, 2), (9, 6, 3)])

#     result = transform_core(test_core, "n", "s")
#     comparison = result == np.rot90(test_core, 2)
#     assert comparison.all()

#     result = transform_core(test_core, "n_i", "s")
#     test = np.fliplr(np.rot90(test_core, 2))
#     comparison = result == test
#     assert comparison.all()

#     result = transform_core(test_core, "w_i", "e")
#     comparison = result[2] == [3, 6, 9]
#     assert comparison.all()


# dirs = {"n": 0, "e": 1, "s": 2, "w": 3}


# def transform_core(core, edge_side: str, target_dir: str):

#     if edge_side in ("n_i", "s_i"):
#         core = np.fliplr(core)
#     elif edge_side in ("w_i", "e_i"):
#         core = np.rot90(core, axes=(1, 0))  # clockwise
#         core = np.fliplr(core)
#         core = np.rot90(core, axes=(0, 1))  # counter-clockwise

#     rots = determine_cw_rotations(edge_side, target_dir)
#     try:
#         core = np.rot90(core, k=rots, axes=(1, 0))  # clockwise
#     except Exception as e:
#         import pdb

#         pdb.set_trace()

#     return core


# @lru_cache()
# def determine_cw_rotations(current: str, target: str) -> int:
#     dirs_list: List[str] = ["n", "e", "s", "w"]
#     start: int = dirs_list.index(current[0]) + 1
#     end: int = dirs_list.index(target[0]) + 1

#     # [1,2,3,4]
#     # 1, 2 == 1
#     # 2, 1 == -1
#     return end - start


# def normalize_inverted(core_full, match_side: str):
#     if match_side in ("n_i", "s_i"):
#         return np.fliplr(core_full)

#     elif match_side in ("w_i", "e_i"):
#         core_full = np.rot90(core_full)  # counter-clockwise
#         core_full = np.fliplr(core_full)
#         return np.rot90(core_full, axes=(1, 0))


# def get_edge(core, side: str) -> Tuple[str]:
#     try:
#         if side == "n":
#             return tuple([x[len(core) - 1] for x in core])
#         elif side == "e":
#             return tuple(core[len(core) - 1])
#         elif side == "w":
#             return tuple(core[0])
#         elif side == "s":
#             return tuple([x[0] for x in core])
#     except Exception as e:
#         import pdb

#         pdb.set_trace()

#     return tuple()


# def print_square(square):
#     square_rows = []
#     for i in range(len(square)):
#         rows = [x[i] for x in square]
#         square_rows.append(np.concatenate(rows, axis=1))

#     # reverse so printing is done correctly
#     square_rows.reverse()

#     for row in square_rows:
#         print()
#         for line in row:
#             print("".join(line))


def test_find_sea_monsters():
    test_case = """.####...#####..#...###..
#####..#..#.#.####..#.#.
.#.#...#.###...#.##.##..
#.#.##.###.#.##.##.#####
..##.###.####..#.####.##
...#.#..##.##...#..#..##
#.##.#..#.#..#..##.#.#..
.###.##.....#...###.#...
#.####.#.#....##.#..#.#.
##...#..#....#..#...####
..#.##...###..#.#####..#
....#.##.#.#####....#...
..##.##.###.....#.##..#.
#...#...###..####....##.
.#.##...#.##.#.#.###...#
#.###.#..####...##..#...
#.###...#.##...#.######.
.###.###.#######..#####.
..##.#..#..#.#######.###
#.#..##.########..#..##.
#.#####..#.#...##..#....
#....##..#.#########..##
#...#.....#..##...###.##
#..###....##.#...##.##.#"""

    rows = [list(x) for x in test_case.split("\n")]
    rows.reverse()
    joined_rows = ["".join(x) for x in rows]

    result = find_sea_monsters(joined_rows)
    print("TEST: find_sea_monster result: ", result)
    assert result == 2


monster_upper_pattern = re.compile(r"#")
monster_middle_pattern = re.compile(r"#[\.\#]{4}##[\.\#]{4}##[\.\#]{4}###")
monster_lower_pattern = re.compile(
    r"#[\.\#]{2}#[\.\#]{2}#[\.\#]{2}#[\.\#]{2}#[\.\#]{2}#"
)


def find_sea_monsters(square: List[str]) -> int:
    # #.###...#.##...#.##O###.
    # .O##.#OO.###OO##..OOO##.
    # ..O#.O..O..O.#O##O##.###

    sea_monsters: int = 0
    for i, row in enumerate(square):

        # Need an inner row to properly match
        if i in (0, len(square) - 1):
            continue

        matches = re.finditer(monster_middle_pattern, row)

        for match in matches:
            
            begin, end = match.span()

            if i - 1 >= 0 and i + 1 < len(square):
                # Values reversed (above looking forward and lower looking backwards)
                #  because we are iterating over the list of strings in reverse
                upper_matches = re.finditer(monster_upper_pattern, square[i + 1])
                lower_matches = re.finditer(monster_lower_pattern, square[i - 1])

                # head location
                upper_found = False
                for m in upper_matches:
                    if m.span()[0] == end -1:
                        upper_found = True
                        break

                # bottom location
                lower_found = False
                for m in lower_matches:
                    if m.span()[0] == begin + 1:
                        lower_found = True
                        break
                
                if upper_found and lower_found:
                    import pdb; pdb.set_trace()
                    sea_monsters += 1

    return sea_monsters


# def trim_tile(tile):
#     l = len(tile) - 1
#     # remove top and bottom

#     tile = np.delete(tile, [0, l], axis=1)
#     # remove left and right
#     tile = np.delete(tile, [0, l], axis=0)

#     return tile


def test_generate_tile_arrangements():
    tile = np.full([3, 3], [(1, 2, 3), (4, 5, 6), (7, 8, 9)])
    t = T("1", tile)
    arrs = t.generate_arrs()

    # import pdb; pdb.set_trace()
    assert (tile == arrs[0].tile).all()
    assert (np.full([3, 3], [[3, 2, 1], [6, 5, 4], [9, 8, 7]]) == arrs[1].tile).all()
    assert (np.full([3, 3], [[7, 8, 9], [4, 5, 6], [1, 2, 3]]) == arrs[2].tile).all()
    assert (np.full([3, 3], [[7, 4, 1], [8, 5, 2], [9, 6, 3]]) == arrs[3].tile).all()
    assert (np.full([3, 3], [[9, 6, 3], [8, 5, 2], [7, 4, 1]]) == arrs[4].tile).all()
    assert (np.full([3, 3], [[1, 4, 7], [2, 5, 8], [3, 6, 9]]) == arrs[5].tile).all()
    assert (np.full([3, 3], [[9, 8, 7], [6, 5, 4], [3, 2, 1]]) == arrs[6].tile).all()
    assert (np.full([3, 3], [[3, 6, 9], [2, 5, 8], [1, 4, 7]]) == arrs[7].tile).all()


# def generate_tile_arrangements(tile: T) -> List[T]:
#     # start
#     # flip horiz
#     # flip vert
#     # 1 cw
#     # 1 cw, flip horiz
#     # 1 cw, flip vert
#     # 2 cw
#     # 3 cw

#     arrangements: List[T] = [
#         tile,  # 0
#         np.fliplr(tile),  # 1 -- Technically flipping along 'columns'
#         np.flip(tile, axis=0),  # 2 -- flipping along rows
#         np.rot90(tile, axes=(1, 0)),  # 3
#         np.fliplr(np.rot90(tile, axes=(1, 0))),  # 4
#         np.flip(np.rot90(tile, axes=(1, 0)), axis=0),  # 5
#         np.rot90(tile, k=2, axes=(1, 0)),  # 6
#         np.rot90(tile, k=3, axes=(1, 0)),  # 7
#     ]

#     return arrangements


def main(rows: List[str]) -> int:
    tiles = parse_input(rows)

    matches: Dict[str, Set[T]] = collections.defaultdict(set)

    # import pdb; pdb.set_trace()

    # generate all permutations of tiles
    full_set: List[T] = []
    for tile in tiles.values():
        full_set.extend(tile.generate_arrs())

    matches: Dict[str, Set[T]] = collections.defaultdict(set)
    for left_tile in full_set:
        for right_tile in full_set:
            if left_tile.id == right_tile.id:
                continue

            if len(set(left_tile.edges()) & set(right_tile.edges())):
                matches[left_tile.id].add(right_tile)
                matches[right_tile.id].add(left_tile)

    # print(matches)
    # import pdb; pdb.set_trace()
    # sys.exit(1)

    # for l_tile_number, l_image_data in images.items():
    #     for l_key, l_data in l_image_data.items():
    #         for r_tile_number, r_image_data in images.items():
    #             for r_key, r_data in r_image_data.items():
    #                 if (
    #                     r_tile_number == l_tile_number
    #                     or l_key.startswith("core")
    #                     or r_key.startswith("core")
    #                 ):
    #                     continue
    #                 if l_data == r_data:
    #                     matches[l_tile_number].add(r_tile_number)
    #                     matches[r_tile_number].add(l_tile_number)

    # Matches is dict of tile numbers to a set of Tuple pairs which match
    #  an edge of the titular tile number to another tile.

    sq: int = int(math.sqrt(len(matches.keys())))
    corners: Set[T] = {tiles[k] for k, v in matches.items() if len(v) == 2}
    corner_coords = ((0, sq - 1), (sq - 1, 0), (0, 0), (sq - 1, sq - 1))

    square: List[List[T]] = np.full((sq, sq), None)
    available_tiles: Set[T] = set(tiles.values())

    x, y = 0, 0
    prev_tile = None
    while available_tiles:
        # we're starting off, just select a corner
        if (x, y) == (0, 0):
            square[x][y] = sorted(list(corners))[0]

        # if y > 0, we can use the last available tile of the neighbor on the previous row
        elif y > 0 and square[x][y - 1]:
            square[x][y] = (matches[square[x][y - 1].id] & available_tiles).pop()

        # In a corner, must select tiles with only two matches
        elif (x, y) in corner_coords:
            for tile in matches[prev_tile.id] & corners & available_tiles:
                if len(matches[tile.id]) == 2:
                    square[x][y] = tile
                    break

        # We are on the edge, must select tiles with three matches
        elif x == 0 or x == sq - 1 or y == 0 or y == sq - 1:
            for tile in matches[prev_tile.id] & available_tiles:
                if len(matches[tile.id]) == 3:
                    square[x][y] = tile
                    break

        # Otherwise we are in the interior of the square
        elif 0 < x < sq and 0 < y < sq:
            for tile in set(matches[prev_tile.id]) & available_tiles:
                if len(matches[tile.id]) == 4:
                    square[x][y] = tile
                    break

        else:
            # fuck
            import pdb

            pdb.set_trace()

        # Wrap up for next iteration
        prev_tile: Optional[T] = square[x][y]
        available_tiles.remove(square[x][y])

        if x == sq - 1:
            x, y = 0, y + 1
        else:
            x, y = x + 1, y

    print(square)

    # Have complete square, now need to rotate tiles to ensure row alignment
    square_filled = np.full((sq, sq), None)

    init = False
    for left_tile in square[0][0].generate_arrs():
        for right_tile in square[1][0].generate_arrs():
            for i, left_side in enumerate(left_tile.edges()):
                for j, right_side in enumerate(right_tile.edges()):
                    # print(left_tile.id, left_side, right_tile.id, right_side)
                    # if set(left_side) & right_side:
                    # Want eastern edge for left and western for right
                    # print(left_tile.id, left_tile.side('e'), right_tile.id, right_tile.side('w'), right_tile.edges())
                    # if left_tile.side('e') in right_tile.edges():
                    if left_tile.side("e") == right_tile.side("w"):

                        # while left_tile.side('e') != right_tile.side('w'):
                        #     import pdb; pdb.set_trace()
                        #     right_tile.rotate(1)

                        square_filled[0][0] = left_tile
                        square_filled[1][0] = right_tile
                        init = True
                        break

        if init:
            break
    else:
        # fuck
        import pdb

        pdb.set_trace()

    # Handle first two: (0,0) & (1,0)
    # for left, right in itertools.permutations(
    #     ["n", "e", "s", "w", "n_i", "e_i", "w_i", "s_i"], r=2
    # ):
    #     left_tile = transform_core(images[square[0][0]]["core_full"], left, "e")
    #     right_tile = transform_core(images[square[1][0]]["core_full"], right, "w")

    #     if get_edge(left_tile, "e") == get_edge(right_tile, "w"):
    #         square_filled[0][0] = left_tile
    #         square_filled[1][0] = right_tile
    #         break
    # else:
    #     # wtf
    #     import pdb

    #     pdb.set_trace()

    print("First two set")
    x, y = 0, 0
    prev_tile = None
    while not np.all(square_filled):
        # print(x, y, square_filled)
        # for y in range(0, len(square)):
        #     for x in range(0, len(square)):

        # If already placed, ignore, (0,0) and (1,0)
        if square_filled[x][y] is not None:
            if x == sq - 1:
                x, y = 0, y + 1
            else:
                x, y = x + 1, y

        # If we are on the first row then we want to match on the 'w' side of the
        #  square
        if y == 0:
            prev_tile = square_filled[x - 1][y]
            prev_tile_e_edge = prev_tile.side("e")
            # for dir in ("n", "e", "s", "w", "n_i", "e_i", "w_i", "s_i"):
            #     tile_n = square[x][y]
            #     tile = transform_core(images[square[x][y]]["core_full"], dir, "w")
            #     tile_core = transform_core(images[square[x][y]]["core"], dir, "w")
            #     if get_edge(tile, "w") == prev_tile_e_edge:
            #         square_filled[x][y] = tile
            #         break
            for right_tile in square[x][y].generate_arrs():
                # print(
                #     prev_tile.id,
                #     prev_tile.side("e"),
                #     right_tile.id,
                #     right_tile.side("w"),
                # )
                if prev_tile_e_edge == right_tile.side("w"):
                    square_filled[x][y] = right_tile
                    break
            else:
                import pdb

                pdb.set_trace()
                sys.exit(1)

        # If we are on any row above the first one we want to match on the 's' side of the
        # square to match the prior row
        elif y > 0:
            prev_tile = square_filled[x][y - 1]
            prev_tile_n_edge = prev_tile.side("n")
            # for dir in ("n", "e", "s", "w", "n_i", "e_i", "w_i", "s_i"):
            #     tile_n = square[x][y]
            #     tile = transform_core(images[square[x][y]]["core_full"], dir, "s")
            #     tile_core = transform_core(images[square[x][y]]["core"], dir, "s")
            #     if get_edge(tile, "s") == prev_tile_n_edge:
            #         square_filled[x][y] = tile
            #         break
            for right_tile in square[x][y].generate_arrs():
                # print(
                #     x,
                #     y,
                #     prev_tile.id,
                #     prev_tile.side("n"),
                #     right_tile.id,
                #     right_tile.edges(),
                # )
                if prev_tile_n_edge == right_tile.side("s"):

                    square_filled[x][y] = right_tile
                    break

            # if square_filled[x][y] is None:
            #     # Need to flip the previous row
            #     for i in range(sq):
            #         square_filled[x][y - 1] = np.flip(
            #             square_filled[x][y - 1], axis=1
            #         )

            # prev_tile = square_filled[x][y - 1]
            # prev_tile_n_edge = get_edge(prev_tile, "n")
            # for dir in ("n", "e", "s", "w", "n_i", "e_i", "w_i", "s_i"):
            #     tile_n = square[x][y]
            #     tile = transform_core(images[square[x][y]]["core_full"], dir, "s")
            #     tile_core = transform_core(images[square[x][y]]["core"], dir, "s")
            #     if get_edge(tile, "s") == prev_tile_n_edge:
            #         square_filled[x][y] = tile
            #         break

            # If we've arrived here and we are upwards in the square, then we might need
            # to flip the previou rows?
            # if y > 0:
            #     for i in range(y):
            #         for j in range(sq):
            #             square_filled[j][i].flip_vert()

        # print(x, y, square[x][y])
        # print(square_filled[x][y])

    # Assert all edges match appropriately
    for i in range(len(square_filled)):
        for j in range(len(square_filled)):
            if j < len(square_filled) - 1:
                print(
                    j,
                    i,
                    square_filled[j][i].side("e"),
                    j + 1,
                    i,
                    square_filled[j + 1][i].side("w"),
                )
                assert square_filled[j][i].side("e") == square_filled[j + 1][i].side(
                    "w"
                )

            if i < len(square_filled) - 1:
                print(
                    j,
                    i,
                    square_filled[j][i].side("n"),
                    j,
                    i + 1,
                    square_filled[j][i + 1].side("s"),
                )
                assert square_filled[j][i].side("n") == square_filled[j][i + 1].side(
                    "s"
                )

    # Concat horizontonally each tile
    square_rows = []
    for i in range(len(square_filled) - 1, -1, -1):
        import pdb; pdb.set_trace()      
        tile_rows = [x[i].trimmed for x in square_filled]
        square_rows.append(np.concatenate(tile_rows))

    
    # Concat the rows into a final giant tile square
    final_square = np.concatenate(square_rows, axis=1)

    # Finally will need to try all rotations and flips to search for sea monster
    monsters: int = 0
    mod_final_square: T = T(0, final_square.copy())
    final_arrangements: List[T] = mod_final_square.generate_arrs()
    for i, arr in enumerate(final_arrangements):
        tile = arr.tile
        monsters = find_sea_monsters(["".join(x) for x in tile])

        if monsters > 0:
            with open(f"output-{i}", "w") as o:
                o.writelines(["".join(x) + "\n" for x in tile])
            break

    print("Monsters found: ", monsters)
    monster_value: int = monsters * 14

    hashes: int = 0
    for row in final_square:
        hashes += sum(1 for x in row if x == "#")

    not_sea_monster: int = hashes - monster_value

    return not_sea_monster


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    # test_transform_core()
    test_find_sea_monsters()
    test_generate_tile_arrangements()
    test_rotations()
    test()
    print("Main result: ", main(rows))
