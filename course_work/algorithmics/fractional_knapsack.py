# Uses python3
import sys

def get_optimal_value(capacity, weights, values):
    options = zip(weights, values)

    sorted_options = sorted(
        options, 
        key=lambda x: x[1]/x[0], 
        reverse=True
    )

    value = 0
    for option in sorted_options:
        if capacity != 0 and option[0] <= capacity:
            value += option[1]
            capacity -= option[0]

        elif capacity != 0 and option[0] > capacity:
            value += (option[1]/option[0])*capacity
            capacity = 0

        elif capacity == 0:
            return value
    return value


if __name__ == "__main__":
    data = list(map(int, sys.stdin.read().split()))
    n, capacity = data[0:2]
    values = data[2:(2 * n + 2):2]
    weights = data[3:(2 * n + 2):2]
    opt_value = get_optimal_value(capacity, weights, values)
    print("{:.10f}".format(opt_value))
