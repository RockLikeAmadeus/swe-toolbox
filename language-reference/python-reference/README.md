# What makes Python unique?

What are a couple of the most important high-level things to be aware of that make this language what it is? What features and quirks define the language and must be kept in mind at (pretty much) all times?

# What do I need installed to write Python code

Rather than installing `pip` and then needing to install `virtualenv`, install [conda](https://anaconda.org/anaconda/conda), which includes both and simplifies things, even though it's a good deal more than just a Python package manager. Most likely you want to use [miniconda](https://docs.anaconda.com/miniconda/install/), as the full anaconda version is bloated with packages you likely won't use and can be added later if needed. Don't skip the "Verify your install" step. On Windows, Anaconda has a dedicated terminal "Anaconda Prompt".

From Conda, run `pip install ipython` for a more feature-full Python interpreter.

When it comes to which version of Python to use (2 or 3), just use Python 3 unless you're explicitly working on a legacy Python 2 codebase. That's the default behavior with Anaconda, anyway.

# I want to execute simple commands in the REPL. 

Run the interpreter with `$ python` or, for more features, `$ ipython`. If you're on Windows, you may need to do this from the Anaconda Prompt terminal.

To exit either, run `quit()`.

# I just want to write a quick script with <language>

Run scripts with `$ python main.py`

To parse command line arguments, use `sys.argv`. The first element is the script name, so start at 1:

```python
import sys

host = sys.argv[1]
port = sys.argv[2]
```

Or for more complex cases, use the [argparse](https://docs.python.org/3/library/argparse.html#module-argparse) module

# I'm starting a new software project with <language>

Python code is made up of packages which are made up of modules. Modules are just Python files, the things that get imported with `import`.

Packages are directories containing one or more `.py` files and a `__init__.py` file, which is often empty.

# Syntax and Style

## Essential Syntax and Style

Examples of the most important syntax and style to know without having to go all the way to external resources.
Also include the most notable syntax-related things to know about that make this language different than other languages.

## Comprehensive Syntax and Style

Links to relevant example syntax and style guides, like learnxinyminutes.com, or official style guides.

# Continue filling this out

[here](https://docs.python.org/3/tutorial/interpreter.html)