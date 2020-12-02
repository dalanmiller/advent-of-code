from itertools import combinations
from math import prod
from sys import exit

if __name__ == "__main__":

    with open("input") as o:
        # Set comprehension just in case there are doubled numbers
        #  didn't actually check if this was the case
        numbers = {int(x) for x in o.readlines() if x not in ("", "\n")}

    # Had to carefully read that `combinations` does not output 
    #  repeats so combinations('abc', 2) would output ['ab', 'ac', 'bc']
    combos = combinations(numbers, 3)

    for c in combos:
        if sum(c) == 2020:
            print(c, prod(c))
            exit(0)

