# Code Organization

`cargo.toml` for dependencies, test and build instructions, etc (equiv to `package.json` for node)

Crates.

Crates can be published on `crates.io`

# Documentation

`$ rustup doc --book`

# Essential Tools

**rustc**: Rust compiler (I think cargo is enough, though)
**cargo**: package manager
**clippy**: linting?
**rustfmt**: auto-format

# Basic Syntax and Style

```rust
fn main() {
    let name = "world;
    println!("Hello, {}!", name);
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

# Tool and Language Management

Manage tools and versions with **rustup**, equiv. to NVM for Node.

Use `$ rustup` with no args for an overview of usage.

Use `$ rustup show` for Rust version info.

Use `$ rustup update` to get latest Rust toolset.