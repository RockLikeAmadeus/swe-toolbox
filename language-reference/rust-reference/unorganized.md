- You can run `$ rustup` with no arguments for an overview of usage. Rustup is a rust version manager which is the equivalent of Node's NVM. Rustup also manages your other Rust tools like clippy, cargo, and rustfmt
- `$ rustup show` shows your current rust version
- Get the latest with `$ rustup update`
- `$rustup doc --book` opens the Rust book in your browser
- rustc is the Rust compiler, Cargo is the package manager
- Rust crates are published on `crates.io`, and you can register an account and publish your own modules there
- your project's dependencies, and other matters like tests and build instructions, are managed by `cargo.toml`, the rust equivalent of Node's `package.json`
- Create new rust projects with `$ cargo new` and `$ cargo init`. The only difference is that `init` will use your current directory as the project directory, and `new` will expect the name of a new directory which it will create.
- Add dependencies to project with `$ cargo add my-dependency`, much like npm's `$npm install my-dependency`, but you have to have installed `cargo-edit` first: `$ cargo install cargo-edit`.
- `$ cargo install my-dep` is the Rust equivalent of globally installing in Node with `$npm install --global my-dep`
- Build tool: Rust apparently needs something like the "make" tool. You can use make, or there are alternatives like `cargo-make`, or `cargo-cmd`, or as the author of "From JavaScript to Rust" recommends, `just`