from collections import Counter

test_case = """16
10
15
5
1
11
7
19
6
12
4"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 7 * 5
    print("Tests complete.")

def parse_line(row):
    return int(row.strip())

def main(rows):
    adapters = [0] + [parse_line(x) for x in rows]
    adapters.sort()
    adapters.append(adapters[-1] + 3)
    
    count = Counter([x - adapters[i-1] for i, x in enumerate(adapters[1:], start=1)])

    return count[1] * count[3]

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
