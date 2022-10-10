# brickdiff

Have you ever desinged a LEGO MOC with Bricklink Studio, created the wanted list to order the parts and sourced the parts, only to decide to make some changes afterwards in your digital model? How do you check what parts you now need in addition to build your MOC with the new version? That's what `brickdiff` is for: it allows you to "subtract" a wanted list from another one, so that you end up with a new wanted list that only contains the parts you need to make the changes.

## Installation

### With homebrew (Mac and Linux)

This is the easiest way to install `brickdiff` if you have already have `homebrew` installed on your Mac or Linux system.

Open a terminal and type:

```
brew tap geekgasm/homebrew-tap
brew install brickdiff
```

Note: you might get asked for a github password, but you don't actually need to provide one.

### Manually downloading the binary (Windows, Mac and Linux)

* Go to the [latest release page](https://github.com/Geekgasm/brickdiff/releases/) and download the zip file for your platform.
* Unzip the archive
* Copy the executable (`brickdiff.exe` on Windows or `brickdiff` on Mac and Linux) to the desired location (if you have a directory for command-line tools that is included in your PATH, then this would be a good choice)
* In a cmd/Powershell/terminal window try executing it by typing `brickdiff.exe` (Windows) or `brickdiff` (Mac and Linux). It should print out a help message. (If you haven't copied the binary to a directory in your PATH you might need to specify the path where the binary is located.)

Note: you might have to allow the execution in your system settings. Refer to the general instructions for your platform.

## Usage

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
  -c, --clipboard   Copy output to clipboard (default true)
  -h, --help        help for subtract
  -m, --mulitline   Multiline output (default is compact output)
      --o string    Name of output file (default output is to clipboard)
  -s, --stdout      Print output to console (stdout)
```