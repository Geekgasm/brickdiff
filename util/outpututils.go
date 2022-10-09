/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package util

import (
	"bufio"

	"fmt"
	"os"

	"golang.design/x/clipboard"
)

func WriteToFile(outputfile string, output string) {
	var f *os.File

	os.Remove(outputfile)
	f, err := os.OpenFile(outputfile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error opening output file %v: %v\n", outputfile, err)
		os.Exit(5)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(output)
	if err != nil {
		fmt.Printf("Error writing to output file %v: %v\n", outputfile, err)
		os.Exit(6)
	}
	writer.Flush()
	fmt.Printf("Output written to file %v\n", outputfile)
}

func CopyToClipboard(output string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Warning: clipboard could not be initialized, XML is not copied to clipboard.\n")
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(output))
	fmt.Printf("Output copied to clipboard.\n")
}
