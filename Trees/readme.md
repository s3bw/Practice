<h1 align="center">
    Trees
</h1>

## Binary Search Tree

Sometimes called ordered or sorted binary tree, are a
particular type of container data structures that store
"items" (such as numbers, names etc.) in memory. They allow
fast lookup, addition and removal of items, and can be used
to implement either dynamic sets of items, or lookup tables
that allow finding an item by its key.

Items are stored in a binary tree. The item is compared to
the root if the item value is larger than the root it is
attached to the right branch, if this is larger than the
inserted item, this comparison happens recursively until we
find a leaf position that is not occupied.

```
20
└── 16
    ├── 13
    │   ├── 6
    │   │   ├── 1
    │   │   │   └── 2
    │   │   │       └── 3
    │   │   │           └── 5
    │   │   │               └── 4
    │   │   └── 7
    │   │       └── 8
    │   │           └── 12
    │   │               └── 10
    │   │                   ├── 9
    │   │                   └── 11
    │   └── 14
    │       └── 15
    └── 19
        └── 17
            └── 18

Depth:  9
Tree is NOT balanced!
```

## AVL Tree

The AVL tree is a self-balancing binary search tree. The
heights of the two child subtrees of any node differ by at
most one; if at any time they differ by more than one,
rebalancing is done to restore this property. Lookup,
insertion, and deletion all take O(log n) time in both the
average and worst cases, where n is the number of nodes in
the tree prior to the operation.

Rebalance is done by one or more "tree rotations". AVL trees
are faster than red-black trees.

```
11
├── 5
│   ├── 2
│   │   ├── 1
│   │   └── 3
│   │       └── 4
│   └── 9
│       ├── 7
│       │   ├── 6
│       │   └── 8
│       └── 10
└── 17
    ├── 14
    │   ├── 12
    │   │   └── 13
    │   └── 16
    │       └── 15
    └── 19
        ├── 18
        └── 20

Depth:  5
Tree is balanced!
```

## Red-Black Tree


- b-Trees
- r-Trees
- Red-Black Trees
- DSW - tree

http://btechsmartclass.com/DS/U5_T5.html

https://rosettacode.org/wiki/AVL_tree#Python

