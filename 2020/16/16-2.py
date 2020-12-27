from typing import List, Tuple, Dict
import collections
import math
import re


test_case = """departure_class: 0-1 or 4-19
departure_row: 0-5 or 8-19
departure_seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9"""


def test():
    rows = test_case.split("\n")
    print(rows)
    result = main(rows)
    print(result)
    assert result == 1716


rule_pattern = re.compile(r"([\w\s]+)\: (\d*)-(\d*) or (\d*)-(\d*)")


def parse_rules(rules: List[str]) -> Dict[str, Tuple[Tuple[int, int], Tuple[int, int]]]:
    parsed_rules: Dict[str, Tuple[Tuple[int, int], Tuple[int, int]]] = {}
    for rule in rules:
        type, a1, a2, b1, b2 = re.match(rule_pattern, rule).groups()
        parsed_rules[type] = ((int(a1), int(a2)), (int(b1), int(b2)))

    return parsed_rules


def parse_nearby_tickets(nearby_tickets: List[str]) -> List[List[int]]:
    tickets = nearby_tickets[2:]
    return [[int(y) for y in x.split(",")] for x in tickets]


def invalid_fields(ticket: List[str], rules):
    for field in ticket:
        if not any(
            r[0][0] <= field <= r[0][1] or r[1][0] <= field <= r[1][1] for r in rules
        ):
            yield field


def check_valid(field: int, rule: Tuple[Tuple[int, int], Tuple[int, int]]) -> bool:
    first_range, second_range = rule[0], rule[1]
    return (
        first_range[0] <= field <= first_range[1]
        or second_range[0] <= field <= second_range[1]
    )


def main(lines: List[str]) -> int:
    rules_section_end = lines.index("")
    my_ticket_end = lines.index("", rules_section_end + 1)

    rule_lines: List[str] = lines[:rules_section_end]
    my_ticket_lines: List[str] = lines[rules_section_end:my_ticket_end]

    my_ticket = [int(x) for x in my_ticket_lines[2].split(",")]

    nearby_tickets_lines = lines[my_ticket_end:]

    rules = parse_rules(rule_lines)
    tickets = parse_nearby_tickets(nearby_tickets_lines)

    valid_tickets: List[List[int]] = []
    for ticket in tickets:
        any_true: bool = True
        for field in ticket:
            if not any(check_valid(field, rule) for rule in rules.values()):
                any_true = False
                break

        if any_true:
            valid_tickets.append(ticket)

    rule_col_map: Dict[str, List[int]] = collections.defaultdict(list)
    for name, rule in rules.items():
        for i in range(len(valid_tickets[0])):
            value_column = (x[i] for x in valid_tickets)
            if all(check_valid(x, rule) for x in value_column):
                rule_col_map[name].append(i)

    identified_fields: Dict[str, int] = {}
    rule_col_map = {
        k: v for k, v in sorted(rule_col_map.items(), key=lambda x: len(x[1]))
    }
    while len(identified_fields) != len(my_ticket):
        for k, v in rule_col_map.items():
            if len(v) == 1:
                column_number = v.pop(0)
                identified_fields[k] = column_number
                for v in rule_col_map.values():
                    if column_number in v:
                        v.remove(column_number)

    return math.prod(
        [
            my_ticket[v]
            for k, v in identified_fields.items()
            if k.startswith("departure")
        ]
    )


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
    # 22073
    # 4195872914689
    # 920453493481
    # 1346570764607
