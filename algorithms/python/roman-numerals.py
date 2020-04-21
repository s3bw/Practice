"""
Convert roman numerals into arabic numerals

Tests:
https://www.sporcle.com/games/g/romannumerals
"""
import sys

roman = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}


def compare(c1, c2):
    v1, v2 = roman[c1], roman[c2]
    if v1 < v2:
        return True
    return False


def convert_numerals(string):
    count = 0
    p = 0
    f = 1
    results = []
    while True:
        try:
            c = string[p]
            future = string[f]
        except IndexError:
            if p == len(string):
                break
            results.append(roman[c])
            break

        if compare(c, future):
            count = roman[future] - roman[c]
            results.append(count)
            p += 2
            f += 2

        else:
            results.append(roman[c])
            c = future
            p += 1
            f += 1

    return sum(results)


if __name__ == "__main__":
    result = convert_numerals("VI")
    print(f"Got: {result}, Expected: 6")

    result = convert_numerals("IV")
    print(f"Got: {result}, Expected: 4")

    result = convert_numerals("VII")
    print(f"Got: {result}, Expected: 7")

    result = convert_numerals("CCVII")
    print(f"Got: {result}, Expected: 207")

    string = sys.argv[1].upper()
    print(convert_numerals(string))
