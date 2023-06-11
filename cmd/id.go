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

var idCmd = &cobra.Command{
	Use:   "id bricklist",
	Short: "Idenity function: output list will be the same as the input list",
	Long: `Parses the input list and outputs it again.
	Can be used to split up a list in conjunction with the --chunksize/-l option.
	Can be used to validate the format of the input list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintf(os.Stderr, "One argument is required (filename of the bricklink wanted list)\n")
			os.Exit(1)
		}
		id(args[0], output.GetOutputOptions(cmd.Flags()))
	},
}

func init() {
	rootCmd.AddCommand(idCmd)
}

func id(inputBricklist string, outOptions output.OutputOptions) {
	inventory := bricklist.ReadXmlList(inputBricklist)
	xmlString := bricklist.RenderXML(inventory, outOptions.Multiline, outOptions.ChunkSize)
	output.Output(xmlString, outOptions)
}
