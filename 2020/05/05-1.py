test_case = r"""
FBFBBFFRLR\n
BFFFBBFRRR\n
FFFBBBFRRR\n
BBFFBBFRLL\n
"""

def test(rows):
    results = main(rows)
    results_answers = zip(results, (357, 567, 119, 820))
    for result, answer in results_answers:
        print("Checking: ", result, answer, "✅")
        assert(result == answer)

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

        yield (seats[0] * 8) + columns[0]


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    test([x for x in test_case.split("\n") if x not in ("\n", "")])

    print("Max ID: ", max(main(rows)))
    

    