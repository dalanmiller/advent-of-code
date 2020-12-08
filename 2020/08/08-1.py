import re


test_case = """nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6"""

def test():
    rows = [x for x in test_case.split("\n")]
    result = main(rows)
    assert result == 5
    print(result)
    print("Tests complete.")

line_pattern = re.compile(r"(\w\w\w) ([\+|\-]\d*)")

def parse_line(row):
    operation, number = re.match(line_pattern, row).groups()
    return (operation, int(number))


def main(rows):

    instructions = [parse_line(row) for row in rows]
    for row in rows:
        parse_line(row)

    accumulator = 0
    line = 0
    visited = set()
    while line not in visited:
        print(f"LINE {line}: {instructions[line]}")

        op, n = instructions[line]

        visited.add(line)
    
        if op == 'acc':
            accumulator += n
            line+=1
        elif op == 'jmp':
            line+=n 
        elif op == 'nop':
            line+=1        

    print("ACCU! :", accumulator)
    return accumulator
    
    

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
