If your program accepts command line arguments, you can pass them in when you run via Cargo with `$ cargo run -- first-arg second-arg`

There do exist libraries on crates.io that can be used, but to handle this from scratch:

```rs
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    dbg!(args);
}
```
