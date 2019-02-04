# Bitwise Operation

>Bitwise operations are everywhere. They are perfect for working
with bitfields (a practice that is ubiquitous in C and C++), such
as a 'flags' field in a data structure or function argument.
Basically, `|` combines flags, `^` flips flags, `&` checks if a
flag is set, and the `x &= ~FLAG` pattern clears a flag.

>Bitwise operations are ubiquitous in all thins low-level -
hardware drivers, network protocols, binary file formats - as well
as some higher-level fields like character encodings, cryptography
etc.

Taken from: https://softwareengineering.stackexchange.com/a/112487/292820

## Operands

### NOT

The bitwise Not, or complement, is a unary operation that performs
logical negation on each bit, forming the ones complement of the given
binary value. Bits that are 0 become 1, and those that are 1 become 0.
For example:

```
NOT 0111 (decimal 7)
  = 1000 (decimal 8)
```

### AND

A bitwise AND takes two equal-length binary representations and
performs the logical AND operation on each pair of the corresponding
bits, which is equivalent to multiplying them. Thus, if both bits in
the compared position are 1, the bit in the resulting binary
representation is 1 (1 x 1 = 1); otherwise, the result is 0
(1 x 0 = 0 and 0 x 0 = 0). For example:

```
    0101 (decimal 5)
AND 0011 (decimal 3)
  = 0001 (decimal 1)
```

### OR (also set-to-1)

A bitwise OR takes two bit patterns of equal length and performs the
logical inclusive OR operation on each pair of corresponding bits. The
result in each position is 0 if both bits are 0, while otherwise the
result is 1. For example

```
    0101 (decimal 5)
 OR 0011 (decimal 3)
  = 0111 (decimal 7)
```

### XOR (also toggle or flip)

**One or the Other, Not both**

A bitwise XOR takes two bit patterns of equal length and performs the
logical exclusive OR operation on each pair of corresponding bits. THe
result in each position is 1 if only the first bit is 1 or only the
second bit is 1, but will be 0 if both are 0 or both are 1. In this we
perform the comparison of two bits, being 1 if the two bits are
different, and 0 if they are the same. For example:

```
    0101 (decimal 5)
XOR 0011 (decimal 3)
  = 0110 (decimal 6)
```

## Truth Table

From the above operand we have 16 possible truth functions of two
binary variables.

```
p = [0, 0, 0, 0]
q = [0, 0, 0, 0]

# False
a = [0, 0, 0, 0]

# NOT (p OR q)
# (aka NOR)
a = [0, 0, 0, 1]

# (NOT p) AND q
# (aka Xq)
a = [0, 0, 1, 0]

# NOT p
a = [0, 0, 1, 1]

# p AND (NOT q)
a = [0, 1, 0, 0]

# NOT q
a = [0, 1, 0, 1]

# p XOR q
a = [0. 1. 1. 0]

# NOT (p AND q)
a = [0, 1, 1, 1]

# p AND q
a = [1, 0, 0, 0]

# NOT (p XOR q)
a = [1, 0, 0, 1]

# q
a = [1, 0, 1, 0]

# (NOT p) OR q
a = [1, 0, 1, 1]

# p
a = [1, 1, 0, 0]

# p OR (NOT q)
a = [1, 1, 0, 1]

# p OR q
a = [1, 1, 1, 0]

# True
a = [1, 1, 1, 1]
```

## Bit Shift

The bit shift are sometimes considered bitwise operations, because
they treat a value as a series of bits rather than as a numerical
quantity. In these operations the digits are moved, or shifted, to the
left or right. Registers in a computer processor have a fixed width,
so some bits will be "shifted out" of the register at one end, while
the same number of bits are "shifted in" from the other end; the
differences between bit shift operators lie in how they determine the
values of the shifted-in bits.

### Arithmetic shift

When shifting left includes a zero in the right most place.
When shifting right, the left most position gets carried to resulting
array.

### Logical shift

Includes zeros for both left and right shifts.

### Circular shift

The are rotations, which act as if the registers were joined on either
end.
