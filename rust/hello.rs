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
