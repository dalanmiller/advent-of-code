from itertools import combinations

test_case = """35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows, 5)
    assert result == 127
    print(result)
    print("Tests complete.")

def parse_line(row):
    return int(row.strip())

def main(rows, preamble_length):

    numbers = [parse_line(row) for row in rows]

    sums = []
    broken = None
    for i, number in enumerate(numbers):
        
        if i <= preamble_length:
            continue

        sums = {sum(x) for x in combinations(numbers[i - preamble_length:i], 2)}

        if number not in sums: 
            print("Found break: ", i, number)
            return number    

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows, 25))
