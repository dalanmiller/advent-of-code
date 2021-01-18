import re
import functools
import itertools

test_case = """0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
6: 4 4 4 4 4

ababbb
bababa
abbbab
aaabbb
aaaabbb"""

def test():
    rows = test_case.split("\n")
    
    matches = main(rows)
    print(matches)
    assert matches == 2

text_rule = re.compile(r"(\d*)\: \"([ab])\"")
condition_rule = re.compile(r"(\d*)\: ([\d\s\|]*)" )
def parse_rules(rules_lines):
    rules = {}
    for rule_text in rules_lines:
        if "\"" in rule_text:
            groups = re.match(text_rule, rule_text).groups()
            rules[groups[0]] = tuple(groups[1])
        else:
            rule_number, rule = re.match(condition_rule, rule_text).groups()
            rule_list = rule.split(" ")

            if '|' in rule_list:
                i = rule_list.index("|")
                left = tuple(x for x in rule_list[:i])
                right = tuple(x for x in rule_list[i+1:])
                rule_list = [left, right]
            else:
                rule_list = [tuple(x for x in rule_list)]

            rules[rule_number] = rule_list

    return rules

def parse_messages(messages):
    return set(x for x in messages)

def main(rows) -> int:
    i = rows.index('')
    rules_lines = rows[:i]
    messages_lines = rows[i+1:]

    rules = parse_rules(rules_lines)
    messages = parse_messages(messages_lines)

    @functools.lru_cache(maxsize=len(rules))
    def resolve(rule_number):
        if type(rules[rule_number][0]) == str and not rules[rule_number][0].isdigit():
            return rules[rule_number][0]

        messages = set()
        for rule in rules[rule_number]:
            sub_rules = (resolve(sub) for sub in rule)
            all_rules = {x for x in itertools.product(*sub_rules)}
            joined_rules = {''.join(x) for x in all_rules}

            messages |= joined_rules

        return messages

    possible_messages = resolve('0')

    return len(possible_messages & messages)

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))


