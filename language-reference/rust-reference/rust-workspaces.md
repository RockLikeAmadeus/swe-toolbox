# Advantages of [Workspaces](https://doc.rust-lang.org/cargo/reference/workspaces.html)

- Manage multiple related libraries and applications together
- Compiled dependencies are shared between projects in the same workspace, saving compilation time and disk space

# Working with Rust Workspaces

Workspaces are pretty simple - you just

``` bash
$ cargo new myworkspace
$ cd myworkspace
$ cargo new mynewbinaryproject
```

And then in the root `Cargo.toml`, add

```toml
[workspace]
members = [ "mynewbinaryproject" ]
```

It's very easy to unintentionally `$ cargo run` your root workspace project when you intended to run a subproject. One way to be more aware when this happens is to edit the root-level project's `main.rs` file and have it output something to the effect of "Running workspace root project. Did you mean to run a workspace member?"