# taken from https://github.com/juanplopes/advent-of-code-2022/blob/main/day18.py

import sys, collections

input = open("input").read()
cubes = set(tuple(map(int, x.split(","))) for x in input.splitlines())
mins, maxs = [min(x) - 1 for x in zip(*cubes)], [max(x) + 1 for x in zip(*cubes)]


def neighbors(x, y, z):
    yield from ((x + 1, y, z), (x, y + 1, z), (x, y, z + 1))
    yield from ((x - 1, y, z), (x, y - 1, z), (x, y, z - 1))


total1 = sum(n not in cubes for c in cubes for n in neighbors(*c))
total2 = 0

# Create queue of cubes which starts with the previously found min
# As well, consider all known cubes as 'visited' but some cubes could have empty space
# and we don't want to visit twice.
Q, visited = collections.deque([tuple(mins)]), set(cubes)

# while we have gas in the tank
while len(Q):
    # for each neighboring cube of a popped cube from the deque
    for cube in neighbors(*Q.popleft()):

        # Check if that coordinate is within the bounds of the seen min/max
        # if it isn't, then we skip this neighboring cube.
        if not all(a <= x <= b for a, b, x in zip(mins, maxs, cube)):
            continue

        # We've arrived here, so we know that the cube is within bounds
        # add if the found cube is in the cubes set
        total2 += cube in cubes

        # If we've already visited, then continue
        if cube in visited:
            continue

        # Otherwise add to the visited set, and append as a future cube
        # . to visit
        visited.add(cube)
        Q.append(cube)
print(total1, total2)
