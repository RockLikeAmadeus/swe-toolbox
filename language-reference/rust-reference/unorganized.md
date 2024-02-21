- Build tool: Rust apparently needs something like the "make" tool. You can use make, or there are alternatives like `cargo-make`, or `cargo-cmd`, or as the author of "From JavaScript to Rust" recommends, `just`
- Workspaces: you can manage workspaces for managing multiple small modules in a larger project with the `[workspace]` entry in your `cargo.toml` file
  - Also check out `cargo-workspaces`
- You might want to install `cargo-expand` for demystifying some uses of macros

- Creating a new project
  projects can be created by just creating a directory for your project and adding a `main.rs` file and manually doing a `git init`. Cargo is not necessary for running code either - it can be done with `rustc main.rs` and then just run the resulting executable

come back to: https://rust-book.cs.brown.edu/ch09-02-recoverable-errors-with-result.html
