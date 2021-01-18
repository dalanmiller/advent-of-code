from dataclasses import dataclass
import math
import re
import collections
from typing import List, Dict, Any, Tuple, Set, Optional
import numpy as np

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
        self.edges()

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
        self.edges()
        self.trim()
        return self

    def flip_horiz(self):
        self.tile = np.fliplr(self.tile)
        self.edges()
        self.trim()
        return self

    def flip_vert(self):
        self.tile = np.flip(self.tile, axis=0)
        self.edges()
        self.trim()
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

        joined_edges: Dict[str, str] = {k: "".join(v) for k, v in edges.items()}

        self.top = joined_edges["n"]
        self.right = joined_edges["e"]
        self.left = joined_edges["w"]
        self.bottom = joined_edges["s"]
        self.borders = set([self.top, self.right, self.bottom, self.left])

    def generate_arrs(self):
        # start
        # flip horiz
        # flip vert
        # 1 cw
        # 1 cw, flip horiz
        # 1 cw, flip vert
        # 2 cw
        # 3 cw

        for arr in [
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
        ]:
            yield arr


def test_rotations():
    tile = np.full([3, 3], [("1", "2", "3"), ("4", "5", "6"), ("7", "8", "9")])
    t = T("1", tile)

    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["7", "4", "1"], ["8", "5", "2"], ["9", "6", "3"]])
    ).all()
    assert t.top == "123"
    assert t.right == "963"
    assert t.bottom == "789"
    assert t.left == "741"
    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["9", "8", "7"], ["6", "5", "4"], ["3", "2", "1"]])
    ).all()
    assert t.top == "741"
    assert t.right == "321"
    assert t.bottom == "963"
    assert t.left == "987"
    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["3", "6", "9"], ["2", "5", "8"], ["1", "4", "7"]])
    ).all()
    assert t.top == "987"
    assert t.right == "147"
    assert t.bottom == "321"
    assert t.left == "369"
    t.rotate(1)
    assert (
        t.tile == np.full([3, 3], [["1", "2", "3"], ["4", "5", "6"], ["7", "8", "9"]])
    ).all()
    assert t.top == "369"
    assert t.right == "789"
    assert t.bottom == "147"
    assert t.left == "123"


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


def test_find_sea_monsters():
    test_case = """.####...#####..#...###.....
#####..#..#.#.####..#.#....
.#.#...#.###...#.##.##.....
#.#.##.###.#.##.##.#####...
..##.###.####..#.####.##...
...#.#..##.##...#..#..##...
#.##.#..#.#..#..##.#.#.....
.###.##.....#...###.#......
#.####.#.#....##.#..#.#....
##...#..#....#..#...####...
..#.##...###..#.#####..#...
....#.##.#.#####....#......
..##.##.###.....#.##..#....
#...#...###..####....##....
.#.##...#.##.#.#.###...#...
#.###.#..####...##..#......
#.###...#.##...#.######....
.###.###.#######..#####....
..##.#..#..#.#######.###...
#.#..##.########..#..##....
#.#####..#.#...##..#.......
#....##..#.#########..##...
#...#.....#..##...###.##...
#..###....##.#...##.##.#...
## #.###...#.##...#.######.
#####.###.#######.######...
#.#####.##.##.#######.###..
..........................."""

    rows = [list(x) for x in test_case.split("\n")]
    rows.reverse()
    joined_rows = ["".join(x) for x in rows]

    result = find_sea_monsters(joined_rows)
    assert result == 3


def assert_borders_match(square: List[List[T]]):
    l: int = len(square)
    sides = ((1, 0), (0, 1), (-1, 0), (0, -1))
    compared: Set[Tuple[int, int, int, int]] = set()

    # Assert all edges match appropriately
    for i in range(len(square)):
        for j in range(len(square)):
            adjacent_tiles: List[Tuple[int, int]] = [(i + x, j + y) for x, y in sides]
            valid_tiles = [
                (x, y)
                for (x, y) in adjacent_tiles
                if 0 <= x < l and 0 <= y < l and (i, j, x, y) not in compared
            ]
            for x, y in valid_tiles:
                if not square[x][y].borders & square[i][j].borders:
                    return False
                compared.add((i, j, x, y))

    return True


