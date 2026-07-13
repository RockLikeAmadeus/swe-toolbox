# Get help

- SPC h -- opens Doom's central help hub
- SPC h b b -- Opens a searchable list of all active keybindings for your current file (in org mode, type `org` to filter them)
- SPC h m -- Describes the current major mode and lists every single shortcut
- SPC h k, followed by any shortcut -- "Describe Key"

# Dired

# Workspaces

Workspaces function like virtual desktops for isolating open buffers, window layouts, and project contexts. They do not persist across machine sessions by default, but can be configured to do so. You can also manually save individual workspace layouts to a file.

Workspaces-related functions are all under `SPC TAB`

- `SPACE TAB` - Bring up command list

- `SPC TAB n` - New workspace
- `SPC TAB N` - New named workspace
- `SPC TAB r` - Rename workspace

- `SPC TAB .` - View and select workspace

- `SPC TAB s` - Save workspace layout to file
- `SPC TAB l` - Load workspace layout from file

- `SPC q l` - From the Doom home screen to quick load last saved workspace

## Managing Scoped Buffers

- 'SPC b B' - Switch between all buffers
- `SPC ,` or `SPC b b` - Switch between buffers in the current workspace

# Evil Mode Editing

On a line item (`-`)
- `RET` - Convert to checkbox item (when not in org mode)
