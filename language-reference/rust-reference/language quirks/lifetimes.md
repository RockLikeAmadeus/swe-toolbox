The borrow checker enforces that all values _live longer than any references to themselves_.

Lifetime specifiers are necessary when the borrow checker can't know how long a reference lives--that is, when they are passed into or returned from a function.

**The following function will not compile, with the `expected named lifetime parameter` error:**

```rust
fn first_or(strings: &Vec<String>, default: &String) -> &String {
    if strings.len() > 0 {
        &strings[0]
    } else {
        default
    }
}
```

When you receive the compiler error, `expected named lifetime parameter`, this just means that you have to indicate which of the incoming arguments the return value of the function points to, so the compiler can ensure that the value having its reference passed in lives as long as the reference that's being returned. If the compiler looks at _just_ the signature (and not the body) of the above function, it isn't clear. But even if it did look at the body, the pointed-to value is determined at runtime.

Keep in mind that the `expected named lifetime parameter` error can _also_ occur when we try and return a reference to a value that will be dropped at the end of the function, such as in this case:

```rust
fn return_a_string() -> &String {
    let s = String::from("Hello world");
    let s_ref = &s;
    s_ref // <-- error: s is invalid here
}
```