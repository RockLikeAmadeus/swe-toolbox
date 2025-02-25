- Come back to

  - https://jakevdp.github.io/WhirlwindTourOfPython/06-built-in-data-structures.html

- Need to pull advice out of https://docs.python-guide.org/writing/structure/ and links therein

  - imported with `import`
    - `from` and `as`

```
if __name__ == '__main__':
```

https://realpython.com/if-name-main-python/


---

All python variables should be thought of as pointers, rather than containers. Assignment simply changes what value the variable points to, but this only matters for mutable (i.e. non-primitive) types (and string is a primitive type).

Python values _do_ have types, it's just that those types are not tied to the variables that point to them.

Everything is an object, so even primitive types have attributes and methods that can be called on them.