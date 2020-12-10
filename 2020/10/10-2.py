from functools import lru_cache

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
    assert result == 8
    print("Tests complete.")

def parse_line(row):
    return int(row.strip())

def construct_tree(adapters):
    tree = {}
    for adapter in adapters:
        possible_joltages = {adapter + 1, adapter + 2, adapter +3}
        tree[adapter] = possible_joltages.intersection(adapters)

    return tree

def generate_paths(adapter_tree):
    start = min(adapter_tree.keys())
    end = max(adapter_tree.keys())

    @lru_cache(maxsize=len(adapter_tree))
    def dive(node):
        if node == end:
            return 1
        else:
            return sum([dive(x) for x in adapter_tree[node]])

    return dive(start)
    
def main(rows):
    adapters = [0] + [parse_line(x) for x in rows]
    adapters.sort()
    adapters.append(adapters[-1] + 3)

    adapter_tree = construct_tree(adapters)
    paths = generate_paths(adapter_tree)

    return paths

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
