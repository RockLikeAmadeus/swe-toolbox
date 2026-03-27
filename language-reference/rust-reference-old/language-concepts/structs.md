```rs
#[derive(Debug)] // derived trait
#[derive(Display)] // for end-user
struct User {
    active: bool,
    username: String, // using owned String rather than
    email: String, // &str, to avoid need for lifetimes
    sign_in_count: u64,
}

fn structs() {
    // Instantiate
    let mut user1 = User {
        email: String::from("someone@example.com"),
        username: String::from("someusername123"),
        active: true,
        sign_in_count: 1,
    };

    // Set properties
    user1.email = String::from("anotheremail@example.com");

    let user2 = User { // <-- move here, user1 invalid
        email: String::from("another@example.com"),
        ..user1
    };

    println!("user2 is {:?}", user2); // debug display
    println!("user2 is {:#?}", user2); // detailed debug
    dbg!(&user2); // more detailed, pass ref to avoid move
}

fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username,
        email,
        sign_in_count: 1,
    }
}
```

# Tuple Structs

```rs
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
    let black = Color(0, 0, 0);
    let origin = Point(0, 0, 0);
}
```

# Methods and Associated Functions

```rs
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }

    fn set_width(&mut self, width: u32) {
        self.width = width;
    }

    // Associated function (like static, I think)
    fn make_square(size: u32) -> Self {
        Self {
            width: size,
            height: size,
        }
    } // call with Rectangle::make_square()
}
```

# Constructors

There is no keyword for declaring constructors, but the method name `new` is idiomatic. However, many users expect a `new()` function to never fail, so if the return value of the constructor is instead of type `Result<T>`, we should use the name `build` instead.

```rs
fn build(args: &[String]) -> Result<Config, &'static str> {
    ...
}
```
