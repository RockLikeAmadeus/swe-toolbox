```rs
enum IpAddrKind {
    V4,
    V6,
}

fn main() {
    let four = IpAddrKind::V4;
    let six = IpAddrKind::V6;
}
```

# Embedding Values and Adding Methods

When designing Rust code, I think it's important to think of enums more broadly than you would in other languages. There are lots of cases that would call for a struct or class in another language that would be best handled by an enum in Rust, given these additional capabilities (as real-world examples of such cases come up, note them here and potentially blog about this point):

```rs
enum IpAddr {
    // Any type can be embedded, even structs and enums
    V4(u8, u8, u8, u8),
    V6(String),
}

fn main() {
    // Each variant becomes a constructor function!
    let home = IpAddr::V4(127, 0, 0, 1);
    let loopback = IpAddr::V6(String::from("::1"));
}

enum Message {
    Quit,
    // Can even have named fields
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

// enums can have methods!
impl Message {
    fn call(&self) {
        // method body would be defined here
    }
}
```
