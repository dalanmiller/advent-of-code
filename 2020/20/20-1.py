import re
import functools
import itertools
import math
import collections

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

test_final_image = """.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###"""


def test():
    rows = test_case.split("\n")

    matches = main(rows)
    print(matches)
    assert matches == 20899048083289

    # image = construct_image(rows)


def compute_edges(image_data):
    edges = {}
    left = []
    right = []

    for i, row in enumerate(image_data):
        if i == 0:
            edges["n"] = row  # top edge
        elif i == len(image_data) - 1:
            edges["s"] = row  # bottom edge

        left.append(row[0])
        right.append(row[-1])

    edges["w"] = "".join(left)
    edges["e"] = "".join(right)

    full_edges = edges.copy()
    for k, v in edges.items():
        inv = list(v)
        inv.reverse()
        full_edges[f"{k}_i"] = "".join(inv)

    return full_edges


def parse_input(input):
    images = {}
    tile_number = None
    image_data = []
    for i in input:
        if "Tile" in i:
            tile_number = i.split(" ")[-1][:-1]
            print(tile_number)
        elif i == "\n" or i == "":
            images[tile_number] = image_data
            tile_number, image_data = None, []
        else:
            image_data.append(i)

    # finally
    # import pdb; pdb.set_trace()
    images[tile_number] = image_data

    images = {k: compute_edges(v) for k, v in images.items()}
    # import pdb; pdb.set_trace()
    return images


def adjacent_squares(ix, iy, square):
    potential_coords = [
        tuple([int(x[0]), int(x[1])])
        for x in itertools.product(["-1", "0", "1"], repeat=2)
    ]

    # Remove diagnol and center
    potential_coords.remove(tuple([0, 0]))
    potential_coords.remove(tuple([1, 1]))
    potential_coords.remove(tuple([-1, -1]))
    potential_coords.remove(tuple([-1, 1]))
    potential_coords.remove(tuple([1, -1]))

    shifted_coords = [tuple([x[0] + ix, x[1] + iy]) for x in potential_coords]

    valid_coords = []
    for x, y in shifted_coords:
        if 0 <= x < len(square) and 0 <= y < len(square):
            valid_coords.append(tuple([x, y]))

    return valid_coords


def position_type(ix, iy, square):
    # import pdb; pdb.set_trace()

    if ix == 0 and 0 < iy < len(square) - 1 or iy == 0 and 0 < ix < len(square) - 1:
        return 3
    elif (ix, iy) in [
        (0, 0),
        (len(square) - 1, 0),
        (0, len(square) - 1),
        (len(square) - 1, len(square) - 1),
    ]:
        return 2
    else:
        return 4


def validate_square(square, matches):
    possibly_valid = collections.defaultdict(set)
    for i, x in enumerate(square):
        for j, y in enumerate(x):
            adj = adjacent_squares(i, j, square)
            adjacent_images = tuple(square[x][y] for x, y in adj)
            # import pdb; pdb.set_trace()
            current_image_edges = matches[y]

            for adj_image in adjacent_images:
                # print(adj_image)
                if current_image_edges & matches[adj_image]:
                    possibly_valid[y].add(tuple([y, adj_image]))

            if len(possibly_valid) < position_type(i, j, square):
                return False

            possibly_valid = collections.defaultdict(set)

    return True


def main(rows) -> int:
    images = parse_input(rows)

    print(images)

    matches = collections.defaultdict(set)
    for l_tile_number, l_image_data in images.items():
        for l_dir, data in l_image_data.items():
            for r_tile_number, r_image_data in images.items():
                for r_dir, r_data in r_image_data.items():
                    if r_tile_number == l_tile_number:
                        continue
                    if data == r_data:
                        matches[l_tile_number].add(r_tile_number)
                        matches[r_tile_number].add(l_tile_number)

    product = 1
    for k, v in matches.items():
        if len(v) == 2:
            product *= int(k)

    return product

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
