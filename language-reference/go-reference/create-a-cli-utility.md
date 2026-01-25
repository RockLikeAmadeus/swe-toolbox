# Quick Start

TBD

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

## Cobra CLI

If you want to use the boilerplate code generation that Cobra offers, you also want to run:

```
$ go install github.com/spf13/cobra-cli@latest
$ cobra-cli init
$ go run main.go
```

You can also run this command in a subdirectory, if you want to take advantage of having your application logic and your CLI interface as separate packages. For instance, generate the Go Module at your module's root, but run the `$ cobra-cli init` command inside of a new `cmd` directory (this naming could conflict with the `cmd` folder that the `init` command generates, meaning your commands will be defined in `app/cmd/cmd`, but this is probably alright).

If the cobra-cli command is not working, you may need to add the line `export PATH="/home/alec/go/bin:$PATH"` to the end of your .bashrc file in your `home` directory.

To avoid having to specify author and license information every time you run a cobra-cli  command, you can add a .cobra.yaml file to your home directory (run command `$ cobra-cli` to see where the default location is). The following contents are likely fine:

```yaml
author: Alec Hampton
license: Apache
```

For more info, see [the cobra README](https://github.com/spf13/cobra/tree/main) and [the cobra-cli README](https://github.com/spf13/cobra-cli/blob/main/README.md)

## Code Structure

 The Cobra CLI creates a simple `main.go` file that only imports the `cmd` package and executes the application. The core functionality resides in (or, at least, is called by) the `cmd` package, in which lies the `root.go` file which defines `Execute()` and the general structure of the program.

Commands and subcommands are implemented by instancing the main type defined by Cobra: `cobra.Command`. These commands can be compiled in a parent-child relationship to form a tree structure of subcommands.

In the pre-generated root command, which is executed when your CLI is run without any commands or subcommands, the `Run` property is commented out, since typical CLI behavior is for the root command to just offer usage info, but you can uncomment it if you want the root command to actually do something.

You can also add the `Version` property to the `rootCmd` to have Cobra automatically include the `-v` and `--version` flags.

```go
var rootCmd = &cobra.Command{
	Use:   "pScan",
	Short: "Fast TCP port scanner",
	Long: `pScan - short for Port Scanner - executes TCP port scan
on a list of hosts.

pScan allows you to add, list, and delete hosts from the list.

pScan executes a port scan on specified TCP ports. You can customize
the target ports using a command line flag.`,
	Version: "0.1.0",
}
```

Finally, in the `root.go` file you can implement the functions `init()` and `initConfi()` (the latter is probably only generated if you included Viper when generating the project). You use the function's body to set some config settings as well as persistent and local flags. You can also call `rootCmd.SetVersionTemplate` here to change the way the version prints, such as printing a short description of the application with the version:

```go
    versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
```