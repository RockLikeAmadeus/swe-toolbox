# Command Line Arguments

If your program accepts command line arguments, you can pass them in when you run via Cargo with `$ cargo run -- first-arg second-arg`

The typical way to handle command line arguments in Rust is to utilize the `clap` (command line argument parser) crate, but to handle this from scratch:

```rs
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    dbg!(args);
}
```

## With clap

```bash
$ cargo add clap
```

Clap provides two distinct APIs for handling command line arguments in your application: the _derive_ API and the _builder_ API. According to [this]([source](https://medium.com/@itsuki.enjoy/rust-take-your-cli-to-the-next-level-with-clap-a0f05875ef45)) source, the differences are

> Derive API:
>    - Easier to read, write, and modify
>    - Easier to keep the argument declaration and reading of argument in sync
>    - Easier to reuse
> 
> Builder API
>    - Faster compile times if you aren’t already using other procedural macros
>    - More flexibility, e.g. you can look up an arguments values, their ordering with other arguments, and what set them. The Derive API can only report values and not indices of or other data.
> 
> It’s also worth mentioning that the two methods can be mixed together.

### Derive API

To do...

### Builder API

The **clap** crate defines a `Command` struct (see [docs](https://docs.rs/clap/latest/clap/struct.Command.html)), and instantiating an instance of `Command` represents defining your actual CLI. With that in mind, the simplest way to define your CLI using `clap` would look like this:

```rs
use clap::Command;

fn main() {
    let _matches = Command::new("simple-echo")
        .version("0.1.0")
        .author("Alec Hampton")
        .about("Simplified implementation of the `echo` command")
        .get_matches();
}
```

There are macros and factories provided to help with this instantiation as well.

# Response Codes

Making sure your program returns a response code helps to make it composable with other CLI programs. Response codes should be integer values from 0 to 255 for POSIX compatible interfaces, with 0 indicating success.