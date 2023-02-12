#Uses python3
import sys

# Return the trie built from patterns
# in the form of a dictionary of dictionaries,
# e.g. {0:{'A':1,'T':2},1:{'C':3}}
# where the key of the external dictionary is
# the node ID (integer), and the internal dictionary
# contains all the trie edges outgoing from the corresponding
# node, and the keys are the letters on those edges, and the
# values are the node IDs to which these edges lead.

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

if __name__ == '__main__':
    patterns = sys.stdin.read().split()[1:]
    tree = build_trie(patterns)
    for node in tree:
        for c in tree[node]:
            print("{}->{}:{}".format(node, tree[node][c], c))
