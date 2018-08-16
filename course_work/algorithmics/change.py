# Uses python3
import sys

# Denominations = [10, 5, 1]

def get_change(m):
    coins = 0
    coins += int(m/10)
    m = m % 10
    coins += int(m/5)
    m = m % 5
    coins += int(m/1)

    return coins

if __name__ == '__main__':
    m = int(sys.stdin.read())
    print(get_change(m))
