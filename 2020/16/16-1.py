from typing import List, Tuple, Dict, Callable
import math
import sys
import re
from functools import reduce
import operator

test_case = """class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    print(result)
    assert result == 71
    
    print("Tests complete.")


rule_pattern = re.compile(r"([\w\s]+)\: (\d*)-(\d*) or (\d*)-(\d*)")
def parse_rules(rules: List[str]) -> Dict[str, Tuple[Tuple[int,int], Tuple[int,int]]]:
    parsed_rules = {}
    for rule in rules:
        type, a1, a2, b1, b2 = re.match(rule_pattern, rule).groups()
        parsed_rules[type] = ((int(a1), int(a2)), (int(b1),int(b2)))

    return parsed_rules

def parse_nearby_tickets(nearby_tickets: List[str]) -> List[List[int]]:
    tickets = nearby_tickets[2:]
    # print(tickets)
    return [[int(y) for y in x.split(",")] for x in tickets]

def main(rows: List[str]) -> int:
    
    lines = rows
    rules_section_end = lines.index('')
    my_ticket_end = lines.index('', rules_section_end + 1)

    rule_lines = lines[:rules_section_end]

    nearby_tickets_lines = lines[my_ticket_end:]

    rules = parse_rules(rule_lines)
    tickets = parse_nearby_tickets(nearby_tickets_lines)

    valid_tickets = []
    valid_field_count = {k:{i:0 for i in range(len(tickets[0]))} for k in rules.keys()}
    invalid_fields: List[int] = []
    for ticket in tickets:
        for i, val in enumerate(ticket): 
            valid = False
            for rule_name, rule in rules.items():
                first_range, second_range = rule[0], rule[1]
                if first_range[0] <= val <= first_range[1] or second_range[0] <= val <= second_range[1]:
                    valid = True
                    valid_field_count[rule_name][i] += 1
                
            if valid: 
                valid_tickets.append(ticket)
            else:
                invalid_fields.append(val)
    
    return sum(invalid_fields)

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
    # 1939887
