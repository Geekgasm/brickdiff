/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package output

import (
	"fmt"

	"golang.design/x/clipboard"
)

func CopyToClipboard(output string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Warning: clipboard could not be initialized, XML is not copied to clipboard.\n")
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(output))
	fmt.Printf("Output copied to clipboard.\n")
}
