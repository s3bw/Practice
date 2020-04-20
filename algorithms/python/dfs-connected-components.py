"""
DFS to find largest number of connected components.
"""


def most_connected(grid):
    nrows = len(grid)
    ncols = len(grid[0])

    def travel(colour, i, j):
        # Mark as seen
        grid[i][j] = False
        result = 1

        if i > 0 and grid[i - 1][j] == colour:
            result += travel(colour, i - 1, j)

        if j > 0 and grid[i][j - 1] == colour:
            result += travel(colour, i, j - 1)

        if i < len(grid) - 1 and grid[i + 1][j] == colour:
            result += travel(colour, i + 1, j)

        if j < len(grid[0]) - 1 and grid[i][j + 1] == colour:
            result += travel(colour, i, j + 1)
        return result

    component_size = 0

    for i in range(nrows):
        for j in range(ncols):
            if grid[i][j]:
                temp = travel(grid[i][j], i, j)
                if temp > component_size:
                    component_size = temp

    return component_size


if __name__ == "__main__":
    tiles = [
        ["G", "B", "B", "B", "P"],
        ["P", "B", "G", "P", "G"],
        ["P", "B", "G", "G", "G"],
        ["B", "B", "P", "P", "G"],
    ]
    result = most_connected(tiles)
    print(result)
