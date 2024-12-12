## [Step-by-step Processes/How do I do X?](how-to)

## Understand Rust

- [How is Rust code organized?](#code-organization)
- [Essential syntax and style](#basic-syntax-and-style)
- [Ownership and borrowing](#ownership-essentials)

# Code Organization

Packages contain crates, which contain modules.

Use Cargo. Cargo will expect your code to be in the `src` directory.

`cargo.toml` is for dependencies, test and build instructions, etc (equiv to `package.json` for node)

### Crates

Crates are the smallest unit of code that compilers considers at once. Crates can either be _library_ crates or _binary_ crates (think of them like C# projects, which can be either library/DLL projects or executable projects). Binary crates must have a `main()` function.

All crates have a crate _root_ file. By convention, Cargo considers `src/main.rs` the crate root of a binary crate with the same name as the package. The same applies for library crates, with the crate root considered to be `src/lib.rs`.

Crates can be published on `crates.io`

### Packages

Packages are bundles of one or more crates, and might best be thought of like C# solutions (I think). Packages contain the `cargo.toml` file, and are what get created by `$ cargo new`. A package can contain any number of binary crates, but only one library crate. If a package contains multiple binary crates, they should each be a separate file in `src/bin/`

### Modules

[Modules Cheat Sheet](language-concepts/code-organization.md#modules-cheat-sheet)

Modules are sort of like C# namespaces (I think). Include modules in your crates with

```rs
pub mod my_module;
```

Then define the module in a separate source file.

### Workspaces

### Access Modifiers

Modules as well as the definitions contained within them are private by default. Expose them by prefixing with the `pub` keyword.

### [More](./language-concepts/code-organization.md)

**For more about code organization, go [here](./language-concepts/code-organization.md).**

# Documentation

`$ rustup doc` and `$ rustup doc --book` work offline.

Run `$ cargo doc --open` to build and display documentation for your specific project locally.

# Running

You can use the Rust compiler `rustc` manually, but there isn't a good reason not to just use Cargo for everything instead.

From your project's root, use the appropriate command from the list:

```
$ cargo run
```

Or, to build without running (defaults to debug configuration):

```
$ cargo build
```

or

```
$ cargo build --release
```

Or, to check for problems without having to wait for a long build:

```
$ cargo check
```

# Basic Syntax and Style

Bring in external code like:

```rs
use std::io;
use std::fmt::Result;
use std::io::Result as IoResult;
use std::{cmp::Ordering, io}; // nested path
use std::io::{self, Write}; // nested path
use std::collections::*; // glob
pub use crate::front_of_house::hosting; // re-export
...
io::stdin().read_line(&mut my_input)
```

You can also reference external code without a `use` by utilizing the fully-qualified name, i.e. `std::io::stdin`. When using the `use` keyword, it is idiomatic to bring in a function's parent module but _not_ the function itself, so a namespaced path with `::` is necessary to call the function, making it clear that the function is defined externally. The same is _not_ true for structs and enums, however.

More syntax:

```rust
fn main() {
    let name = "world;"
    println!("Hello, {}!", name);
    println!("Hello, {name}!"); // also valid
}

// main can return a Result as well:
fn main() -> Result<(), Box<dyn Error>> {
    let greeting_file = File::open("hello.txt")?;
    Ok(())
}

fn data_structures() {
    // tuple
    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    let one = z.2

    // array (stack-allocated fixed-size)
    // type specifier optional
    let half_dozen: [i32; 5] = [1, 2, 3, 4, 5, 6];
    // write `[3, 3, 3, 3, 3]` as:
    let a = [3; 5];

    // slice
    let a = [1, 2, 3, 4, 5];
    let slice = &a[1..3];

    // vector
    let mut v: Vec<i32> = Vec::new();
    let mut w = vec![1, 2, 3];
    w.push(4);
    let does_not_exist = &v[100]; // panics
    let does_not_exist = v.get(100); // returns option
    for n_ref in &v {
        // n_ref has type &i32
        let n_plus_one: i32 = *n_ref + 1;
        println!("{n_plus_one}");
    }
    for n_ref in &mut v {
        // n_ref has type &mut i32
        *n_ref += 50;
    }

    // string and string literals
    let mut s = String::new() // strings are special vectors
    let s = "initial contents".to_string();
    let s = String::from("hello world");
    s.push_str("bar");
    s.push('l');
    let s1 = String::from("Hello, ");
    let s2 = String::from("world!");
    let s3 = s1 + &s2; // actually standard, s1 moved here
    let s1 = String::from("tic");
    let s2 = String::from("tac");
    let s3 = String::from("toe");
    // this doesn't take ownership:
    let s = format!("{s1}-{s2}-{s3}");
    // string literals are slices (type &str)!
    let hello: &str = &s[..5];
    let world: &str = &s[6..11];
    // Generally, you should use &str in function arguments and
    // returns when you can (but you can't always).

    // hash map: `use std::collections::HashMap;`
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    let team_name = String::from("Blue");
    let score = scores.get(&team_name).copied().unwrap_or(0);
    for (key, value) in &scores { // arbitrary order
        println!("{key}: {value}");
    }
    scores.entry(String::from("Yellow")).or_insert(50);
    let text = "hello world wonderful world";
    let mut map = HashMap::new();
    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }

}

fn funcs_and_expressions(a: i32) -> i32 {
    if a == 1 {
        return 0; // early return
    }
    let y = {
        let x = 3;
        x + 1 // no semicolon
    };
    5 // no semicolon, no return

    // define functions as f(s: &str)
    // not as f(s: &String
    // the former works for all cases
}

fn control_flow(condition: bool) {
    let number = if condition { 5 } else { 6 };
    // types of loops:
    // loop; labels allowed for disambiguation
    let mut counter = 0;
    let result = loop {
        counter += 1;
        if counter == 10 {
            break counter * 2;
        }
    };
    // while
    while condition {
        println!("again!");
    }
    // for
    let a = [10, 20, 30, 40, 50];
    for element in a {
        println!("the value is: {element}");
    }
    for number in (1..4).rev() { // reverse!
        println!("{number}!");
    }
    // match
    match coin { // match on a ref `&coin` to avoid moves
        Coin::Penny => {
            println!("Lucky penny!");
            1
        }
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        }
        _ => 0 // catch-all can be named and used, too
        // can also do nothing with `()`
    }
}
}
```

# Ownership Essentials

Every value has an owner and can have just one owner at a time. Once a value has been moved, the previous owner is invalidated.

Moves of heap allocated values occur when:

- You assign an existing value to another variable
- You pass a value to a method or function
- You return a value from a function

```rust
    let s1 = String::from("hello");
    let s2 = s1; // <-- Move!!
    println!("{}, world!", s1); // <-- Compile error! s1 is invalid here!
```

### [More](language-concepts/idiosyncracies/ownership.md)

**For more about ownership and borrowing, go [here](language-concepts/idiosyncracies/ownership.md).**

# Option Type

```rs
let some_number = Some(5);
let some_char = Some('e');
let absent_number: Option<i32> = None;
```

# Essential Tools

- **rustc**: Rust compiler (but just use cargo)
- **cargo**: package manager
- **clippy**: linting
- **rustfmt**: auto-format

# Tool and Language Management

Manage tools and versions with **rustup**, equiv. to NVM for Node.

Use `$ rustup` with no args for an overview of usage.

Use `$ rustup show` for Rust version info.

Use `$ rustup update` to get latest Rust toolset.

Use `$ rustup doc` to read the Rust docs offline
