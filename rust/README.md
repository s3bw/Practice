# Rust

```bash
cargo build
./hello
```

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

### Scopes (Ownership, borrowing and lifetimes)

Indications to the compiler when borrows are valid, when resources
can be freed, and when variables are created or destroyed.

Variable in Rust do more than just hold data in the stack: they
also _own_ resources, e.g. `Box<T>` owns memory in the head. Rust
enforces `RAII` (Resource Acquisition Is Initialization), so
whenever an object goes out of scope, its destructor is called and
its owned resources are freed.
This behavior shields against _resource leak_ bugs, so you'll never
have to manually free memory or worry about memory leaks again!

<b>Destructor:</b> The notion of a destructor in Rust is provided
through the `Drop` trait. The destructor is called when the resource
goes out of scope. This trait is not required to be implemented for
every type, only implement it for your type if you require its own
destructor logic.

Ownership and moves. Because variables are in charge of freeing
their own resources, resources can only have one owner. This also
prevents resources from being freed more than once. Note that not
all variables own resources.

When doing assignments (`let x = y`) or passing function args by
value (`foo(x)`), the _ownership_ of the resource is transferred.
This is known as a _move_. (Rust-speak).

```rust
fn destroy(c: Box<i32>) {
    println!("Destroying a box that contains {}", c);
}

fn main() {
    let a = Box::new(5i32);
    println!("a contains: {}", a);

    // *Move* `a` into `b`
    // The pointer address of `a` is copied (not the data)
    // into `b`. Both are now pointers to the same heap
    // allocated data, but `b` now owns it.
    let b = a;

    // Error! `a` can no longer access the data, because it
    // no longer owns the heap memory
    //println!("a contains: {}", a);

    // This function takes ownership of the heap allocated
    // memory from `b`
    destroy(b);

    // Since the heap memory has been freed at this point,
    // the action would result in a dereferencing freed
    // memory, but it's forbidden by the compiler
    // Error! Same reason as the previous Error
    //println!("b contains: {}", b);
}
```

Mutability, of data can be changed when ownership is transferred.

A "borrow" refers to a way in which code can access data without
taking full ownership of it. Rust's borrowing system is a concept
designed to ensure memory safety and prevent data races. It plays
a crucial role in enforing rules such as the ownership, borrowing
and lifetimes of data.

1. Read-Only Borrows: These allow multiple parts of the code to
read the data at the same time. While data is immutably borrowed,
it cannot be modified until the borrow is released.

```rust
fn main() {
    let x = 42;
    let y = &x; // Immutable borrow
    println!("x: {}", x);  // Valid, because we're only reading x
}
```

2. Mutable Borrows: (Read/Write), There allow one part of your code
to modify the data while ensuring that no other part of the code is
simultaneously reading or modifying the data. There can be only one
mutable borrow at a time.

```rust
fn main() {
    let mut x = 42;
    let y = &mut x;  // Mutable borrow
    *y = 10;  // Valid, because we're modifying x through the mutable
              // borrow.
}
```

3. Interior Mutability: In some cases, you might want to mutate data
behind an immutable reference. This is possible with types like
`Cell`, `RefCell` and `Mutex`, which allow interior mutability while
preserving Rust's safety guarantees. This is considered advanced and
is typically used in certain safe patterns, such as for thread sync

```rust
use std::cell::Cell;

fn main() {
    let x = Cell::new(42);
    let y = &x  // Immutable borrow
    x.set(10);  // Valid, even though we have an immutable borrow.
}
```

### Explicit Typing

Given the two examples `.collect::<String>()` and `.collection()`.
Both are used to collect items into a string, but there is a key
difference between the two.

1. `.collect::<String>()`. This method explicitly specifies the
type you want to collect into, in this case, a `String`. It is
useful when you want to be explicit about the type of the collection

2. `.collect()`. This method infers the type of the collection from
the context in which it's used. In most cases, this will work just
fine, as the Rust compiler can usually infer the appropriate type.
It is more concise and is often used when the target type can be
easily determined from the surrounding code.

### Lifetime specifier



### Traits

...[progress](https://doc.rust-lang.org/rust-by-example/custom_types/enum/enum_use.html)...
