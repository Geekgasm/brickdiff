/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package output

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type OutputOptions struct {
	Multiline  bool
	Clipboard  bool
	Stdout     bool
	OutputFile string
	ChunkSize  int
}

func GetOutputOptions(flags *pflag.FlagSet) OutputOptions {
	multiline, _ := flags.GetBool("multiline")
	clipboard, _ := flags.GetBool("clipboard")
	stdout, _ := flags.GetBool("stdout")
	outputfile, _ := flags.GetString("outfile")
	chunksize, _ := flags.GetInt("chunksize")
	if len(outputfile) == 0 && !stdout {
		clipboard = true
	}
	return OutputOptions{
		Multiline:  multiline,
		Clipboard:  clipboard,
		Stdout:     stdout,
		OutputFile: outputfile,
		ChunkSize:  chunksize,
	}
}

func Output(output []string, options OutputOptions) {
	if options.Stdout {
		fmt.Println(strings.Join(output, "\n"))
	}
	baseFilename := options.OutputFile
	fileExtension := ""
	if pos := strings.LastIndexByte(options.OutputFile, '.'); pos != -1 {
		baseFilename = options.OutputFile[:pos]
		fileExtension = options.OutputFile[pos:]
	}
	if len(options.OutputFile) > 0 {
		if len(output) == 0 {
			WriteToFile(options.OutputFile, output[0])
		} else {
			for i := 0; i < len(output); i++ {
				WriteToFile(fmt.Sprintf("%v-%v%v", baseFilename, i+1, fileExtension), output[i])
			}
		}
	}
	if options.Clipboard {
		CopyToClipboard(strings.Join(output, "\n"))
	}
}
