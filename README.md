# brickdiff

Have you ever desinged a LEGO MOC with BrickLink Studio, created the wanted list to order the parts and sourced the parts, only to decide to make some changes afterwards in your digital model? How do you check what parts you now need in addition to build your MOC with the new version? That's what `brickdiff` is for: it allows you to "subtract" a wanted list from another one, so that you end up with a new wanted list that only contains the parts you need to make the changes.

## Installation

### With homebrew (Mac and Linux)

This is the easiest way to install `brickdiff` if you have already have `homebrew` installed on your Mac or Linux system.

Open a terminal and type:

```
brew tap geekgasm/homebrew-tap
brew install brickdiff
```

Note: you might get asked for a github password, but you don't actually need to provide one.

To upgrade to the latest released version later, type 

```
brew upgrade brickdiff
```

### Manually downloading the binary (Windows, Mac and Linux)

* Go to the [latest release page](https://github.com/Geekgasm/brickdiff/releases/) and download the zip file for your platform.
* Unzip the archive
* Copy the executable (`brickdiff.exe` on Windows or `brickdiff` on Mac and Linux) to the desired location (if you have a directory for command-line tools that is included in your PATH, then this would be a good choice)
* In a cmd/Powershell/terminal window try executing it by typing `brickdiff.exe` (Windows) or `brickdiff` (Mac and Linux). It should print out a help message. (If you haven't copied the binary to a directory in your PATH you might need to specify the path where the binary is located.)

Note: you might have to allow the execution in your system settings. Refer to the general instructions for your platform.

To upgrade to the latest released version later, follow the same steps and simply replace the old version with the new one.

## Usage

```
> brickdiff --help
A command line tool for comparing BrickLink Wanted lists

Usage:
  brickdiff [command]

Available Commands:
  add         Adds two BrickLink Wanted lists in XML format
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  id          Idenity function: output list will be the same as the input list
  intersect   Creates the intersection of two BrickLink Wanted lists in XML format
  multiply    Multiplies the quantity of all parts in BrickLink Wanted list with a given factor
  subtract    Subtracts two BrickLink Wanted lists in XML format
  union       Creates the union of two BrickLink Wanted lists in XML format

Flags:
  -l, --chunksize int    Maximum chunk size limit for the output files. Longer lists will be split into several files. Can not be combined with the clipboard option
  -c, --clipboard        Copy output to clipboard (default)
  -h, --help             help for brickdiff
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
  -v, --version          Print version information

Use "brickdiff [command] --help" for more information about a command.
````

## Subtracting two wanted lists

Subtracting two wanted lists is useful if you have modified your MOC digitally after already building a physical version. It allows you to easily find out which parts you need to make the changes on your physical build.

```
> brickdiff subtract --help
Subtracts two BrickLink Wanted lists in XML format and copies the output to the clipboard or writes it to the console or a file.
        If an item is in the first list, but not in the second, it will be unchanged
        If an item is in both lists, the quantity in the second list is subtracted from the quantity in the first,
        and if the result is >0 the item will be in the result with the remaining quantity.

Usage:
  brickdiff subtract bricklist1 bricklist2 [flags]

Aliases:
  subtract, sub, s

Flags:
  -h, --help   help for subtract

Global Flags:
  -c, --clipboard        Copy output to clipboard (default)
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
```

## Adding two wanted lists

Adding two wanted is useful if you have digitally created parts of your MOC in separate files and now want to have a complete list for all the parts from the separate builds.

```
> brickdiff add --help
Adds two BrickLink Wanted lists in XML format and copies the output to the clipboard or writes it to the console or a file.
        If an item is in both lists, the quantities are added up in the result list.
        If an item is in only one list, the item will be added to the reault list.

Usage:
  brickdiff add bricklist1 bricklist2 [flags]

Aliases:
  add, a

Flags:
  -h, --help   help for add

Global Flags:
  -c, --clipboard        Copy output to clipboard (default)
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
```

## Union of two wanted lists

Building the union of two lists can be useful if you have two variations of a MOC and want to create a list for a 2-in-1 model. With the parts of the list either of the MOCs could be built, but not both at the same time.

```
> brickdiff union --help
Creates the union of two BrickLink Wanted lists in XML format and copies the output to the clipboard or writes it to the console or a file.
        If an item is in both lists, the item is added with the higher quantity to the result list.
        If an item is in only one list, the item will be added to the result list.

Usage:
  brickdiff union bricklist1 bricklist2 [flags]

Aliases:
  union, u

Flags:
  -h, --help   help for union

Global Flags:
  -c, --clipboard        Copy output to clipboard (default)
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
```

## Multiplying a wanted list with a factor

Imagine you have downloaded a wanted list for a MOC, and you want to order parts for multiples. This function allows you to multiply the quaontities of all the parts of a given wanted list with a given factor.

```
> brickdiff multiply --help
Multiplies the quantity of all parts in a BrickLink Wanted list with a given positive integer number and copies the output to the clipboard or writes it to the console or a file.

Usage:
  brickdiff multiply bricklist factor [flags]

Aliases:
  multiply, mult, m

Flags:
  -h, --help   help for multiply

Global Flags:
  -c, --clipboard        Copy output to clipboard (default)
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
```

## Intersection of two wanted lists

Intersecting two lists will result in a list that only contains parts that are present in both lists with the lower quantity for the part.

```
> brickdiff intersect --help
Creates the intersection of two BrickLink Wanted lists in XML format and copies the output to the clipboard or writes it to the console or a file.
        If an item is in both lists, the item is added with the lower quantity to the result list.

Usage:
  brickdiff intersect bricklist1 bricklist2 [flags]

Aliases:
  intersect, int, i

Flags:
  -h, --help   help for intersect

Global Flags:
  -c, --clipboard        Copy output to clipboard (default)
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
```

## Identity function

To allow for splitting a list without performing any transformation function, the `id` function was added. But it also can serve to just validate an XML file.

```
> brickdiff id --help
Parses the input list and outputs it again.
        Can be used to split up a list in conjunction with the --chunksize/-l option.
        Can be used to validate the format of the input list.

Usage:
  brickdiff id bricklist [flags]

Flags:
  -h, --help   help for id

Global Flags:
  -l, --chunksize int    Maximum chunk size limit for the output files. Longer lists will be split into several files. Can not be combined with the clipboard option
  -c, --clipboard        Copy output to clipboard (default)
  -m, --multiline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file
  -s, --stdout           Print output to console (stdout)
```

# Examples

Let's assume that you have the parts for my [Funny Birds: Chicken](https://rebrickable.com/mocs/MOC-71294/olivercgoetz/funny-birds-chicken/#details). 

And now you want to change it into the [Funny Birds: Rooster](https://rebrickable.com/mocs/MOC-83803/olivercgoetz/funny-birds-rooster/#details), which shares a lot of the same pieces with the chicken.

You want to know what pieces you would need in addition to the parts you already have for the chicken in order to build the rooster instead.

## Getting the parts lists

First step is to obtain the parts list:

* For a MOC you find on [Rebrickable](https://rebrickable.com) you can go to the MOC inventory, click on _Export Parts_ and select _BrickLink XML_. Copy the data from the popup window and save in into a file on your computer.
* For a wanted list you already have in your [BrickLink Account](https://www.bricklink.com/v2/wanted/list.page), select the desired wanted list and click on _Download_ to save the XML file to your computer.
* For a MOC you have created in [BrickLink Studio](https://www.bricklink.com/v3/studio/download.page) select _File_ ➔ _Export As_ ➔ _Export as WL xml..._, select a name and location, and save the file.

In the `examples` folder of this repository, you will find the parts list for the chicken and the rooster, so that you can more easily try this out.

## Subtracting the parts list

If you want to know now what additional parts you need to build the rooster, you use the `brickdiff` tool to simply subtract the chicken list from the rooster list by typing the following command in a terminal/cmd window:

```
brickdiff subtract examples/rooster.xml examples/chicken.xml
```
(Note: on Windows use the backslash "\\" instead of the forward slash "/" to separate the folder names)

You should see an output like:

```
Output copied to clipboard.
```

## Splitting a list in smaller chunks

Sometimes it's useful to split up a long list into smaller chunks. For example, when ordering from the LEGO "Pick a Brick" service, there is a limit of how many lots can be added in one order. To that end, there is an output option `--chunkzize` that limits the number of lots to include in each output XML document. When using in conjunction with the `--outfile` output option, this will result in several files with a trailing number (starting with 1) in case the result list is longer than the chunksize.

```
> brickdiff id examples/rooster.xml --chunksize 10 --outfile rooster-chunk.xml
Output written to file rooster-chunk-1.xml
Output written to file rooster-chunk-2.xml
Output written to file rooster-chunk-3.xml
```

This will create three XML files with 10, 10 and 1 lot (the original file has 21 lots):
```
> ls -l rooster-chunk-* 
-rw-r--r--  1 olivergotz  staff  976 Jun 11 14:16 rooster-chunk-1.xml
-rw-r--r--  1 olivergotz  staff  969 Jun 11 14:16 rooster-chunk-2.xml
-rw-r--r--  1 olivergotz  staff  158 Jun 11 14:16 rooster-chunk-3.xml
```

## Uploading the parts list to BrickLink

To create the new wanted list on Bricklink with the result from `brickdiff`:

1. Go to the [BrickLink Wanted List Upload Page](https://www.bricklink.com/v2/wanted/upload.page?utm_content=subnav)
1. Click on _Upload BrickLink XML format_
1. Select a wanted list you want to add the parts to or select _Create New Wanted List_ from the dropdown and provide a name for it.
1. Click on the area which says _Copy and paste here_
1. Paste the content of your clipboard (Command-C on Mac, Ctrl-C on Windows, or use the context menu)
1. Click on _Proceed to verify items_. You should now see the list of parts with preview pictures.
1. Click on _Add to Wanted List_

## Ordering parts from LEGO Pick a Brick

The XML files cannot be directly uploaded to the _LEGO Pick a Brick_ service for ordering single parts, as this service is designed to be used interactively in a web browser.

Fortunately, other talented LEGO fans have developed a browser plugin [BrickHunter](https://github.com/BrickTwo/BrickHunter) which can import brick lists in different formats, including the Bricklink XML format. BrickHunter is an open source project, and you can install the plugin for Google Chrome, Microsoft Edge and Mozilla Firefox browsers from their respective extensions store.

Note that you might need to split up longer wanted lists, as there are limitations how many lots can be ordered at once from LEGO. The BrickHunter pluging will just report a warning that the list was not transferred to the LEGO shopping cart.


