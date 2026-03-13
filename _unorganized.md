# Rust

Come back to VS Code setup section of TypeScript to Rust book! TODO

for all languages, create a (short as possible) example code file that's just hello world with as many language features on display as possible, with comments, in a way that is easy to just grab syntax all from one place. For Rust, start with this for examples etc.: https://learnxinyminutes.com/docs/rust/

- Note, everything it creates for you is pretty much what's already in the code crafter's basic solution, with the exception of the stuff only code crafters uses, and the commented out code in main, so that's good.
- Strings
    - Two types of strings, `str` and `String`
- `.unwrap()` is fine in example code but really shouldn't be used in production code unless you're certain the operation you're unwrapping can't fail.
- Directly from the Rust book: "Defining a function to take a string slice instead of a reference to a String makes our API more general and useful without losing any functionality"
    `fn first_word(s: &str) -> &str {` instead of `fn first_word(s: &String) -> &str {`
- From _From JavaScript to Rust_, the suggestion is, when you want to turn a string literal into a String object, use the `.to_owned()` function, rather than `String::from()`, `to_string()`, `.into()`, or `format!()`. The explanation of why is not something I want to list here, but I'm sure there are resources on it.
- Rust variables use `snake_case`. That sucks, but you should probably jsut get used to it. See https://doc.rust-lang.org/1.0.0/style/style/naming/README.html for more.
- Whenever you see an error complaining about `()`, it's frequently a matter of adding or removing a semi-colon somewhere, or a return type on a function declaration.
- When to use which type of collection, handy reference: https://doc.rust-lang.org/std/collections/index.html#use-a-btreemap-when
- When an error mentions `&'static str`, it's probably referring to a string literal. Remember string literals are not type `str`, they're type `&str`

# Visual Studio

Some things Josh taught me I might want to be aware of in visual studio., related to debugging multiple process at the same time in the same instance of visual studio.

- In the past I've either tried to have two instances of the same solution open (definitely causes problems) or have multiple startup projects selected. This is not optimal. Instead, to get a second process running with the debugger attached, you can:
    - Right click the project in the solution explorer, assuming it's an executable of some sort, and select Debug >> Start New Instance
    - Or, you can separately start up an instance of the process you're trying to debug (assuming it's a debug build with debug symbols, not a release build), and at the top of Visual Studio go to Debug >> Attach to Process and select a running process.
- Also, you can change the maximum number of build threads in the IDE settings
- Also, you can view the project Build Order by right clicking on the solution in the Solution Explorer

# Python



---

# Resource:

https://www.programming-books.io/

---

Think about _plugins_ vs _extensions_. In a nutshell, _plugins_ are a mechanism for polymorphic behavior, and _extensions_ (I think) are an incarnation of the open-closed principal (my words, and I could be wrong).
    How do these relate to "mods" in video games?
    It seems like there truly is a lot of conflation of these two terms, but maybe it's worth writing a post talking about a meaningful difference between the two and being precise in language.