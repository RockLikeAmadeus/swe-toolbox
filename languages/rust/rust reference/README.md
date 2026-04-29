# Contents

- [Contents](#contents)
- [What makes Rust unique?](#what-makes-rust-unique)
- [What do I need installed to write Rust code](#what-do-i-need-installed-to-write-rust-code)
- [I just want to write a quick script with ](#i-just-want-to-write-a-quick-script-with-)
- [I'm starting a new software project with ](#im-starting-a-new-software-project-with-)
- [Syntax and Style](#syntax-and-style)
  - [Essential Syntax and Style](#essential-syntax-and-style)
  - [Comprehensive Syntax and Style](#comprehensive-syntax-and-style)
  - [Documenting  Code](#documenting--code)
- [Basics of Testing](#basics-of-testing)
- [Useful Library Imports](#useful-library-imports)
- [Essential Language Toolset Commands](#essential-language-toolset-commands)
- [Tools](#tools)
- [Utilities](#utilities)
- [How To](#how-to)

# What makes Rust unique?

Note: rewrite this later, this is just random notes to be re-organized.

- variables are mutable by default, use `mut` to declare them mutable.
- Rust has special syntax to return a value from a function without using the return keyword or a semicolon.
- Like some other languages, has two types of strings: `str` (string literals, entered in the actual source code), and dynamic `String`s that store location, length, and capacity.
- Methods are static (called "associated functions") unless the first parameter is `&self`.
  - In method implementations (inside the `impl` body), `Self` refers to the struct type itself, while `self` refers to the _instance_ of the structure.

_What are a couple of the most important high-level things to be aware of that make this language what it is? What features and quirks define the language and must be kept in mind at (pretty much) all times? Note: this section is not for notes on unique syntax--that section is below. Delete this block once the reference for this language is "complete"_

[Contents](#contents)

# What do I need installed to write Rust code

The best approach for installing Rust and it's various tools are to install [the RustUp _toolchain_](https://rustup.rs/). RustUp is more-or-less analogous to NVM for Node.

Use `$ rustup` with no args for an overview of usage.

Use `$ rustup show` for Rust version info.

Use `$ rustup update` to get latest Rust toolset.

Use `$ rustup doc` to read the Rust docs offline

[Contents](#contents)

# I just want to write a quick script with <language>

_Notes on writing and then running a standalone script or small process here. Include the simplest way to retrieve command line arguments. Delete this block once the reference for this language is "complete"_

[Contents](#contents)

# I'm starting a new software project with <language>

To create a new project in the current dir:

`$ cargo init`

To create a new project in a new dir:

`$ cargo new my_app`

Note: use `snake_case` for your project name. This can always be changed in `cargo.toml`.

Additional note: these commands create executable projects by default. Use `$ cargo new --lib` to bootstrap a library project, which may be easier to start with since it will include test scaffolding.

Additional Note: these commands _do_ generate a git repository automatically, unless one already exists. Change this by adding the `--vcs none` flag

To run your project, execute `$ cargo run`.

_Notes on the structure and terminology of more complex, multi-file projects here. Talk about how code is organized in <language>, how files are named, common patterns for specific types of projects, and tools commonly used for scaffolding new projects. Keep it brief, and leave the details to [this page I haven't created yet]. Delete this block once the reference for this language is "complete"_

[Contents](#contents)

# Syntax and Style

## Essential Syntax and Style

`snake_case` is used most often for naming variables in Rust.

_Examples of the most important syntax and style to know without having to go all the way to external resources.
Also include the most notable syntax-related things to know about that make this language different than other languages. Delete this block once the reference for this language is "complete"_

[Contents](#contents)

## Comprehensive Syntax and Style

_Links to relevant example syntax and style guides, like learnxinyminutes.com, or official style guides. Delete this block once the reference for this language is "complete"_

## Documenting <Language> Code

_What is the idiomatic way to document your code, especially any public API it might expose? What other details are relevant to documentation?_

[Contents](#contents)

# Basics of Testing

_What are the basics of automated tests for this language? If testing isn't built into the language, what testing frameworks are most popular? What is the basic syntax for writing automated tests, and how are test functions and files organized?_

[Contents](#contents)

# Useful Library Imports

_What are the most commonly included libraries, preferably from the language's standard library? Specify the syntax for including them, what each library is used for, and optionally, a link to the libraries documentation page._

[Contents](#contents)

# Essential Language Toolset Commands

_What commands with the languages compiler and/or included workflow command-line tools should be known, or at least be known of?_

[Contents](#contents)

# Tools

_What are some well-known and widely-used third party tools and frameworks that it's important to be aware of, such as CLI or TUI frameworks, or web frameworks? Be sure to include what they're used for and maybe some well established pros, cons, and use cases._

[Contents](#contents)

# Utilities

_What are some well-known and widely-used third party libraries that it's important to be aware of, such as logging, financial libraries, math libraries, or data structure utilities? Be sure to include what they're used for and maybe some well established pros, cons, and use cases._

[Contents](#contents)


# How To

_Define an organized list, with subheadings as appropriate, of links to pages in this reference *OR* to a project or file that serves as an example, of patterns or processes for which it would be beneficial to have a reference for demonstration. There is no reason two or more links cannot point to the same example, particularly when multiple patterns are represented in a single example project or file._

[Contents](#contents)