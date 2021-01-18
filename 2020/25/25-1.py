import re
import functools
import itertools
import math
import collections
from dataclasses import dataclass
from typing import List, Optional, Tuple, Literal, Dict, Callable, Set

test_case = """5764801
17807724"""


def test():
    rows = test_case.split("\n")
    # card = int(rows[0])
    # door = int(rows[1])

    v = 1
    for _ in range(8):
        v = transform(v)    
    assert v == 5764801

    v = 1
    for _ in range(11):
        v = transform(v)    
    assert v == 17807724

    v = 1
    for _ in range(8):
        v = transform(v, subject_num = 17807724)
    assert v == 14897079

    v = 1
    for _ in range(11):
        v = transform(v, subject_num = 5764801)
    assert v == 14897079

@functools.lru_cache(maxsize=1024)
def transform(v: int, subject_num: int = 7) -> int:
    v *= subject_num
    v %= 20201227
    return v

def create_encryption_key(pk: int, loops: int) -> int:
    v: int = 1
    for _ in range(loops):
        v = transform(v, subject_num = pk)
    return v

def find_loop(pk: int):
    i, v = 0, 1
    while v != pk:
        v = transform(v)
        i+=1
    return i

def main(card_pk, door_pk) -> int:
    
    card_loop = find_loop(card_pk)
    door_loop = find_loop(door_pk)

    print("card loop: ", card_loop, "door loop: ", door_loop)

    card_ek = create_encryption_key(card_pk, door_loop)
    door_ek = create_encryption_key(door_pk, card_loop)

    if card_ek != door_ek:
        print("ded", card_ek, door_ek)
        import sys
        sys.exit(1)
    
    print(transform.cache_info())
    return card_ek
    
    
if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x.strip() for x in o.readlines()]

    door_pk: int = int(rows[0])
    card_pk: int = int(rows[1])
    
    test()
    print("Main result: ", main(card_pk, door_pk))
    # 2408539
    # 16939008
