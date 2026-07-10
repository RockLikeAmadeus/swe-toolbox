First make sure `.cobra.yaml` file exists in `$HOME` directory with contents:

```yaml
author: Alec Hampton
license: Apache
```

Then:

```bash
$ mkdir myapp # contains everything
$ cd myapp
$ go mod init github.com/RockLikeAmadeus/myapp
$ touch myapp.go # *must insert `package myapp` at top of file*, contains just business logic
$ go get -u github.com/spf13/cobra@latest
$ mkdir myappcli # contains just the CLI interface, name of this folder = name of exe
$ cd myappcli
$ cobra-cli init
```

If the `cobra-cli` command fails, may need to run `$ go install github.com/spf13/cobra-cli@latest`.

Now use `cobra-cli` by itself for command list.

Inside `myappcli` run with `$ go run .`, build with `$ go build`, and install to `$GOPATH` with `$ go install`, after which you can just run `myappcli` from terminal.