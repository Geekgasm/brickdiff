/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package output

import (
	"bufio"

	"fmt"
	"os"
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
