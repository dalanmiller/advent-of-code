from functools import lru_cache
import itertools
import math
import sys
import collections
from typing import List, Dict, Any, Tuple, Set
import numpy as np
from numpy.testing._private.utils import integer_repr

def compute_edges(image_data) -> Dict[str, Any]:
    edges: Dict[str, Any] = {}
    left: List[str] = []
    core: List[List[str]] = []
    right: List[str] = []
    core_full = []

    for i, row in enumerate(image_data):
        if i == 0:
            edges["n"] = row  # top edge
        elif i == len(image_data) - 1:
            edges["s"] = row  # bottom edge

        left.append(row[0])
        core.append(row[1:-1])
        right.append(row[-1])
        core_full.append(list(row))

    core.reverse()
    edges["w"] = "".join(left)
    edges["e"] = "".join(right)
    edges["core"] = np.array(core)
    edges["core_full"] = np.array(core_full)

    full_edges = edges.copy()
    for k, v in edges.items():
        if k.startswith("core"):
            continue
        inv = list(v)
        inv.reverse()
        full_edges[f"{k}_i"] = "".join(inv)

    return full_edges


def parse_input(input: List[str]):
    images: Dict[Any, List[str]] = {}
    tile_number: Any = None
    image_data: List[str] = []
    for i in input:
        if "Tile" in i:
            tile_number = i.split(" ")[-1][:-1]
        elif i == "\n" or i == "":
            images[tile_number] = image_data
            tile_number, image_data = None, []
        else:
            image_data.append(i)

    images[tile_number] = image_data

    return {k: compute_edges(v) for k, v in images.items()}


dirs = {"n": 0, "e": 1, "s": 2, "w": 3}
def transform_core(core, edge_side: str, target_dir: str):

    if edge_side in ("n_i", "s_i"):
        core = np.fliplr(core)
    elif edge_side in ("w_i", "e_i"):
        core = np.rot90(core, axes=(1, 0))  # clockwise
        core = np.fliplr(core)
        core = np.rot90(core, axes=(0, 1))  # counter-clockwise


    rots = determine_cw_rotations(edge_side, target_dir)
    core = np.rot90(core, k=rots, axes=(1, 0))  # clockwise

    return core

@lru_cache()
def determine_cw_rotations(current: str, target: str) -> int:
    dirs_list = ['n','e', 's', 'w']
    i_dirs_list = ['n_i','e_i', 's_i', 'w_i']

    if current.endswith('i'):
        start = i_dirs_list.index(current) + 1
    else:
        start = dirs_list.index(current) + 1
    
    end = dirs_list.index(target) + 1

    # [1,2,3,4]
    # 1, 2 == 1
    # 2, 1 == -1 
    return end - start

def normalize_inverted(core_full, match_side: str):
    if match_side in ('n_i', 's_i'):
        return np.fliplr(core_full)
    
    elif match_side in ('w_i', 'e_i'):
        core_full = np.rot90(core_full) # counter-clockwise
        core_full = np.fliplr(core_full)
        return np.rot90(core_full, axes=(1,0))

def get_edge(core, side: str) -> Tuple[str]:
    if side == 'n':
        return tuple([x[9] for x in core])
    elif side == 'e':
        return tuple(core[9])
    elif side == 'w':
        return tuple(core[0])
    elif side == 's':
        return tuple([x[0] for x in core])

    return tuple()

def print_square(square):
    square_rows = []
    for i in range(len(square)):
        rows = [x[i] for x in square]
        square_rows.append(np.concatenate(rows, axis=1))

    # reverse so printing is done correctly
    square_rows.reverse() 

    for row in square_rows:
        print()
        for line in row:
            print(''.join(line))

