# Modes

Normal Mode (`Esc`)
Insert Mode (`i`, `a`, `o`)
Replace Mode (`r`, `R`)
Visual Mode (`v`, `V`, `Ctrl+v`)
Command Mode (`:`)

# Critical Combinations

## Exiting Vim

Discard changes and exit: `:q!`

Save changes: `:w`

Save changes and exit: `:wq`

# Normal Mode

Some of these are operators, some are motions (which just enable cursor navigation when used without an operator), and some are combinations of operators, counts, and motions.

***Single Character***

(a)ppend word -- (A)ppend line -- (i)nsert -- (r)eplace character -- enter (R)eplace mode -- For(w)ord -- (b)ackword -- (x)-out -- (u)ndo -- (U)ndo line change

***Multi-Character (Operator Followed by Motion)***

(d)elete (w)ord -- (d)elete to (e)nd of word -- (d)elete re($)t of line -- (d)elete the next (3) (w)ords --

## Cursor Movement

Left: `h`

Right:  `l`

Up: `k`

Down: `j`


# Operators

(d)elete

(dd)elete the whole line (can use `2dd` to delete the next two lines)

(u)ndo

(U)ndo line change

(CTRL+R)edo



# Motions

_Motions can be preceded with a **Count** (number) to be repeated that many times:_

to start of next (w)ord

to (e)nd of current word

(b)ack one word

($)hoot to end of line

(0)th character of the line (technically not a motion I think)