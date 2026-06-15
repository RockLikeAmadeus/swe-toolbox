# Helpful commands

## On a feature branch, reset the state of a specific file to it's state from the source branch

```bash
git restore --source <source_branch> -- <path/to/file>
```

e.g.

```bash
git restore --source main -- src/app.js
```