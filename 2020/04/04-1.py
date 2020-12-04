import re

if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
    joined_fields = "|".join(fields)

    # group 1 field key/name
    # group 2 field value/string
    pattern = re.compile(rf"^({joined_fields})\:[\#\w\d]*")

    current_passport = {f: False for f in fields}
    valid_passports = 0
    for row in rows:
        if row.startswith("\n"):
            if all([x for x in current_passport.values()]):
                valid_passports += 1
            current_passport = {f: False for f in fields}

        for item in row.strip().split(" "):
            if match := re.match(pattern, item.strip()):
                if match.groups()[0] in fields:
                    current_passport[match.string[:3]] = True

    print("Valid passports: ", valid_passports)
    