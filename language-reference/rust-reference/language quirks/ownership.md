# My own rules of ownership:
- Each value in rust has an owner, and can only have one at a time
- When the owner goes out of scope, the value will be dropped
- Once a value has been moved, the previous owner is invalidated.
- Moves of heap allocated values occur when:
    - You assign an existing value to another variable
    - You pass a value to a method or function
    - You return a value from a function

Ownership allows us to treat heap values like stack values--the memory taken up by stack variables is automatically freed when the variable goes out of scope. This is harder with heap variables in most languages because multiple references can exist, so the moment when a variable goes out of scope is not sufficient to say that value isn't used anymore. Rust solves this by only allowing one variable to "own" that value, and then treats it like a stack variable.

# Cloning and Copying

Types annotated with the  `Copy` trait is for types that are stored on the stack. If a type implements `Copy`, variables that use it don't perform a _move_, because the copy is trivial

Use `clone()` to deep copy. The below example describes the same behavior as if these were primitive types using simple copying, but we want the explicit call because `clone()` could be expensive, so we want to call it out:

```rust
let s1 = String::from("hello");
let s2 = s1.clone(); // <-- Here, the copy is explicit.
println!("s1 = {}, s2 = {}", s1, s2);
```

# References and Borrowing

Of course, sometimes we want to assign a value or call a function but we don't want to _move_ the value into the new variable--we want to maintain the ownership. This is where _borrowing_ and _references_ come in.

A reference is like a pointer _except that a reference is guaranteed to point to a valid value_.

```rs
fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1); // <-- Pass a reference to s1
    println!("The length of '{}' is {}.", s1, len);
}

// Accept a reference to a string
fn calculate_length(s: &String) -> usize {
    s.len()
    // s goes out of scope here, but doesn't own its value, so it's not dropped
} 
```

### Mutable References

 We can't modify something we're borrowing without a _mutable reference_. The major restriction with mutable references is that if you have a mutable reference to a value, no other references to that value can exist simultaneously.

 ### The Scope of References

 A reference's scope is different than the scope of other types, and this is really important, although confusing. A reference's scope starts from where it is introduced and continues through _the last time that reference is used._ For example:

 ```rust
let mut s = String::from("hello");

let r1 = &s; // no problem
let r2 = &s; // no problem
println!("{} and {}", r1, r2);
// variables r1 and r2 will not be used after this point

let r3 = &mut s; // no problem
println!("{}", r3);
```

# Tips:

Dealing with all of this complexity is probably much easier to think through if we keep our functions small.