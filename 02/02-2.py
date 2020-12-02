from collections import Counter
from sys import exit

if __name__ == "__main__":
    valid = 0
    with open("input", "r") as o:
        for line in o.readlines():
            char_range, char, password = line.split(" ")

            char = char[:-1]  # remove :
            password = password.strip()  # remove newline

            first_position, second_position = [int(x) for x in char_range.split("-")]

            first_present = password[first_position - 1] == char
            second_present = password[second_position - 1] == char

            if not all([first_present, second_present]) and (
                first_present or second_present
            ):
                valid += 1

    print("Valid passwords: ", valid)
