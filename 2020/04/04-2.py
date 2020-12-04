import re

def tests(fields):
    assert fields["byr"]("2002") == True
    assert fields["byr"]("2003") == False

    assert fields["hgt"]("60in") == True
    assert fields["hgt"]("190cm") == True
    assert fields["hgt"]("190in") == False
    assert fields["hgt"]("190") == False
    assert fields["hgt"]("1900") == False
    assert fields["hgt"]("300cm") == False
    assert fields["hgt"]("0cm") == False
    assert fields["hgt"]("cm") == False

    assert fields["hcl"]("#123abc") == True
    assert fields["hcl"]("#123abz") == False
    assert fields["hcl"]("123abc") == False

    assert fields["ecl"]("brn") == True
    assert fields["ecl"]("wat") == False

    assert fields["pid"]("000000001") == True
    assert fields["pid"]("0123456789") == False
    assert fields["pid"]("0") == False
    print("Tests complete.")


if __name__ == "__main__":
    rows = None
    with open("input") as o:
        rows = [x for x in o.readlines()]

    # keys are passport field names, values are rules for passport field value
    fields = {
        "byr": lambda x: len(x) == 4 and 1920 <= int(x) <= 2002,
        "iyr": lambda x: len(x) == 4 and 2010 <= int(x) <= 2020,
        "eyr": lambda x: len(x) == 4 and 2020 <= int(x) <= 2030,
        "hgt": lambda x: re.match(r"\d*[cm$|in$]", x) 
            and len(x[:-2]) and 150 <= int(x[:-2]) <= 193
            if x.endswith("cm") 
            else len(x[:-2]) and x != "" and 59 <= int(x[:-2]) <= 76,
        "hcl": lambda x: bool(re.match(r"^#[0-9a-fA-F]{6}", x)),
        "ecl": lambda x: x in ("amb", "blu", "brn", "gry", "grn", "hzl", "oth"),
        "pid": lambda x: bool(re.match(r"^\d{9}$", x)),
    }

    # Ensure basic assertions pass
    tests(fields)

    joined_fields = "|".join(fields.keys())

    # group 1 field key/name
    # group 2 field value/string
    pattern = re.compile(rf"^({joined_fields})\:([\#\w\d]*)")

    # init empty passport, all values defaulted to False
    current_passport = {f: False for f in fields.keys()}
    valid_passports = 0
    for row in rows:
        if row == "\n" or row == "":  # empty row signals end of field grouping
            if all(
                x for x in current_passport.values()
            ):  # proceed if all fields of passport are found and validated
                valid_passports += 1
            current_passport = {f: False for f in fields}  # reset passport

        for item in row.strip().split(" "):  # for individual fields:values in row

            if match := re.match(pattern, item.strip()): # walrus operator! So good! 

                field_name = match.groups()[0]
                field_string = match.groups()[1]

                # 1. check if field name we care about
                # 2. check if field string passes validation
                if field_name in fields and fields[field_name](field_string):
                    current_passport[field_name] = True

    print("Valid passports: ", valid_passports)
