from os import listdir, mkdir
import re

if __name__ == "__main__":
    dirs = listdir(".") # root of dir

    days = [int(x) for x in dirs if re.match(r"\d\d", x)]

    next_day = str(max(days) + 1)
    mkdir(f"{next_day}")

    for x in ["-1", "-2"]:
        open(f"./{next_day}/{next_day}{x}.py", "w")

    open(f"./{next_day}/input", "w")

