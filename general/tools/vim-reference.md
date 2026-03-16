# Modes

Normal Mode (`Esc`)

Insert Mode (`i`, `a`, `o`)

Replace Mode (`r`, `R`)

Visual Mode (`v`, `V`, `Ctrl+v`)

Command Mode (`:`)

# Useful Combinations

## Exiting Vim

Discard changes and exit: `:q!`

Save changes: `:w`

Save changes and exit: `:wq`

## Move a line to a different location

(`dd`)elete line --> navigate to intended location --> (`p`)ut/(`p`)aste line under current line.

# Normal Mode

Some of these are operators, some are motions (which just enable cursor navigation when used without an operator), and some are combinations of operators, counts, and motions.

***Single Character***

(`a`)ppend word -- (`A`)ppend line -- (`i`)nsert -- (`r`)eplace character -- enter (`R`)eplace mode -- For(`w`)ord -- (`b`)ackword -- (`x`)-out -- (`u`)ndo -- (`U`)ndo line change

***Multi-Character (Operator Followed by Motion)***

(`d`)elete (`w`)ord -- (`d`)elete to (`e`)nd of word -- (`d`)elete re(`$`)t of line -- (`d`)elete the next (`3`) (`w`)ords --

(`c`)hange (`i`)nside (`(`)parenthesis

## Motions (Cursor Movement)

_Motions can be preceded with a **Count** (number) to be repeated that many times:_

Left: `h`

Right:  `l`

Up: `k`

Down: `j`

to start of next (`w`)ord

to (`e`)nd of current word

(`b`)ack one word

(`$`)hoot to end of line

(`0`)th character of the line 


# Operators (Verbs)

For many of these motions, typing them _twice_ (as in `dd`) will operate on the entire current line. _Capitalizing_ the operator often performs an action to the _end of the line_.

(`d`)elete

(`dd`)elete the whole line (can use `2dd` to delete the next two lines)

(`c`)hange

(`a`)ppend (or (`A`)ppend to end of line)

(`y`)ank or cop(`y`)

(`u`)ndo

(`U`)ndo line change

(`CTRL+R`)edo

(`o`)pen a new line below the cursor

--

Come back to Chapter 3 of the VIM tutor - between this point and the end of the tutor there are a few more commands that I want to add to this reference before I start using the Vim extension for text editing in VS Code for real, like (c)hange and (o)pen, etc.