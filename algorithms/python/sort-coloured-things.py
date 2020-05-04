x = ["G", "B", "G", "R", "B", "G"]


def sort(array):
    d = {
        "B": 0,
        "G": 0,
        "R": 0,
    }

    for n in array:
        d[n] += 1

    i = 0
    while True:
        while d["R"] > 0:
            array[i] = "R"
            d["R"] -= 1
            i += 1

        while d["G"] > 0:
            array[i] = "G"
            d["G"] -= 1
            i += 1

        while d["B"] > 0:
            array[i] = "B"
            d["B"] -= 1
            i += 1

        break

    return array


print(sort(x))
