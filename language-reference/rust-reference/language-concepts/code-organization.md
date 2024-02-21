# Common Library Structure

A common pattern is for a package to contain a single binary crate and along with a single library crate. This pattern is typically incarnated with the behavior being exposed by the library, and the binary being a default client for that library. The binary crate will have just enough code to interface with the user and with the library behind it. See the "Best Practices" note on [this](https://rust-book.cs.brown.edu/ch07-03-paths-for-referring-to-an-item-in-the-module-tree.html) page.

# Modules Cheat Sheet

Pulled directly from the [Rust book](https://rust-book.cs.brown.edu/ch07-02-defining-modules-to-control-scope-and-privacy.html):

- Start from the crate root: When compiling a crate, the compiler first looks in the crate root file (usually `src/lib.rs` for a library crate or `src/main.rs` for a binary crate) for code to compile.
- Declaring modules: In the crate root file, you can declare new modules; say, you declare a “garden” module with `mod garden;`. The compiler will look for the module’s code in these places:
  - Inline, within curly brackets that replace the semicolon following mod garden
  - In the file `src/garden.rs`
  - In the file `src/garden/mod.rs`
- Declaring submodules: In any file other than the crate root, you can declare submodules. For example, you might declare mod vegetables; in src/garden.rs. The compiler will look for the submodule’s code within the directory named for the parent module in these places:
  - Inline, directly following mod vegetables, within curly brackets instead of the semicolon
  - In the file `src/garden/vegetables.rs`
  - In the file `src/garden/vegetables/mod.rs`
- Paths to code in modules: Once a module is part of your crate, you can refer to code in that module from anywhere else in that same crate, as long as the privacy rules allow, using the path to the code. For example, an `Asparagus` type in the garden vegetables module would be found at `crate::garden::vegetables::Asparagus`.
- Private vs public: Code within a module is private from its parent modules by default. To make a module public, declare it with `pub mod` instead of `mod`. To make items within a public module public as well, use `pub` before their declarations.
- The `use` keyword: Within a scope, the `use` keyword creates shortcuts to items to reduce repetition of long paths. In any scope that can refer to `crate::garden::vegetables::Asparagus`, you can create a shortcut with `use crate::garden::vegetables::Asparagus;` and from then on you only need to write `Asparagus` to make use of that type in the scope.

# Paths

Paths are for telling Rust where to find entities in the module tree. It works similarly to file paths, with absolute and relative paths possible.

Absolute path (from current crate root): `crate::`

Absolute path (from external crate): `my_crate::`

Relative path (uses `self` or `super` or identifier in current module). `super` is equivalent to `..` in a file path.

e.g.

```rs
mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}
    }
}

pub fn eat_at_restaurant() {
    // Absolute path
    crate::front_of_house::hosting::add_to_waitlist();

    // Relative path
    front_of_house::hosting::add_to_waitlist();
}
```

# Multiple Files

https://rust-book.cs.brown.edu/ch07-05-separating-modules-into-different-files.html
