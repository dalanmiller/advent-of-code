
def main(rows):
    for row in rows: 
        seats = [x for x in range(128)]
        columns = [x for x in range(8)]
        for char in row.strip():
            if char == "F": # lower half
                seats = seats[:len(seats)//2]
            elif char == "B": # upper half
                seats = seats[len(seats)//2:]

            if char == "L": # lower half
                columns = columns[:len(columns)//2]
            elif char == "R": # upper half
                columns = columns[len(columns)//2:]

        yield (seats[0], columns[0])


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    # Create empty plane || 2D grid
    airplane_seats = [[False for x in range(8)] for x in range(128)]
    
    # Mark seats as taken
    for taken_seat in main(rows):
        airplane_seats[taken_seat[0]][taken_seat[1]] = True

    # Create output to visualize
    with open("output", "w") as o:
        for i, column in enumerate(airplane_seats):
            o.write(f"row:{i:3} | ")
            for j, row in enumerate(column):
                o.write(f"{str(row).ljust(5)}/{((i * 8) + j):4} |")
                if j == len(column) - 1:
                    o.write("\n")
    
    


    

    