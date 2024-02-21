```rs
pub trait Summary {
    fn summarize(&self) -> String;
}

pub struct NewsArticle {
    pub headline: String,
    pub location: String,
    pub author: String,
    pub content: String,
}

impl Summary for NewsArticle {
    fn summarize(&self) -> String {
        format!("{}, by {} ({})", self.headline, self.author, self.location)
    }
}

pub struct Tweet {
    pub username: String,
    pub content: String,
    pub reply: bool,
    pub retweet: bool,
}

impl Summary for Tweet {
    fn summarize(&self) -> String {
        format!("{}: {}", self.username, self.content)
    }
}
```

### Defining and Using a Default Implementation

```rs
pub trait Summary {
    fn summarize(&self) -> String {
        String::from("(Read more...)")
    }
}

impl Summary for NewsArticle { }
```

This is also allowed:

```rs
pub trait Summary {
    fn summarize_author(&self) -> String;

    fn summarize(&self) -> String {
        format!("(Read more from {}...)", self.summarize_author())
    }
}

impl Summary for Tweet {
    fn summarize_author(&self) -> String {
        format!("@{}", self.username)
    }
}
```

from the Rust book

>>> it isn’t possible to call the default implementation from an overriding implementation of that same method

# Accepting Polymorphic Arguments

```rs
pub fn notify(item: &impl Summary) {
    println!("Breaking news! {}", item.summarize());
}
```

The above is syntactic sugar for the more verbose "Trait Bound Syntax", which is more appropriate in some cases:

```rs
pub fn notify<T: Summary>(item: &T) {
    println!("Breaking news! {}", item.summarize());
}
```

### Requiring Multiple Traits on Arguments

```rs
pub fn notify(item: &(impl Summary + Display)) { ... }
```

and

```rs
pub fn notify<T: Summary + Display>(item: &T) { ... }
```

#### Specifying More Complex Trait Bounds With `where` Syntax

```rs
fn some_function<T, U>(t: &T, u: &U) -> i32
where
    T: Display + Clone,
    U: Clone + Debug,
{ ... }
```

# Returning Polymorphic Arguments

```rs
fn returns_summarizable() -> impl Summary {
    Tweet {
        username: String::from("horse_ebooks"),
        content: String::from(
            "of course, as you probably already know, people",
        ),
        reply: false,
        retweet: false,
    }
}
```

### Conditionally Implementing Traits for Types That Implement Other Traits

```rs
impl<T: Display> ToString for T {
    // --snip--
}
```

# Traits Aren't Quite Interfaces

Traits are essentially what interfaces are in most other languages, but have a few unique aspects.