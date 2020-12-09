import itertools

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
    result = main(rows, 127)
    print(result)
    assert result == 62
    print("Tests complete.")

def parse_line(row):
    return int(row.strip())


def main(rows, contiguous_list_sum):

    numbers = [parse_line(row) for row in rows]

    for i, _ in enumerate(numbers):
        current_index = i
        current_sum = numbers[i]
        current_numbers_list = [numbers[i]]
        while current_sum <= contiguous_list_sum: 
            print(i, current_index, current_sum)
            if current_sum == contiguous_list_sum:
                current_numbers_list.sort()
                return current_numbers_list[0] + current_numbers_list[-1]

            current_index += 1
            current_numbers_list.append(numbers[current_index])
            current_sum += numbers[current_index]
    

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows, 248131121))