def find_sea_monsters(square: List[str]) -> int:
    # #.###...#.##...#.##O###.
    # .O##.#OO.###OO##..OOO##.
    # ..O#.O..O..O.#O##O##.###

    # Couldn't get lookahead searching to work as some of the sea monsters
    #  are overlapping. I'll save this for another day.

    # monster_upper_pattern = re.compile(r"#")
    # monster_middle_pattern = re.compile(r"(?=(#[\.\#]{4}##[\.\#]{4}##[\.\#]{4}###))")
    # monster_lower_pattern = re.compile(
    #     r"(?=(#[\.\#]{2}#[\.\#]{2}#[\.\#]{2}#[\.\#]{2}#[\.\#]{2}#))"
    # )

    sea_monsters: int = 0

    # matches = re.finditer(monster_middle_pattern, row)

    # for match in matches:

    #     begin, end = match.span()

    #     if i - 1 >= 0 and i + 1 < len(square):
    #         # Values reversed (above looking forward and lower looking backwards)
    #         #  because we are iterating over the list of strings in reverse
    #         upper_matches = re.finditer(monster_upper_pattern, square[i + 1])
    #         lower_matches = re.finditer(monster_lower_pattern, square[i - 1])

    #         # head location
    #         upper_found = False
    #         for m in upper_matches:
    #             if m.span()[0] == end - 1:
    #                 upper_found = True
    #                 break

    #         # bottom location
    #         lower_found = False
    #         for m in lower_matches:
    #             if m.span()[0] == begin + 1:
    #                 lower_found = True
    #                 break

    #         if upper_found and lower_found:
    #             sea_monsters += 1

    for i, row in enumerate(square):

        # Need an inner row to properly match
        if i in (0, len(square) - 1):
            continue

        for j, char in enumerate(row):
            if char == "#" and j + 19 < len(row):

                if all(
                    [
                        row[j] == "#",
                        row[j + 5] == "#",
                        row[j + 6] == "#",
                        row[j + 11] == "#",
                        row[j + 12] == "#",
                        row[j + 17] == "#",
                        row[j + 18] == "#",
                        row[j + 19] == "#",
                        # upper
                        square[i + 1][j + 18] == "#",
                        # lower
                        square[i - 1][j + 1] == "#",
                        square[i - 1][j + 4] == "#",
                        square[i - 1][j + 7] == "#",
                        square[i - 1][j + 10] == "#",
                        square[i - 1][j + 13] == "#",
                        square[i - 1][j + 16] == "#",
                    ]
                ):

                    sea_monsters += 1

    return sea_monsters


def test_generate_tile_arrangements():
    tile = np.full([3, 3], [("1", "2", "3"), ("4", "5", "6"), ("7", "8", "9")])
    t = T("1", tile)
    arrs = [x for x in t.generate_arrs()]

    assert (tile == arrs[0].tile).all()
    assert (
        np.full([3, 3], [["3", "2", "1"], ["6", "5", "4"], ["9", "8", "7"]])
        == arrs[1].tile
    ).all()
    assert (
        np.full([3, 3], [["7", "8", "9"], ["4", "5", "6"], ["1", "2", "3"]])
        == arrs[2].tile
    ).all()
    assert (
        np.full([3, 3], [["7", "4", "1"], ["8", "5", "2"], ["9", "6", "3"]])
        == arrs[3].tile
    ).all()
    assert (
        np.full([3, 3], [["9", "6", "3"], ["8", "5", "2"], ["7", "4", "1"]])
        == arrs[4].tile
    ).all()
    assert (
        np.full([3, 3], [["1", "4", "7"], ["2", "5", "8"], ["3", "6", "9"]])
        == arrs[5].tile
    ).all()
    assert (
        np.full([3, 3], [["9", "8", "7"], ["6", "5", "4"], ["3", "2", "1"]])
        == arrs[6].tile
    ).all()
    assert (
        np.full([3, 3], [["3", "6", "9"], ["2", "5", "8"], ["1", "4", "7"]])
        == arrs[7].tile
    ).all()


