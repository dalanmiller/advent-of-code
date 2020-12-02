from collections import Counter
from sys import exit

if __name__ == "__main__":    
    valid = 0
    with open('input', 'r') as o:
        for line in o.readlines():
            char_range, char, password = line.split(" ")

            char = char[0] # remove :
            password = password.strip() # remove newline
            
            min_chars, max_chars = [int(x) for x in char_range.split("-")]

            c = Counter(password)

            if c[char] >= min_chars and c[char] <= max_chars:
                valid += 1
    
    print("Valid passwords: ", valid)

