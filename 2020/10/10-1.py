from itertools import combinations

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
    adapters = [parse_line(x) for x in rows]
    adapters.sort()
    adapters.insert(0, 0)

    count = {1: 0, 2: 0, 3: 1}
    print(adapters)
    
    for i, current_adapter in enumerate(adapters, start=0):

        if i + 1 < len(adapters):
            next_adapter = adapters[i + 1]
        
            joltage_delta = next_adapter - current_adapter
            print(i, current_adapter, next_adapter, joltage_delta)
            count[joltage_delta] += 1

        
    print("Final count", count)

    return count[1] * count[3]

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
