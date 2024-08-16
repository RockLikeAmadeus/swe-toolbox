# Importing

From [Stack Overflow](https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go):

> You should have created your package with go mod init e.g. `$ go mod init github.com/my-org/my-package`

> Now in my-package you have a sub module called utils for example.

```
main.go
utils
 |- randstr.go
```

> And your randstr.go looks like this:

```go
package utils

func RandStr(n int) string {
    // TODO: Generate random string....
    return "I am a random string"
}
```

>And then anywhere in your project you would use exported (capitalized) functions from the utils package like this, for example in main.go:

```go
package main

import (
    "fmt"

    // "github.com/my-org/my-package" is the
    // module name at the top of your `go.mod`
    "github.com/my-org/my-package/utils"
)

func main() {
    fmt.Printf("Random string: %s\n", utils.RandStr(20))
}
```

You can reference the _exported_ types, variables, and functions from the package as `packagename.Something`. To make this shorter, you can [alias](https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/) the package name with a `.`, in your import statement, like this:

```go
import (
    . "math"
    "fmt"
)

fmt.Println(Pi)
```

# Organizing Command-Line Projects

[This](https://medium.com/swlh/how-to-structure-a-go-command-line-project-788c318a1d8c) blog post mentions the following typical structure (while stating that it didn't work in his case):

> One of the most common project structure recommendations is one that uses the following directory structure.
> - internal/app - Place the core application logic in this app package.
> - internal/pkg/ - Place any internal-only packages here (not app though).
> - pkg/ - Place any externally shareable packages here.
> - cmd/<app_name> - Place the main package here under a directory with the application name

