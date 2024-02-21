Lifetimes are a special case of generics. The primary aim of lifetimes are to prevent dangling references.

The borrow checker enforces that all values _live longer than any references to themselves_.

As a prerequisite to understanding this, see [The Scope of References](./ownership.md#critical-the-scope-of-references).

Lifetime specifiers are necessary when the borrow checker can't know how long a reference lives--that is, when they are passed into or returned from a function.

**The following function will not compile, with the `expected named lifetime parameter` error:**

```rustfn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```

When you receive the compiler error, `expected named lifetime parameter`, this just means that you have to indicate the relationship between the references so the compiler can ensure that the value having its reference passed in lives as long as the reference that's being returned. If the compiler looks at _just_ the signature (and not the body) of the above function, it isn't clear. Even if it did look at the body, however, the pointed-to value is determined at runtime.

Lifetime specifiers can be any identifier that starts with a `'` character, such as `'my_lifetime_specifier`, but typically they are short and we use something like `'a` or `'b`.

```rs
&i32        // a reference
&'a i32     // a reference with an explicit lifetime
&'a mut i32 // a mutable reference with an explicit lifetime
```

# Lifetime Annotations for Functions

From the Rust book:
> We want the signature to express the following constraint: the returned reference will be valid as long as both the parameters are valid.

```rs
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```

> The function signature now tells Rust that for some lifetime 'a, the function takes two parameters, both of which are string slices that live at least as long as lifetime 'a. The function signature also tells Rust that the string slice returned from the function will live at least as long as lifetime 'a.  In practice, it means that the lifetime of the reference returned by the longest function is the same as the smaller of the lifetimes of the values referred to by the function arguments.
...
When returning a reference from a function, the lifetime parameter for the return type needs to match the lifetime parameter for one of the parameters.

# Lifetime Annotations for Structs

Lifetime annotations are necessary when structs contain references. Lifetime specifiers are necessary as a generic argument to the struct as well as on every field of the struct that holds a reference.

```rs
struct ImportantExcerpt<'a> {
    part: &'a str,
}

impl<'a> ImportantExcerpt<'a> {
    fn level(&self) -> i32 {
        3
    }
}

fn main() {
    let novel = String::from("Call me Ishmael. Some years ago...");
    let first_sentence = novel.split('.').next().expect("Could not find a '.'");
    let i = ImportantExcerpt {
        part: first_sentence,
    };
}
```

> This annotation means an instance of ImportantExcerpt can’t outlive the reference it holds in its part field.

# Lifetime Elision

Lifetime annotations are not necessary when the compiler might be able to automatically determine input and output lifetimes. See [Lifetime Elision](https://rust-book.cs.brown.edu/ch10-03-lifetime-syntax.html). This typically applies to functions that take a single reference argument, or to most functions that are technically methods.

# The Static Lifetime

All string literals have the `'static` lifetime. References annotated with `'static` can live for the entire duration of the program.

# Lifetime Annotations Combined With Generic Arguments

```rs
use std::fmt::Display;

fn longest_with_an_announcement<'a, T>(
    x: &'a str,
    y: &'a str,
    ann: T,
) -> &'a str
where
    T: Display,
{
    println!("Announcement! {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```

# Things to Keep in Mind About Error Messages

Keep in mind that the `expected named lifetime parameter` error can _also_ occur when we try and return a reference to a value that will be dropped at the end of the function, such as in this case:

```rust
fn return_a_string() -> &String {
    let s = String::from("Hello world");
    let s_ref = &s;
    s_ref // <-- error: s is invalid here
}
```

---
---
---

> Most of the time, an error message suggesting the 'static lifetime results from attempting to create a dangling reference or a mismatch of the available lifetimes. In such cases, the solution is fixing those problems, not specifying the 'static lifetime.