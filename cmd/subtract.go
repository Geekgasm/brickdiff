/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Geekgasm/brickdiff/bricklist"
	"github.com/Geekgasm/brickdiff/output"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract bricklist1 bricklist2",
	Aliases: []string{"sub", "s"},
	Short:   "Subtracts two BrickLink Wanted lists in XML format",
	Long: `Subtracts two BrickLink Wanted lists in XML format and copies the output to the clipboard or writes it to the console or a file.
	If an item is in the first list, but not in the second, it will be unchanged
	If an item is in both lists, the quantity in the second list is subtracted from the quantity in the first,
	and if the result is >0 the item will be in the result with the remaining quantity.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "Two arguments are required (filenames of the bricklink wanted lists)\n")
			os.Exit(1)
		}
		multiline, _ := cmd.Flags().GetBool("mulitline")
		clipboard, _ := cmd.Flags().GetBool("clipboard")
		stdout, _ := cmd.Flags().GetBool("stdout")
		outputfile, _ := cmd.Flags().GetString("outfile")
		subtract(args[0], args[1], multiline, clipboard, stdout, outputfile)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}

func subtract(bricklist1 string, bricklist2 string, multiline bool, clipboard bool, stdout bool, outputfile string) {
	inventory1 := bricklist.ReadXmlList(bricklist1)
	inventory2 := bricklist.ReadXmlList(bricklist2)

	result := bricklist.SubtractInventories(inventory1, inventory2)

	xmlString := bricklist.RenderXML(result, multiline)

	if len(outputfile) > 0 {
		output.WriteToFile(outputfile, xmlString)
	} else if stdout {
		fmt.Println(xmlString)
	} else if clipboard {
		output.CopyToClipboard(xmlString)
	}
}
