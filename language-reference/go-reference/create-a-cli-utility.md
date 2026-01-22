# Setting up the project

- Make a new folder for the utility (module)
- `cd` into the new directory
- Make a new folder called `cmd`. Do not `cd` into this.
- `$ go mod init github.com/RockLikeAmadeus/myproject`
- Create the rest of the project structure according to [Code Organization](#code-organization)

# Code organization

A common pattern for creating a CLI is to have a main package for the CLI executable, and a separate package for the business logic. This way you (or other developers) can re-use the behavior of the tool via another interface (even a different style of CLI). The business logic package sits at the module's root, and the `cmd` package is in a folder beside it. It looks like this:

![alt text](./images/cli-tool-code-organization.png)

It's also a common pattern to have the `main()` function simply parse command line arguments, then call a `run()` function that is responsible for the actual program logic (or for calling into functions that define the logic). If there are many arguments to pass to `run()`, consolidate them into one or more `struct`s, perhaps called `config`. This pattern can be easy for testing, since unlike `main()`, the `run()` function can be called by your tests, although I prefer running the CLI tool itself for a full end-to-end acceptance test.

# Using Cobra

First create the directory and initialize the go module for the project, then run:

```
$ go get -u github.com/spf13/cobra@latest
```

If you want to use the boilerplate code generation that Cobra offers, you also want to run:

```
$ go get -u github.com/spf13/cobra-cli@latest
$ cobra-cli init
$ go run main.go
```

To avoid having to specify author and license information every time you run a cobra CLI command, you can add a .cobra.yaml file to your home directory (try doing this with cobra command)