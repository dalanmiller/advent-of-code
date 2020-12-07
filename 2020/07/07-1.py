import re
from functools import reduce
from functools import lru_cache

test_case = """light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags."""


container_pattern = re.compile(r"(\w+ \w+) (\w+)")
contained_pattern = re.compile(r"\d+ (\w+ \w+) bags?")


def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows, "shiny gold")
    assert result == 4
    print(result)
    print("Tests complete.")


def parse_line(row):
    container, containing = map(lambda x: x.strip(), row.split("contain"))
    container_bag, _ = re.match(container_pattern, container).groups()

    if containing.strip().startswith("no other bags"):
        contained_bags = None
    else:
        contained_bags = re.findall(contained_pattern, containing)

    return (container_bag, contained_bags)


def main(rows, target_bag_color):
    rule_set = {}
    for row in rows:
        if row == "":
            continue
        container, contained_bags = parse_line(row)
        rule_set[container] = contained_bags

    @lru_cache(maxsize=256)
    def recurse_bags(bag_color, target_bag_color):
        if rule_set[bag_color] == None:
            return False
        elif target_bag_color in rule_set[bag_color]:
            return True
        else:
            results = [
                recurse_bags(bag, target_bag_color) for bag in rule_set[bag_color]
            ]
            if len(results) > 1:
                return any(results)
            else:
                return results[0]

    results = []
    for bags in rule_set.keys():
        results.append(recurse_bags(bags, target_bag_color))

    return reduce(lambda x, y: x + y, results, 0)


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows, "shiny gold"))
