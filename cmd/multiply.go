/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Geekgasm/brickdiff/bricklist"
	"github.com/Geekgasm/brickdiff/output"
	"github.com/spf13/cobra"
)

var multiplyCmd = &cobra.Command{
	Use:     "multiply bricklist factor",
	Aliases: []string{"mult", "m"},
	Short:   "Multiplies the quantity of all parts in BrickLink Wanted list with a given factor",
	Long:    `Multiplies the quantity of all parts in a BrickLink Wanted list with a given positive integer number and copies the output to the clipboard or writes it to the console or a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "Two arguments are required (filename of the bricklink wanted lists and a factor)\n")
			os.Exit(1)
		}
		factor, err := strconv.Atoi(args[1])
		if err != nil || factor < 1 {
			fmt.Fprintf(os.Stderr, "Factor argument needs to be a positive integer number.\n")
			os.Exit(1)
		}
		multiline, _ := cmd.Flags().GetBool("mulitline")
		clipboard, _ := cmd.Flags().GetBool("clipboard")
		stdout, _ := cmd.Flags().GetBool("stdout")
		outputfile, _ := cmd.Flags().GetString("outfile")
		multiply(args[0], factor, multiline, clipboard, stdout, outputfile)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
}

func multiply(bricklinklist string, factor int, multiline bool, clipboard bool, stdout bool, outputfile string) {
	inventory := bricklist.ReadXmlList(bricklinklist)

	result := bricklist.Multiply(inventory, factor)

	xmlString := bricklist.RenderXML(result, multiline)

	if len(outputfile) > 0 {
		output.WriteToFile(outputfile, xmlString)
	} else if stdout {
		fmt.Println(xmlString)
	} else if clipboard {
		output.CopyToClipboard(xmlString)
	}
}
