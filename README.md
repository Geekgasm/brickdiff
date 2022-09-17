# brickdiff

```
> ./brickdiff --help
A command line tool for comparing Bricklink Wanted lists

Usage:
  brickdiff [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  subtract    Subtracts two Bricklink Wanted lists in XML format and prints the output to the console

Flags:
  -h, --help     help for brickdiff
  -t, --toggle   Help message for toggle

Use "brickdiff [command] --help" for more information about a command.
````

## Subtracting two wanted lists

```
> ./brickdiff subtract --help
Subtracts two Bricklink Wanted lists in XML format and prints the output to the console
        If an item is in the first list, but not in the second, it will be unchanged
        If an iten is in both lists, the quantity in the second list is subtracted from the quantity in the first,
        and if the result is >0 the item will be in the result with the remaining quantity.

Usage:
  brickdiff subtract bricklist1 bricklist2 [flags]

Flags:
  -h, --help   help for subtract
```