def main(rows: List[str]) -> int:
    tiles = parse_input(rows)

    matches: Dict[str, Set[T]] = collections.defaultdict(set)

    # generate all permutations of tiles
    full_set: List[T] = []
    for tile in tiles.values():
        full_set.extend(tile.generate_arrs())

    matches: Dict[str, Set[T]] = collections.defaultdict(set)
    for left_tile in full_set:
        for right_tile in full_set:
            if left_tile.id == right_tile.id:
                continue

            if len(left_tile.borders & right_tile.borders):
                matches[left_tile.id].add(right_tile)
                matches[right_tile.id].add(left_tile)

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

        # Wrap up for next iteration
        prev_tile: Optional[T] = square[x][y]
        available_tiles.remove(square[x][y])

        if x == sq - 1:
            x, y = 0, y + 1
        else:
            x, y = x + 1, y


    # Have complete square, now need to rotate tiles to ensure border alignment
    square_filled = np.full((sq, sq), None)

    init = False
    for left_tile in square[0][0].generate_arrs():
        for right_tile in square[1][0].generate_arrs():
            if left_tile.right == right_tile.left:
                square_filled[0][0] = left_tile
                square_filled[1][0] = right_tile
                init = True
                break

        if init:
            break

    x, y = 0, 0
    prev_tile: Optional[T] = None
    while not np.all(square_filled):

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
            prev_tile_e_edge: str = prev_tile.right

            for right_tile in square[x][y].generate_arrs():
                if prev_tile_e_edge == right_tile.left:
                    square_filled[x][y] = right_tile
                    break

        # If we are on any row above the first one we want to match on the 's' side of the
        # square to match the prior row
        elif y > 0:
            prev_tile = square_filled[x][y - 1]
            prev_tile_n_edge: str = prev_tile.top

            # if sq > 3:
            # import pdb; pdb.set_trace()

            for right_tile in square[x][y].generate_arrs():
                if prev_tile_n_edge == right_tile.bottom:
                    square_filled[x][y] = right_tile
                    break
            else:
                # If we don't break out of prior for loop,
                #  then the current rows need to be flipped
                #  if this is encountered above y = 1, then we
                #  just flip the starting two and restart for simplicity
                if y == 1:
                    for i in range(sq):
                        square_filled[i][0].flip_horiz()
                else:
                    first = square_filled[0][0].flip_horiz()
                    second = square_filled[1][0].flip_horiz()
                    square_filled = np.full((sq, sq), None)
                    square_filled[0][0] = first
                    square_filled[1][0] = second
                    x = 2
                    y = 0

    # Ensure that all borders match
    assert assert_borders_match(square_filled)

    # Concat horizontonally each tile
    square_rows: List[Any] = []
    for i in range(len(square_filled)):
        tile_rows = [x[i].trimmed for x in square_filled]
        square_rows.append(np.concatenate(tile_rows))

    # Concat the rows into a final giant tile square and rotate ccw
    #  because
    final_square = np.rot90(np.concatenate(square_rows, axis=1))

    # Finally will need to try all rotations and flips to search for sea monster
    max_monsters: int = 0
    mod_final_square: T = T(0, final_square.copy())
    final_arrangements = mod_final_square.generate_arrs()
    for i, arr in enumerate(final_arrangements):
        tile = arr.tile
        monsters = find_sea_monsters(["".join(x) for x in tile])

        if monsters > 0:
            with open(f"output-{i}", "w") as o:
                o.writelines(["".join(x) + "\n" for x in tile])

        max_monsters = max(monsters, max_monsters)

    monster_value: int = max_monsters * 15

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
