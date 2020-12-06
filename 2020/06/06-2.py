from functools import reduce

test_case = """abc

a
b
c

ab
ac

a
a
a
a

b
"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    assert result == 6
    print(result)
    print("Tests complete.")

def main(rows):
    total = []
    group = []
    for row in rows:
        if row in ("\n", "") and len(group):
            total.append(group)
            group = []  # reset group
        else:
            group.append({x for x in row})

    if len(group):  # Weirdly empty row on input file, not appending last group
        total.append(group)

    # Reduce each group via intersection, then length of that set
    intersected_groups = [len(reduce(lambda a, b: a.intersection(b), x)) for x in total]
    # Sum all lengths
    return sum(intersected_groups)

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
