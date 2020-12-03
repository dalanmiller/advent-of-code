from math import prod

if __name__ == "__main__":
    rows = None
    with open('input') as o:
        rows = [x for x in o.readlines()]

    paths = [(1, 1), (3, 1), (5,1), (7,1), (1, 2)]
    path_outcomes = []

    for path in paths:
        x, y, trees = 0,0,0

        while y < len(rows):
            x = (x + path[0]) % 31 
            y += path[1]
            if y < len(rows) and rows[y][x] == "#":
                trees+=1
            
        path_outcomes.append(trees)
        
    print("Hit trees: ", path_outcomes)
    print("Product: ", prod(path_outcomes))

    