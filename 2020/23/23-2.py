import dataclasses
from typing import List, Optional, Dict, Tuple

test_case = """389125467"""

@dataclasses.dataclass
class N:
    __slots__ = ['value', 'prev', 'next']
    value: int
    prev: Optional["N"]
    next: Optional["N"]

def test():
    rows: List[str] = test_case.split("\n")

    matches = main(rows[0])
    print(matches)
    assert matches == 149245887792

def parse_cups(row: str):
    return [int(x) for x in row]

def main(row: str) -> int: 
    cups = parse_cups(row)
    counter = max(cups) + 1
    val_node_map: Dict[int, N] = {}

    # starter cup
    prev_cup = int(cups.pop(0))
    # convert into N class    
    prev_cup = N(prev_cup, None, None)
    # add to quick lookup
    val_node_map[prev_cup.value] = prev_cup

    # remember this is head
    head = prev_cup

    for cup in cups + [x for x in range(counter, 1000001)]:
        # create new N, prev is previous N
        n = N(cup, prev_cup, None)
        # add to lookup
        val_node_map[cup] = n
        # connect prev N to point to this new one
        prev_cup.next = n
        # previous cup is now this one
        prev_cup = n

    # wrap up by connecting last item to head and vice versa
    prev_cup.next = head 
    head.prev = prev_cup
    
    # 🤷🏻‍♂️
    del cups, counter

    # set current cup to head
    current_cup = head
    # precompute max
    max_num: int = max(val_node_map.keys())
    for _ in range(1, 10000001):
        
        # left most pick up cup and right most pick up cup
        pick_up_cups_left, pick_up_cups_right = current_cup.next, current_cup.next.next.next
        # restitch slice ends
        pick_up_cups_left.prev.next = pick_up_cups_right.next
        pick_up_cups_right.next.prev = pick_up_cups_left.prev
        # put values in tuple to check in later loop
        pick_up_cups_values: Tuple[int] = tuple([pick_up_cups_left.value, pick_up_cups_left.next.value, pick_up_cups_right.value])

        # init i to current cup - 1
        i: int = current_cup.value - 1
        destination: int = -1
        while destination == -1:
            if i in val_node_map and i not in pick_up_cups_values:
                destination = val_node_map[i].value
            else:
                i -= 1

            if i < 1:
                i = max_num

        # need to insert pick up cups to right of 'left' N
        left: N = val_node_map[destination]
        # right N of where new cups will be inserted
        right = left.next

        # connect left to leftmost pick up cup
        left.next = pick_up_cups_left
        # ... and vice versa
        pick_up_cups_left.prev = left

        # connect right to rightmost pick up cup
        right.prev = pick_up_cups_right
        # ... and vice versa
        pick_up_cups_right.next = right
        
        # new current cup move over one
        current_cup = current_cup.next
        
    # find N of value 1
    one = val_node_map[1]

    # grab adjacent two values
    adj, next_adj = one.next.value, one.next.next.value
        
    return adj * next_adj

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    test()
    print("Main result: ", main(rows[0]))
