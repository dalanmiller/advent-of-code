import re
import functools
import itertools

test_case = """42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba"""

def test():
    rows = test_case.split("\n")
    
    matches = main(rows)
    print(matches)
    assert matches == 12

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
    cache = {}

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

        cache[rule_number] = messages
        return messages

    @functools.lru_cache(maxsize=len(rules))
    def expand_rule(rule, msg):
        if msg in cache[rule]:
            return True

        # If the message isn't in a pre-derived list of existing rules, try to dynamically
        # expand it using the rule 0 => 8 11 by finding msg = ab such that 8 derives a and
        # 11 derives b.
        if rule == "0":
            step = min(map(len, rules["8"]))
            for i in range(step, len(msg) - 1, step):
                a, b = msg[:i], msg[i:]
                if expand_rule("8", a) & expand_rule("11", b):
                    return True

        # 8 => 42 | 42 8  - translates to 42+ in regex terms
        if rule == "8":
            while True:
                for w in cache["42"]:
                    if msg.startswith(w):
                        msg = msg[len(w) :]
                        if len(msg) == 0:
                            return True
                        break
                else:
                    break

        # 11 => 42 31 | 42 11 31  - translate to 42{n}31{n} in regex terms
        if rule == "11":
            while True:
                for w, v in itertools.product(cache["42"], cache["31"]):
                    if msg.startswith(w) and msg.endswith(v):
                        msg = msg[len(w) : -len(v)]
                        if len(msg) == 0:
                            return True
                        break
                else:
                    break

        return False

    resolve('0')
    return sum(expand_rule("0", m) for m in messages)

if __name__ == "__main__":
    rows = None
    with open("input2") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))


