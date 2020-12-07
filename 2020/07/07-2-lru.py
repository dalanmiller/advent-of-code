import re
from functools import lru_cache

test_case = """shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags."""

container_pattern = re.compile(r"(\w+ \w+) (\w+)")
contained_pattern = re.compile(r"(\d)+ (\w+ \w+) bags?")


def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows, "shiny gold")
    assert result == 126
    print(result)
    print("Tests complete.")


def parse_line(row):
    container, containing = map(lambda x: x.strip(), row.split("contain"))
    container_bag, _ = re.match(container_pattern, container).groups()

    if containing.strip().startswith("no other bags"):
        contained_bags = None
    else:
        contained_bags = {
            v: int(k) for k, v in re.findall(contained_pattern, containing)
        }

    return (container_bag, contained_bags)


def main(rows, target_bag_color):
    rule_set = {}
    for row in rows:
        if row == "":
            continue
        container, num_bags = parse_line(row)
        rule_set[container] = num_bags

    # Originally had this outside this method, but can't take
    #  advantage of lru_cache in this case given that dicts are not hashable
    #  having an inline function prevents needing a signature like the following
    #  recurse_bags(rule_set, bag_color)...
    @lru_cache(maxsize=256)
    def recurse_bags(bag_color):
        if rule_set[bag_color] == None:
            return 0
        else:
            results = sum(v for v in rule_set[bag_color].values()) 
            for color, num in rule_set[bag_color].items():

                # multiply sum of inner bags by number of bags observed
                results += num * recurse_bags(color)

            return results

    result = recurse_bags(target_bag_color)

    return result

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows, "shiny gold"))
