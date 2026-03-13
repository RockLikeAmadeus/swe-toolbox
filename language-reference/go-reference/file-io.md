# Generate a file that has fixed content, then inject dynamic data at runtime

Use Go's templates constructs. Depending on the type of file, you should either use [text templates](https://pkg.go.dev/text/template), or for HTML output, [HTML templates](https://pkg.go.dev/html/template).