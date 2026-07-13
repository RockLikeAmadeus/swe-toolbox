# General

- `TAB` - Toggle folding for the current heading subtree
- `Shift+TAB` Cycle global visibility states for the entire document

- `M`
- `C-RET` -- Create a new line of the same type as previous
- `M-RET` -- Create a new heading at current level
~~- `C-RET` -- Create a new heading at current level and enter INSERT mode~~
- `M-h` and `M-l` -- Promotes or demotes the current heading level
- `M-j` and `M-k` -- Moves the current heading or subtree up or down in the file
- `g j` and `g k` -- Moves the cursor to the next or previous visible heading at the same level

- `m i` -- Toggles the current plain text line to a list item, or the current task to a checkbox item

# Formatting

- Highlight text in Visual Mode (v) and use the Surround (`S`) method
  - `S *` -- Bold
  - `S /` -- Italic
  - `S ~` -- Code style
  - `S =` -- Verbatim
  - `S _` -- Underline
  - `S +` -- Strikethrough

# Task Management
- `SPC m t` -- Opens the menu to change task type or status, I guess
- `SPC m t t` -- Turns the heading into a task
- `SPC m x` - Toggle checkbox on current line (doesn't need to be a task)
- `RET` on a task -- Cycle task state
- If you `C-RET` on a task line, it will add another task

