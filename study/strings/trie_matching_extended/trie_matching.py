# python3
import sys

NA = -1

def build_trie(patterns):
    nodes = 0
    tree = {0: {}}
    for pattern in patterns:
        # Root
        curr_node = tree[0]
        for symbol in pattern:
            if symbol in curr_node:
                pointer = curr_node[symbol]
                curr_node = tree.get(pointer, {})
            else:
                nodes += 1
                tree[nodes] = {}
                curr_node[symbol] = nodes
                curr_node = tree[nodes]

    return tree

def prefix_match(text, tree, pos):
    result = []

    text_i = 0
    node = tree[0]
    node_i = 0
    symbol = text[text_i]
    while True:
        if symbol in node and "$" in tree[node[symbol]]:
            result.append(pos)
            break
        elif symbol in node:
            text_i += 1
            if text_i >= len(text):
                break
            node = tree[node[symbol]]
            symbol = text[text_i]
        else:
            break

    return result


def solve(text, n, patterns):
    tree = build_trie(patterns)

    i = 0
    result = []
    while text:
        r = prefix_match(text, tree, pos=i)
        result.extend(r)
        text = text[1:]
        i += 1

    # write your code here
    return result

text = sys.stdin.readline().strip()
n = int(sys.stdin.readline().strip())

patterns = []
for i in range(n):
    patterns += [sys.stdin.readline().strip() + "$"]

ans = solve(text, n, patterns)

print(' '.join(map(str, ans)) + '\n')
