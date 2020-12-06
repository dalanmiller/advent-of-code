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
    assert(result == 11)
    print(result)
    print("Tests complete.")


def main(rows):
    total = []
    group = set()
    for row in rows:
        if row in ("\n", "") and len(group):
            total.append(group)
            group = set() # reset group
        else:
            group = group.union({x for x in row})

    return sum(len(x) for x in total)

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
