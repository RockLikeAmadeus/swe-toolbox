To create new project in current dir:

`$ cargo init`

To create new project in new dir:

`$ cargo new my-app`

Note: these commands create executable projects by default. Use `$ cargo new --lib` to bootstrap a library project, which may be easier to start with since it will include test scaffolding.

Additional Note: these commands _do_ generate a git repository automatically, unless one already exists. Change this by adding the `--vcs none` flag

## Starting a command line project

See [this](start-a-command-line-project.md).