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

### Manually downloading the binary (Windows, Mac and Linux)

* Go to the [latest release page](https://github.com/Geekgasm/brickdiff/releases/) and download the zip file for your platform.
* Unzip the archive
* Copy the executable (`brickdiff.exe` on Windows or `brickdiff` on Mac and Linux) to the desired location (if you have a directory for command-line tools that is included in your PATH, then this would be a good choice)
* In a cmd/Powershell/terminal window try executing it by typing `brickdiff.exe` (Windows) or `brickdiff` (Mac and Linux). It should print out a help message. (If you haven't copied the binary to a directory in your PATH you might need to specify the path where the binary is located.)

Note: you might have to allow the execution in your system settings. Refer to the general instructions for your platform.

## Usage

```
> ./brickdiff --help
A command line tool for comparing BrickLink Wanted lists

Usage:
  brickdiff [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  subtract    Subtracts two BrickLink Wanted lists in XML format and prints the output to the console

Flags:
  -h, --help      help for brickdiff
  -v, --version   Print version information

Use "brickdiff [command] --help" for more information about a command.
````

## Subtracting two wanted lists

```
> ./brickdiff subtract --help
Subtracts two BrickLink Wanted lists in XML format and prints the output to the console
        If an item is in the first list, but not in the second, it will be unchanged
        If an iten is in both lists, the quantity in the second list is subtracted from the quantity in the first,
        and if the result is >0 the item will be in the result with the remaining quantity.

Usage:
  brickdiff subtract bricklist1 bricklist2 [flags]

Flags:
  -c, --clipboard        Copy output to clipboard (default true)
  -h, --help             help for subtract
  -m, --mulitline        Multiline output (default is compact output)
  -o, --outfile string   Name of output file (default output is to clipboard)
  -s, --stdout           Print output to console (stdout)
```

# Example

Let's assume that you have the parts for my [Funny Birds: Chicken](https://rebrickable.com/mocs/MOC-71294/olivercgoetz/funny-birds-chicken/#details). 

And now you want to change it into the [Funny Birds: Rooster](https://rebrickable.com/mocs/MOC-83803/olivercgoetz/funny-birds-rooster/#details), which shares a lot of the same pieces with the chicken.

You want to know what pieces you would need in addition to the parts you already have for the chicken in order to build the rooster instead.

## Getting the parts lists

First step is to obtain the parts list:

* For a MOC you find on [Rebrickable](https://rebrickable.com) you can go to the MOC inventory, click on "Export Parts" and select "BrickLink XML". Copy the data from the popup window and save in into a file on your computer.
* For a wanted list you already have in your [BrickLink Account](https://www.bricklink.com/v2/wanted/list.page), select the desired wanted list and click on "Download" to save the XML file to your computer.
* For a MOC you have created in [BrickLink Studio](https://www.bricklink.com/v3/studio/download.page) use the "Add to Wanted List" feature to first upload the parts list as wanted list to your BrickLink account and then follow the steps in the previous bullet point.

In the `examples` folder of this repository, you will find the parts list for the chicken and the rooster, so that you can more easily try this out.

## Subtracting the parts list

If you want to know now what additional parts you need to build the rooster, you use the `brickdiff` tool to simply subtract the chicken list from the rooster list by typing the following command in a terminal/cmd window:

```
subtract examples/rooster.xml examples/chicken.xml
```

You should see an output like:

```
Output copied to clipboard.
```

## Uploading the parts list to BrickLink

To create the new wanted list with the result:

1. Go to the [BrickLink Wanted List Upload Page](https://www.bricklink.com/v2/wanted/upload.page?utm_content=subnav)
1. Click on "Upload BrickLink XML format"
1. Select a wanted list you want to add the parts to or select "Create New Wanted List" from the dropdown and provide a name for it.
1. Click on the area which says "Copy and paste here"
1. Paste the content of your clipboard (Command-C on Mac, Ctrl-C on Windows, or use the context menu)
1. Click on "Proceed to verify items". You should now see the list of parts with preview pictures.
1. Click on "Add to Wanted List"



