# Stroll

> One day, shortly after waking up, you experience an unexpected rush of ambition and motivation. *"It's such a nice day, maybe I should take a stroll around the neighborhood..."*
>
> After just a few moments of shoe-tying, notebook-grabbing, and knob-turning, you head out the door, ready to experience the world!

Stroll is a character-based, 2D, register-storage esoteric programming language about taking a quick stroll around the neighborhood.

**[Click here](using-the-stroll-tool.md) to view documentation on compiling and running the `stroll` binary, which can interpret valid Stroll files.**

## Taking a Stroll

A program written in Stroll is called a Stroll. It depicts the path you travel and the actions to take on your walk around the neighborhood. It is a text file composed of ASCII characters, arranged in a grid.

A stroll is composed of various nodes ("points"), which are connected by edges ("paths"). This effectively makes a stroll a graph. Each points and path may have an action associated with it, and the entire program is run incrementally until it terminates or crashes. No validation is done before the program arrives at the next action.

### Home

> A stroll always begins at your home! You walk out the door, which has always faced east, and begin your journey with a sense of wonder and excitement.

Your home is represented by an `H`, which can be placed anywhere. You **always** start walking to the east from your home.

### Waypoint

> After strolling for a bit, you see a sign up ahead. There's nothing special about the spot it marks, at least, not to you. After noting the sign, you continue on your journey, in the direction you best see fit.

The most basic point is a waypoint, at which no action occurs. From a waypoint, the path continues in any random direction that is not the one you came from. Usually, this is just one path, but it can also be used to travel on a random path.

### Skipping

> The sense of novelty you feel compels whimsical action. You find that you just can't help but skip along a large portion of your walk!

Skipping allows you to travel to another point. Skipping is noted with either a `-` (moving east/west) or a `|` (moving north/south)

We now have enough to assemble our first Stroll!
```
H-#
| |
#-#
```
In the above stroll, you leave your home heading east, then you head south, then west, then north back home. This program doesn't do anything at all, but it does run! An important note here is that **a stroll should ALWAYS return back home! If you don't return back home, you will be "lost" and the program will remind you to plan a better route next time!** While you don't *technically* have to return home for the program to function, it is considered a very good practice to do so.

### Walking and Taking Notes

> To keep yourself occupied, you decide to take note of the amount of steps you take between points. You note them down on 10 different notebook pages with tally marks. Each step you take walking forward is 1 single step, obviously. After each step you take walking *backwards* is then, by extension, -1 steps. Everyone knows that.

Instead of skipping everywhere on your stroll, you can just walk! Walking in a direction is noted with either a `^`, `v`, `>`, or `<` character.

Walking east, then, would look like this:
```
H>>>>>#
```
In this (invalid) Stroll, you walk 5 steps to the east. Note that the direction you walk in is *implicit* based on the direction you were travelling before. You can just as easily walk backwards 5 steps to the east like so:
```
H<<<<<#
```

Note: You cannot switch between walking forwards, walking backwards, and skipping mid-travel. You must stop at a point in order to change how you stroll. This is what waypoints are for.


---

As you stroll, you note the number of steps you take in your notebook. Each page of 10 holds a different count of steps. When you begin your walk, your notebook is flipped to page 0 (notebooks always start on page 0).

### Pages

> As you count your steps, you think it might be useful to separate your steps in some way. After all, the steps you took before this last waypoint aren't *really* the same steps as those you took after. Something about them just... *feels* different. You flip to another page and restart your count, while wondering what it all means.

You can flip to another page in your notebook by simply referencing that page at a point, such as `0`, `1`, `2`, ..., `9`. There are exactly 10 pages in your notebook, `0`-`9`. It's not a very thick notebook.

The following stroll demonstrates counting to 3 in page `0`, which we start our stroll on, and then counting to -3 on page `1`, before ending our stroll:
```
H>>>#
|   |
#>>>1
```

Note that an indentically functioning stroll can be written as follows:
```
#>>1<<#
^     ^
#--H--#
```

Strolls exist in many unique shapes and paths.

### Zero

> Shucks. You've lost count of your steps. Well, there's nowhere to go but forwards now! You hastily scribble out your old tallies, ready to start anew.

The zero point, `Z`, will reset the count on the current page to 0.

### Memorize and Recall

> You decide to take a mental note of your current step count. Focusing on a single number, you are able to persist it in your memory indefinitely. Who knows, at some point son on your stroll you may find it helpful to recall what it was!

You can memorize a number with `M` and later recall it with `R`. You can only memorize a single number at a time, and it will stay memorized until the stroll is completed. Memorizing a number can be very helpful for otherwise tedious tasks, such as swapping two numbers.

