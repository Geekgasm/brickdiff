/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Geekgasm/brickdiff/build"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "brickdiff",
	Short:   "A command line tool for comparing BrickLink Wanted lists",
	Version: fmt.Sprintf("%v\ncommit:   %v\nbuilt at: %v", build.Version, build.Commit, build.Date),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.brickdiff.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
	rootCmd.PersistentFlags().BoolP("clipboard", "c", false, "Copy output to clipboard (default)")
	rootCmd.PersistentFlags().BoolP("multiline", "m", false, "Multiline output (default is compact output)")
	rootCmd.PersistentFlags().BoolP("stdout", "s", false, "Print output to console (stdout)")
	rootCmd.PersistentFlags().StringP("outfile", "o", "", "Name of output file")
	rootCmd.PersistentFlags().IntP("chunksize", "l", 0, "Maximum chunk size limit for the output files; longer lists will be split into several XML documents")
}
