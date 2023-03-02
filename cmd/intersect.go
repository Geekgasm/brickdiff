/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Geekgasm/brickdiff/bricklist"
	"github.com/Geekgasm/brickdiff/output"
	"github.com/spf13/cobra"
)

var intersectCmd = &cobra.Command{
	Use:     "intersect bricklist1 bricklist2",
	Aliases: []string{"int", "i"},
	Short:   "Creates the intersection of two BrickLink Wanted lists in XML format",
	Long: `Creates the intersection of two BrickLink Wanted lists in XML format and copies the output to the clipboard or writes it to the console or a file.
	If an item is in both lists, the item is added with the lower quantity to the result list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "Two arguments are required (filenames of the bricklink wanted lists)\n")
			os.Exit(1)
		}
		intersect(args[0], args[1], output.GetOutputOptions(cmd.Flags()))
	},
}

func init() {
	rootCmd.AddCommand(intersectCmd)
}

func intersect(bricklist1 string, bricklist2 string, outOptions output.OutputOptions) {
	inventory1 := bricklist.ReadXmlList(bricklist1)
	inventory2 := bricklist.ReadXmlList(bricklist2)
	result := bricklist.IntersectInventories(inventory1, inventory2)
	xmlString := bricklist.RenderXML(result, outOptions.Multiline)
	output.Output(xmlString, outOptions)
}
