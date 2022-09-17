/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Geekgasm/brickdiff/util"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:   "subtract bricklist1 bricklist2",
	Short: "Subtracts two Bricklink Wanted lists in XML format and prints the output to the console",
	Long: `Subtracts two Bricklink Wanted lists in XML format and prints the output to the console
	If an item is in the first list, but not in the second, it will be unchanged
	If an iten is in both lists, the quantity in the second list is subtracted from the quantity in the first,
	and if the result is >0 the item will be in the result with the remaining quantity.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "Two arguments are required (filenames of the bricklink wanted lists)\n")
			os.Exit(1)
		}
		subtract(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}

func subtract(bricklist1 string, bricklist2 string) {
	// fmt.Printf("Opening file %v\n", bricklist1)
	inventory1 := util.ReadList(bricklist1)
	// printInventory(inventory1)

	// fmt.Printf("Opening file %v\n", bricklist2)
	inventory2 := util.ReadList(bricklist2)
	// printInventory(inventory2)

	result := util.CreateSubtractionInventory(inventory1, inventory2)
	// fmt.Printf("RESULT\n")
	// printInventory(result)

	util.PrintInventoryXML(result)
}
