from math import prod

if __name__ == "__main__":
    rows = None
    with open('input') as o:
        rows = [x for x in o.readlines()]

    trees = 0
    x,y = 0, 0
        
    while y < len(rows):
        x = (x + 3) % 31 
        y += 1
        if y < len(rows) and rows[y][x] == "#":
            trees+=1
        
    print("Hit trees: ", trees)

    