Swapping two numbers without memorize:
```
    1---s>1
    |   | ^
Y-H-r-9 l-2
|   v v |
3   2-# 9-e-r-3
|         v | |
Y         9<1 Z
|             v
2-Y-1<<<<<<<<<#
```

With memorize:
```
H--1-M-0-R-2-M-1-R-0-M-2
|                      |
Y<<<<<<<<<<Z-3-Y-2-Y-1-R
```
### Prepared Notes

> After grabbing a random notebook from the shelf, you notice that its pages already have a bunch of tallies on them! Oh well, you'll just have to continue from where they left off, I guess...

A stroll can accept user input. You can supply a single argument, a string, the characters of which will get pre-filled into pages `1`-`9` when you leave your home. The values of the characters as a Unicode decimal value will be used. They can then be changed or referenced like normal. Characters past the 9th in a supplied string are ignored.

### Yelling!

> This point feels different. You are overcome with energy! In this sudden surge, you find yourself unable to hold back from sharing your progress with the world! Since you know no one would understand the true meaning on your abstract tallies in your notes, you decide to take it upon yourself to translate your step count into something you *know* the world appreciates: Unicode.

The `Y` point represents a point to yell on your stroll. This point takes the current notebook page's value, and outputs the Unicode codepoint corresponding to that number.

### Cardinal Directions

> At the latest point, the sign is clear: though there are multiple paths forwards, one must only travel towards the direction indicated. You shrug this seemingly-artificial restriction off and continue on your merry way.

You can force a direction change by referencing cardinal directions (`n`, `s`, `e`, `w`). This *can* be used in place of waypoints, but is most useful when combining branching paths.

Consider the following stroll:
```
H>#>#>#
| | | |
#-w-w-#
```
After the stroll is completed, page `0` is always either 1, 2, or 3, and we always return home in a bounded amount of time.

### Crossings

> You stand before a bridge. You can see a path below you, crossing under the bridge, but it cannot be travelled upon from here. You hesistate in thought of where this path might go. Eventually, refocusing on your journey straight ahead, you continue onwards.

Crossings can be added with `+`. When you encounter a crossing, you must travel straight. You otherwise perform no action.

The following valid stroll prints out the input character, and is fully deterministic.
```
#---#
|   |
H-1-+-Y
    | |
    #-#
```

### Forks

> You've reached a fork in the road. You feel like you might have been here before...? That can't be right. Under normal circumstances you think maybe making a turn would be the right call, but something strikes you a bit differently this time... You decide to go straight!

Forks in the path will change directions conditionally. You can fork left or fork right, with `l` or `r` respectively. A fork in the road will turn in the direction indicated *unless the current notebook page's value is exactly 0.*

The following stroll will print the first character in the argument 10 times.
```
#--------------#
|              |
H-2>>>>>>>>>>0-r-2<#
               |   |
               2-Y-1
```

### Portals

> You happen upon a nexus of magical energy. It shimmers and coalesces before you, until eventually converging into an identifiable form... It appears to lead to an entire separate trail! With uncertainty and intrigue in your heart, you take a timid step into the radiant sphere, where your ultimate fate is not known...

Portals can be placed on a stroll with `@`. A portal can be entered from any direction, and will always transport you to the location of another portal on the stroll. If multiple portal other than the one you entered exist, one is chosen at random. The direction you enter a portal is not taken into account after being transported: you can always exit a portal on any valid path.

The following valid stroll will print the first character in the argument.
```
@-H-1-Y-@
```

### Comments

> In the distance you can barely make out a few shapes. They look bizarre to you. Strangely... alien. You decide it would be best not to loiter thinking about it for too long, and swiftly continue on your way.

Characters are checked for validity as they are encountered on the path, and any unknown character is interpreted as a `comment` node. Comments are always invalid to encounter, but if they are unencounterable on a stroll, they are therefore valid to include.

```
I'm the above program,

@-H-1-Y-@ but with included

comments that cannot be reached!
```

## Example Programs

### Hello World

The following stroll prints `hello, world` if you provide the following argument: `helo, wrd`.

```
H-1-Y-2-Y-3-Y-Y-4-Y-5-Y
|                     |
Y-9-Y-3-Y-8-Y-4-Y-7-Y-6
```

### Print 10 Times

As seen above, this stroll will print the first provided character 10 times.

```
#--------------#
|              |
H-2>>>>>>>>>>0-r-2<#
               |   |
               2-Y-1
```

### Random Bits

The following stroll will print either `0` or `1` *x* times, where *x* is the numeric value of the 1st character in the provided argument.

```
#-----#   y---#
|     |   ^   ^
H-2 0-r-1<2-Y-s
  v ^ |       |
  v ^ #-1-----#
  v ^
  v #<<<<<<<<<<<<<<<<<<<#
  v                     |
  #>>>>>>>>>>>>>>>>>>>>>#
```

