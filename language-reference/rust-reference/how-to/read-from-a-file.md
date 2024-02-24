```rs
use std::fs;

fn main() {
    let file_path = "test.txt"
    let contents = fs::read_to_string(file_path)
        .expect("An error occurred when attempting to read from the specified file");
}
```
