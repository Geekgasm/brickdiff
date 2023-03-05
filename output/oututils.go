/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package output

import (
	"fmt"

	"github.com/spf13/pflag"
)

type OutputOptions struct {
	Multiline  bool
	Clipboard  bool
	Stdout     bool
	OutputFile string
}

func GetOutputOptions(flags *pflag.FlagSet) OutputOptions {
	multiline, _ := flags.GetBool("multiline")
	clipboard, _ := flags.GetBool("clipboard")
	stdout, _ := flags.GetBool("stdout")
	outputfile, _ := flags.GetString("outfile")
	if len(outputfile) == 0 && !stdout {
		clipboard = true
	}
	return OutputOptions{
		Multiline:  multiline,
		Clipboard:  clipboard,
		Stdout:     stdout,
		OutputFile: outputfile,
	}
}

func Output(output string, options OutputOptions) {
	if options.Stdout {
		fmt.Println(output)
	}
	if len(options.OutputFile) > 0 {
		WriteToFile(options.OutputFile, output)
	}
	if options.Clipboard {
		CopyToClipboard(output)
	}
}