def main(rows: List[str]) -> int:
    images = parse_input(rows)

    matches: Dict[str, Set[str]] = collections.defaultdict(set)
    match_sides: Dict[Tuple[str, str], List[Tuple[str, str]]] = collections.defaultdict(
        list
    )
    for l_tile_number, l_image_data in images.items():
        for l_key, l_data in l_image_data.items():
            for r_tile_number, r_image_data in images.items():
                for r_key, r_data in r_image_data.items():
                    if (
                        r_tile_number == l_tile_number
                        or l_key.startswith("core")
                        or r_key.startswith("core")
                    ):
                        continue
                    if l_data == r_data:
                        matches[l_tile_number].add(r_tile_number)
                        matches[r_tile_number].add(l_tile_number)

                        match_sides[tuple([l_tile_number, r_tile_number])].append(
                            tuple([l_key, r_key])
                        )
                        match_sides[tuple([r_tile_number, l_tile_number])].append(
                            tuple([r_key, l_key])
                        )

    # Matches is dict of tile numbers to a set of Tuple pairs which match
    #  an edge of the titular tile number to another tile.

    sq: int = int(math.sqrt(len(matches.keys())))
    corners: Set[str] = {k for k, v in matches.items() if len(v) == 2}
    corner_coords = ((0, sq - 1), (sq - 1, 0), (0, 0), (sq - 1, sq - 1))

    square: List[List[str]] = np.full((sq, sq), None)
    available_tiles: Set[str] = {x for x in images.keys()}

    x, y = 0, 0
    prev_tile = ""
    while available_tiles:
        # if we're starting off, just select a corner
        if (x, y) == (0, 0):
            square[x][y] = list(corners)[0]

        # if y > 0, we can use the last available tile of the neighbor on the previous row
        elif y > 0 and square[x][y - 1]:
            square[x][y] = (matches[square[x][y - 1]] & available_tiles).pop()

        # In a corner, must select tiles with only two matches
        elif (x, y) in corner_coords:
            for tile in matches[prev_tile] & corners & available_tiles:
                if len(matches[tile]) == 2:
                    square[x][y] = tile
                    break

        # We are on the edge, must select tiles with three matches
        elif x == 0 or x == sq - 1 or y == 0 or y == sq - 1:
            for tile in matches[prev_tile] & available_tiles:
                if len(matches[tile]) == 3:
                    square[x][y] = tile
                    break

        # Otherwise we are in the interior of the square
        else:
            for tile in set(matches[prev_tile]) & available_tiles:
                if len(matches[tile]) == 4:
                    square[x][y] = tile
                    break

        # Wrap up for next iteration
        prev_tile = square[x][y]
        available_tiles.remove(square[x][y])
        if x == sq - 1:
            x, y = 0, y + 1
        else:
            x, y = x + 1, y

    # Have complete square, now need to rotate tiles to ensure row alignment
    square_filled = np.full((sq, sq), None)

    left_core = images[square[0][0]]['core_full']
    right_core = images[square[1][0]]['core_full']
    # matching_edges: List[Tuple[str, str]] = match_sides[(square[0][0], square[1][0])]
    # left_edge_dir, right_edge_dir = [x for x in filter(lambda x: not x[0].endswith("i"), matching_edges)][0]
    
    # (0,0)
    square_filled[0][0] = left_core

    # (1,0)
    square_filled[1][0] = right_core

    for left, right in itertools.permutations(['n', 'e', 's', 'w', 'n_i', 'e_i', 'w_i', 's_i'], r=2):
        left_tile = transform_core(images[square[0][0]]['core_full'], left, 'e')
        right_tile = transform_core(images[square[1][0]]['core_full'], right, 'w')
        
        if get_edge(left_tile, 'e') == get_edge(right_tile, 'w'):
            square_filled[0][0] = left_tile
            square_filled[1][0] = right_tile
            break

    print(square[0][0])
    print(square_filled[0][0])
    print()
    print(square[1][0])
    print(square_filled[1][0])

    for y in range(0, len(square)):
        for x in range(0, len(square)):
        
            # If already placed, ignore, (0,0) and (1,0)
            if square_filled[x][y] is not None:
                continue

            core_full = images[square[x][y]]["core_full"]

            # If we are on the first row then we want to match on the 'w' side of the 
            #  square
            if y == 0:
                prev_tile = images[square[x-1][y]]['core_full']
                for dir in ('n', 'e', 's', 'w', 'n_i', 'e_i', 'w_i', 's_i'): 
                    tile = transform_core(images[square[x][y]]['core_full'], dir, 'w')
                    if get_edge(tile, 'w') == get_edge()
                
                
            # If we are on any row above the first one we want to match on the 's' side of the 
            # square to match the prior row
            elif y > 0:
                # square_rows = np.concatenate([x[0] for x in square_filled[0]], axis=0)
                
                print_rows = collections.defaultdict(str)
                for i in range(len(square_filled)): # for each tile in row
                    for j in range(len(square_filled[i][0])): # for each row in tile:
                        print_rows[j] += ''.join(x[j] for x in square_filled[i][0])
                        

                for k,v in print_rows.items():
                    print(v)
                import pdb; pdb.set_trace()
                sys.exit(1)

                matching_edges: List[Tuple[str, str]] = [x for x in filter(lambda x: not x[0].endswith("i"), match_sides[(square[x][y], square[x][y-1])])]
                left_edge_dir, _ = matching_edges[0]
                if left_edge_dir.endswith('i'):
                    core_full = normalize_inverted(core_full, left_edge_dir)
                rots = determine_rotations(left_edge_dir, "s")
                core_full = np.rot90(core_full, k=rots)

            square_filled[x][y] = core_full
            
            print(x, y, square[x][y])
            for i in range(len(core_full)):
                row = []
                for j in range(len(core_full)):
                    row.append(core_full[i][j])
                
                print(''.join(row))
            print()


            # if x == len(square) - 1:
            #     # left, _ = match_sides[(square[x][y], square[x - 1][y])][0]
            #     # core_full = transform_core(core_full, left)
            #     left_edge = [x[0] for x in core_full]
            #     right_edge = [y[9] for y in images[square[x-1][y]]["core_full"]]
            #     print(square[x][y], square[x-1][y])
            # else:
            #     left_edge = [x[9] for x in core_full]
            #     right_edge = [y[0] for y in images[square[x+1][y]]["core_full"]]
            #     print(square[x][y], square[x+1][y])
                
            # Need to make sure match side faces 'right', default is 'e'
            # left, _ = match_sides[(square[x][y], square[x + 1][y])][0]

            # i = 0
            # while left_edge != right_edge:
            #     if i == 4:
            #         core_full = np.fliplr(core_full)
            #     else:   
            #         core_full = np.rot90(core_full)
                
            #     if x == len(square) - 1:
            #         left_edge = [x[0] for x in core_full]
            #     else:
            #         left_edge = [x[9] for x in core_full]

            #     print(left_edge, right_edge)
            #     i+=1

                    


            # import pdb; pdb.set_trace()
            # square_filled[x][y] = core_full

    square_rows = []
    for i in range(len(square_filled)):
        rows = [x[i] for x in square_filled]
        square_rows.append(np.concatenate(rows, axis=1))

    import pdb; pdb.set_trace()

    # reverse so printing is done correctly
    square_rows.reverse() 
    for row in square_rows:
        print()
        for line in row:
            print(''.join(line))
    
    # Finally will need to try all rotations and flips to search for sea monster

    not_sea_monster: int = 0

    return not_sea_monster


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    # test()
    print("Main result: ", main(rows))
