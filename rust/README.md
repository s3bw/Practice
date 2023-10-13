# Rust

[Rust-By-Example](https://doc.rust-lang.org/rust-by-example)

## Primitives

### Scalar Types

* Signed Integers: i8, i16, i32, i64, i128, and isize (pointer size)
    Signed Integers are typically whole numbers which have a
    signed bit which indicate whether they are positive or
    negative. i8 is an 8-bit integer, so the values range
    from -128 to 127.
* Unsigned Integers: u8, u16, u32, u64, u128 and usize (pointer size)
    These are unsigned therefore they don't have a bit indicating
    whether they are positive or negative, therefore the represent
    whole numbers in the range 0 to 255 (for u8 as example).
* Floating point: f32, f64
    Approximations to real numbers, called a floating-point
    since the decimal point can "float" in different positions
    within the number. Typically used to represent values with
    fractional parts
    Single Precision (f32) and Double Precision (f64)
    f32's typically have 23 bits for the significand and 8 bits
    for the exponent.
* char: unicode scalar values like `'a'`, `'α'` and `'∞'` (4 bytes each)
* bool: either `true` or `false`
* the unit type `()`, whose only possible value is an empty tuple: `()`

Integers default to `i32` and float to `f64`.


### Compound Types

* Arrays like `[1, 2, 3]`
* Tuples like `(1, true)`

```rust
fn main() {
    // Variables can be type annotated.
    let logical: bool = true;

    // Regular annotation
    let a_float: f64 = 1.0;
    // Suffix annotation
    let an_integer   = 5i32;

    // Or a default will be used.
    let default_float = 3.0;  // `f64`
    let default_integer = 7;  // `i32`

    // A type can also be inferred from context.
    let mut inferred_type = 12;  // Type i64 is inferred from another line.
    inferred_type = 4294967296i64;

    // A mutable variable's value can be changed.
    let mut mutable = 12; // Mutable `i32`
    mutable = 21;

    // Error! The type of a variable can't be changed.
    mutable = true;

    // Variables can be overwritten with shadowing.
    let mutable = true;
}
```

...[progress](https://doc.rust-lang.org/rust-by-example/custom_types/enum/enum_use.html)...


### Traits
