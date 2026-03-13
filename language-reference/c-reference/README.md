
# What makes C unique?

_What are a couple of the most important high-level things to be aware of that make this language what it is? What features and quirks define the language and must be kept in mind at (pretty much) all times? Note: this section is not for notes on unique syntax--that section is below. Delete this block once the reference for this language is "complete"_

C is the lowest-level language you can write before assembly, and knowing it will teach you a great deal about how the machine works at a more fundamental level. According to Beej, the only thing that really stands out for people when it comes to C are pointers and manual memory management, but I suspect this isn't quite the whole story.

Of course, C doesn't have fancy string functionality, and if your application relies heavily on string parsing and manipulation you should use a different language. Strings in C have the type `char *`, to be read as "char pointer".

A few more things:
- Never ignore compiler warnings.
- Variables need to be initialized explicitly or they can contain nonsense data.
- Functions must be defined before they're used--or, more commonly, a function _prototype_ must be declared before the function is used.
- You should always use `void` between parens to indicate a function has no parameters, in the function definition and _especially_ in the prototype.
- _Everything_ in C is pass by value, that is, a copy is put into the parameter. Even pointers are copies of pointers.
- Arrays and pointers are pretty much the same thing.
- Explicit casting is sometimes necessary, but rarely. If you find yourself doing it, ask if you really need to. See Chapter 15 of [Beej's Guide](https://beej.us/guide/bgc/html/split/types-iii-conversions.html#types-iii-conversions)


[contents]

# What do I need installed to write C code

_A couple quick notes on what should be installed, i.e. compilers, interpreters, package managers, common workflow tools, etc... Include notes on specific language/compiler versions if appropriate. Delete this block once the reference for this language is "complete"_

You'll need a C compiler, which may come installed with your OS if it's not Windows. Try running `$ cc` or `$ gcc` or `$ clang` (Mac) from the command line. For Mac specifically, you may want to install XCode. For Windows, install WSL for gcc, or use an IDE. VS Code can easily run C programs with a debugger if you install the C/C++ extension.

# I just want to write a quick script with C

_Notes on writing and then running a standalone script or small process here. Include the simplest way to retrieve command line arguments. Delete this block once the reference for this language is "complete"_

[contents]

# I'm starting a new software project with C

_Notes on the structure and terminology of more complex, multi-file projects here. Talk about how code is organized in <language>, how files are named, common patterns for specific types of projects, and tools commonly used for scaffolding new projects. Keep it brief, and leave the details to [this page I haven't created yet]. Delete this block once the reference for this language is "complete"_

Refer to [Beej's Guide](https://beej.us/guide/bgc/html/split/multifile-projects.html#multifile-projects) for multi-file projects. But you should probably utilize the `make` utility in most cases.

[contents]

# Syntax and Style

## Essential Syntax and Style

_Examples of the most important syntax and style to know without having to go all the way to external resources.
Also include the most notable syntax-related things to know about that make this language different than other languages. Delete this block once the reference for this language is "complete"_

First use [Learn C in Y Minutes]. If the answer isn't there, look up the relevant chapter in [Beej's Guide](https://beej.us/guide/bgc/html/split/), which is really well organized and sufficiently concise.

[contents]

## Comprehensive Syntax and Style

[Essential Syntax](https://learnxinyminutes.com/c/)

[contents]



## Basics of Testing

_What are the basics of automated tests for this language? If testing isn't built into the language, what testing frameworks are most popular? What is the basic syntax for writing automated tests, and how are test functions and files organized?_

## Essential Language Toolset Commands

_What commands with the languages compiler and/or included command-line tools should be known, or at least be known of?_

## Tools

_What are some well-known and widely-used third party tools and frameworks that it's important to be aware of, such as CLI or TUI frameworks, or web frameworks? Be sure to include what they're used for and maybe some well established pros, cons, and use cases._

## Utilities

_What are some well-known and widely-used third party libraries that it's important to be aware of, such as logging, financial libraries, math libraries, or data structure utilities? Be sure to include what they're used for and maybe some well established pros, cons, and use cases._