### Nondeterminism

The following stroll represents non-deterministic behaviour. It has a 25% change of exiting each loop. Otherwise, it prints the first character.

```
  #-------#
  |       |
H-e-1-#-Y-n
|   |     |
|   #---Y-n
|   |
#---#
```

### Hello World, Goodbye Input

This stroll prints `Hello, World!\n` with no input needed. It is obviously not space-optimized.

```
#---------------------H>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y
|                                                                                              |
Y                               #<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<1-Y<<<<<<<<<<<<<<<<<<<<<<<<<<<<<#
v                               |
v                               #>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y--Y
v                                                                                                                 |
v Y<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<2
v |
v 3>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y<<<<<<<<<<<<Y>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y
v                                                                                                                   |
#>>>>>>>>>>>>>>>>Y>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y>>>>>>>>Y-1-Y<<<<<<<<<<<<<0-Y-2
```

### Hello World, Many Times

This stroll prints `Hello, World!\n` *x* times, where *x* is the numeric value of the first character in the argument.

```
                     #-#
                     | |
                     H-r---1<0>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y
                       |                                                                              |
Y-0-Z-2-Z-3-Z-4-Z-1----#        #<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<2--------Y<<<<<<<<<<<<<<<<<<<<<<<<<<<<<#
v                               |
v                               #>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y--Y
v                                                                                                                 |
v Y<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<3
v |
v 4>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y<<<<<<<<<<<<Y>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y
v                                                                                                                   |
#>>>>>>>>>>>>>>>>Y>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Y>>>>>>>>Y-2-Y<<<<<<<<<<<<<0-Y-3
```

### Fibonacci

The following program will print a Unicode character with a codepoint equal to the *n*th Fibonacci number, where *n* is the Unicode decimal value of the first character of the argument, minus `48`. 48 is subtracted to allow you to start with "0" corresponding to the number `0`, and so one.

One good example pairing is inputting a `=` (`U+003D`, or `61`, which corresponds to the `13`th Fibonacci number [`61` - `48` = `13`]). The expected output is `é` (`U+00E9`, or `233` a.k.a. "Latin Small Letter E with acute").

```
#----------------w-----w-#
|                |     | |
|                | #<9 | |
|                | | | | |
|                | 1<l-# |
|                |   |   |
|                #-# 0   |
|                  | |   |
9-Y-Z Y-H-1 #----r-# | #-#
    v ^   ^ v    ^   | |
# # v ^   ^ v    #-r># |
v ^ v ^   ^ v      |   |
v ^ v ^   ^ v      9>0-r-7-Z-8-r-----e-9-r-----e-#
v ^ v ^   ^ v          v       |     |   |     | |
v ^ #-#   ^ v          1       e-8-r-#   e-9-r-# |
v ^       ^ v          |       ^   ^     ^   ^   |
v ^       ^ v          |       7---#     8---#   |
v ^       ^ v          |                         |
v ^       ^ v          |    #--------------------#
v ^       ^ v          |    |
v ^       ^ v          |    7-r-----------e-8-r-----------e-#
v ^       ^ v          |      |           |   |           | |
v ^       ^ v          |      e-7-r-e-0-r-#   e-8-r-e-0-r-# |
v ^       ^ v          |      |   | |   |     |   | |   |   |
v ^       ^ v          |      |   0 #>0<7     |   0 #>0<8   |
v ^       ^ v          |      |   v           |   v         |
v ^       ^ v          |      #>7<9           #>8<9         |
v ^       ^ v          |                                    |
v ^       ^ v          #------------------------------------#
v ^       ^ v
v ^       ^ v
v ^       ^ v
v ^       ^ v
v ^       #-#
v ^
v ^
#-#
```

### Print Character Range

The following stroll will print all characters between *a* and *b*, where the argument provided is "ab". Providing an argument such that a > b will print a single newline.

| input | output                                                     |
| ----- | ---------------------------------------------------------- |
| ae    | abcde                                                      |
| 09    | 0123456789                                                 |
| Za    | Z[\]^_`a                                                   |
| Za    | ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz |

```
copy 1 to    copy 2 to   ensure 4 is      subtract 0   print starting at 8
0, 3, and 8  4 and 9     greater than 3   from 9       until 9 is zero
                                                                           
H-1-e-r------2-e-r-------e-4-r----# #-----0-e-r--------9>e-r-#
|   v |        v |       v   |    | |       v |          v | |
Y   1 0        2 4       3   3-r--+-#       0>9          9 | |
^   ^ v        | v       v     |  |                      ^ | |
^   3<8        #<9       4-----#  |                      Y-8 |
^                                 |                          |
#<<<<<<<Z-0-----------------------w--------------------------#

print newline
```