

- When you're building a type in your library that is meant to be exposed to consumers of your library, it's good practice to implement the built-in Rust trait `Default`
- `(0..1000).for_each(|_| { ....` is the same as `for _ in 0..1000`
- Rust ownership is a completely new _model_ for thinking about memory (rather than thinking about memory as just an array of bytes).
    - Stack is made up of frames which are made up of variables mapped to values (or pointers).
    - Heap is made up of boxes.