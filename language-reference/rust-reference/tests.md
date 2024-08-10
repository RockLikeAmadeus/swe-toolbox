# Tools

Add the `pretty_assertions` crate for more friendly test output.

Also consider [nextest](https://nexte.st/) for a better testing experience.

# Integration Tests

The convention in Rust is to create a `tests` directory at the same level as your `src` directory.

## Testing command line binaries

For a command line program, you can test the actual execution of your code by directly executing the binary, like this:

```rs
use std::process::Command;

#[test]
fn it_works() {
    let mut cmd = Command::new("my_program");
    let res = cmd.output();
    assert!(res.is_ok());
}
```

This won't work on its own, since `my_program` doesn't exist in the same direrctory as this test. Add the crate `assert_cmd` to the project and change the test code to:

```rs
use assert_cmd::Command;

#[test]
fn runs_without_errors() {
    let mut cmd = Command::cargo_bin("my_program").unwrap();
    cmd.assert().success();
}
```

You should also test that your command line programs return 0, indicating a successful run, which `assert().success()` takes care of. Making sure your program returns a response code helps to make it composable with other CLI programs. Test for expected failures with

```rs
#[test]
fn return_error_code() {
    let mut cmd = Command::cargo_bin("my_program").unwrap();
    cmd.assert().failure();
}
```

### Verifying the output of a command line program

```rs
use assert_cmd::Command;
use pretty_assertions::assert_eq;

#[test]
fn produces_expected_output() {
    let mut cmd = Command::cargo_bin("my_program").unwrap();
    let output = cmd.output().expect("Something went wrong");
    assert!(output.status.success());
    let stdout = String::from_utf8(out.stdout).expect("invalid UTF-8");
    assert_eq!(stdout, "Hello, world!\n");
}
```

# Writing Tests

Rust tests are just functions annotated with the `test` attribute (e.g. `#[test]`. Library projects created using Cargo have default scaffolding provided for automated tests.

The body of a test function typically does these three things:

1. Set up any data or state that the test needs to run.
2. Run the code that we want to test.
3. Assert that the results are what we expect.

Tests fail when something in the test function panics.

### Syntax

```rs
#[cfg(test)]
mod tests {
    #[test]
    fn exploration() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn another() {
        panic!("Make this test fail");
    }
}
```

##### Assertions

```rs
assert!()
// For custom types passed to these,
// use `#[derive(PartialEq, Debug)]`
assert_eq!() // equal
assert_ne!() // not equal
```

##### Custom Failure Messages

```rs
#[test]
fn greeting_contains_name() {
    let result = greeting("Carol");
    assert!(
        result.contains("Carol"),
        "Greeting did not contain name, value was `{}`",
        result
    );
}
```

##### Testing Handling of Error Conditions

```rs
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    // `expected` is optional, but usually wise.
    // Verifies the panic message contains provided string
    #[should_panic(expected = "less than or equal to 100")]
    fn guess_constructor_panics_when_greater_than_100() {
        Guess::new(200);
    }
}

pub struct Guess {
    value: i32,
}

impl Guess {
    pub fn new(value: i32) -> Guess {
        if value < 1 {
            panic!(
                "Guess value must be greater than or equal to 1, got {}.",
                value
            );
        } else if value > 100 {
            panic!(
                "Guess value must be less than or equal to 100, got {}.",
                value
            );
        }

        Guess { value }
    }
}
```

##### Using Result in Tests

```rs
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() -> Result<(), String> {
        if 2 + 2 == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
}
```

### Test Organization

# Running Tests

Run tests in parallel with application output suppressed:

```
$ cargo test
```

> Some command line options go to cargo test, and some go to the resulting test binary. To separate these two types of arguments, you list the arguments that go to cargo test followed by the separator -- and then the ones that go to the test binary.

```
$ cargo test --help
```

```
$ cargo test -- --help
```

Limit parallel test runs:

```
$ cargo test -- --test-threads=1
```

Prevent application output suppression

```
$ cargo test -- --show-output
```

Run a specific test:

```
$ cargo test my_test
```

Run all tests who's name contains a string:

> Also note that the module in which a test appears becomes part of the test’s name, so we can run all the tests in a module by filtering on the module’s name.

```
$ cargo test part_of_test_name
```

Ignore specific tests:

```rs
#[test]
#[ignore]
fn expensive_test() {
    // ...
}
```

Run only ignored tests:

```
$ cargo test -- --ignored:
```

Run all tests including ignored tests:

```
$ cargo test -- --include-ignored
```

# Test Organization

Come back to https://rust-book.cs.brown.edu/ch11-03-test-organization.html and decide if the idiomatic approach makes sense.
