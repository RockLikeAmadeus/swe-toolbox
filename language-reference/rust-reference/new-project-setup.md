To create new project in current dir:

`$ cargo init`

To create new project in new dir:

`$ cargo new my-app`

Note: these commands create executable projects by default. Use `$ cargo new --lib` to bootstrap a library project.

Additional Note: these commands _do_ generate a git repository automatically, unless one already exists. Change this by adding the `--vcs` flag