# Uses python3
import sys

def fib_list(n):
    i = [0, 1]
    for m in range(n):
        new_fib = (i[m + 1] + i[m]) % 10
        i.append(new_fib)
    return(i[n])

def get_fibonacci_last_digit_naive(n):
    if n <= 1:
        return n

    previous = 0
    current  = 1

    for _ in range(n - 1):
        previous, current = current, previous + current

    return current % 10

if __name__ == '__main__':
    input = sys.stdin.read()
    n = int(input)
    print(fib_list(n))
