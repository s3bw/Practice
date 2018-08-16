# Uses python3

def fib_list(n):
    i = [0, 1]
    for m in range(n):
        i.append(i[m + 1] + i[m])
    return(i[n])

def calc_fib(n):
    if (n <= 1):
        return n

    return calc_fib(n - 1) + calc_fib(n - 2)

n = int(input())
print(fib_list(n))
