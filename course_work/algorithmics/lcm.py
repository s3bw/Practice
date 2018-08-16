# Uses python3
import sys

def eul(a, b):
    if b == 0:
        return a
    return eul(b, a%b)

def lcm_ad(a, b):
    compare = [a, b]
    _a, _b = max(compare), min(compare)
    low_com_de = eul(_a, _b)
    return int(_a/low_com_de)*_b

def lcm_naive(a, b):
    for l in range(1, a*b + 1):
        if l % a == 0 and l % b == 0:
            return l

    return a*b

if __name__ == '__main__':
    input = sys.stdin.read()
    a, b = map(int, input.split())
    print(lcm_ad(a, b))
