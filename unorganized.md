# Rust

Come back to VS Code setup section! TODO

for all languages, create a (short as possible) example code file that's just hello world with as many language features on display as possible, with comments, in a way that is easy to just grab syntax all from one place. For Rust, start with this for examples etc.: https://learnxinyminutes.com/docs/rust/

- Note, everything it creates for you is pretty much what's already in the code crafter's basic solution, with the exception of the stuff only code crafters uses, and the commented out code in main, so that's good.
- Strings
    - Two types of strings, `str` and `String`
- Ownership
    - in Rust, you have to be thinking about whether a value is on the stack or the heap (maybe you don't actually, but you have to know when values are allocated on the heap, because ownership's job is to manage memory on the heap)
    - Rules of ownership
        - Each value in rust has an owner, and can only have one at a time
        - When the owner goes out of scope, the value will be dropped
    - Ownership allows us to treat heap values like stack values--the memory taken up by stack variables is automatically freed when the variable goes out of scope. This is harder with heap variables in most languages because multiple references can exist, so when a variable goes out of scope is not sufficient to say that value isn't used anymore. Rust solves this by only allowing one variable to "own" that value, and then treats it like a stack variable.
    - Use `clone()` to deep copy:
        ```rust
        let s1 = String::from("hello");
        let s2 = s1.clone(); // <-- Here, the copy is explicit.
        // (Above) This is the same behavior as if these were primitive types with simple copying, but we want the explicit call because `clone()` could be expensive.
        println!("s1 = {}, s2 = {}", s1, s2);
        ```
    - Types annotated with the  `Copy` trait (interface) for types that are stored on the stack. If a type implements `Copy`, variables that use it don't perform a _move_, because the copy is trivial
    - A thought of my own: dealing with all of this complexity is probably much easier to think through if we keep our functions small.
    - Of course, sometimes we want to assign a value or call a function but we don't want to _move_ the value into the new variable--we want to maintain the ownership. This is where _borrowing_ and _references_ come in.
        - A reference is like a pointer _except that a reference is guaranteed to point to a valid value_.
        - Example:
        ```rs
        fn main() {
            let s1 = String::from("hello");
            let len = calculate_length(&s1); // <-- Pass a reference to s1
            println!("The length of '{}' is {}.", s1, len);
        }

        fn calculate_length(s: &String) -> usize { // <-- Accept a reference to a string
            s.len()
        } // s goes out of scope here, but doesn't own its value, so it's not dropped
        ```
    - We can't modify something we're borrowing without a _mutable reference_.
        - The major restriction with mutable references is that if you have a mutable reference to a value, no other references to that value can exist simultaneously.
    - A reference's scope is, I think, different than the scope of other types, and I think this is really important, though confusing. A reference's scope starts from where it is introduced and continues through the last time that reference is used
        Example:
        ```rust
        let mut s = String::from("hello");

        let r1 = &s; // no problem
        let r2 = &s; // no problem
        println!("{} and {}", r1, r2);
        // variables r1 and r2 will not be used after this point

        let r3 = &mut s; // no problem
        println!("{}", r3);
        ```
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

# Python



---

? to ask WF, what format will it come in, is it changeable and updatable by me? not just a PDF right?