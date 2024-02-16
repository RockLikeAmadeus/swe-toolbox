# Code Organization

Use Cargo. Cargo will expect your code to be in the `src` directory.

`cargo.toml` for dependencies, test and build instructions, etc (equiv to `package.json` for node)

Crates.

Crates can be published on `crates.io`

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
...
io::stdin().read_line(&mut my_input)
```

You can also reference external code without a `use` by utilizing the fully-qualified name, i.e. `std::io::stdin`

More syntax:
```rust
fn main() {
    let name = "world;"
    println!("Hello, {}!", name);
    println!("Hello, {name}!"); // also valid
}

fn tuples_and_arrays(i: i32) {
    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    let one = z.2

    // stack-allocated fixed-size array
    // type specifier optional
    let half_dozen: [i32; 5] = [i, 2, 3, 4, 5, 6];
    // write `[3, 3, 3, 3, 3]` as:
    let a = [3; 5];
}

fn funcs_and_expressions(a: i32) -> i32 {
    if (a == 1) {
        return 0; // early return
    }
    let y = {
        let x = 3;
        x + 1 // no semicolon
    };
    5 // no semicolon, no return
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

# Essential Tools

**rustc**: Rust compiler (I think cargo is enough, though)
**cargo**: package manager
**clippy**: linting?
**rustfmt**: auto-format

# Tool and Language Management

Manage tools and versions with **rustup**, equiv. to NVM for Node.

Use `$ rustup` with no args for an overview of usage.

Use `$ rustup show` for Rust version info.

Use `$ rustup update` to get latest Rust toolset.

Use `$ rustup doc` to read the Rust docs offline