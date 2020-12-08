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
    assert result == 8
    print(result)
    print("Tests complete.")

line_pattern = re.compile(r"(\w\w\w) ([\+|\-]\d*)")

def parse_line(row):
    operation, number = re.match(line_pattern, row).groups()
    return (operation, int(number))

def main(rows):
    instructions = [parse_line(row.strip()) for row in rows]

    jmps_and_nops = (i for i, x in enumerate(instructions) if x[0] in ('nop', 'jmp'))

    for location in jmps_and_nops:
        changed_instructions = instructions.copy()
        op, num = changed_instructions[location] 
          
        if op == 'nop' and num == 0:
            continue # changing to a jmp +0 won't do anything so we can skip 
        elif op == 'nop':
            changed_instructions[location] = ('jmp', num) 
        else:
            changed_instructions[location] = ('nop', num)

        accumulator = 0
        line = 0
        visited = set()
        while line not in visited and line < len(changed_instructions):

            op, n = changed_instructions[line]
            
            visited.add(line)
        
            if op == 'acc':
                accumulator += n
                line+=1
            elif op == 'jmp':
                line+=n 
            elif op == 'nop':
                line+=1
                        
        if line == len(instructions): # can jump the for loop early if we find answer
            break 
    
    return accumulator

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows))
