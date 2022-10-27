/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadXmlList(filename string) Inventory {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %v: %v\n", filename, err)
		os.Exit(2)
	}

	// fmt.Printf("Successfully Opened %v\n", list)
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing XML file %v: %v\n", filename, err)
		os.Exit(3)
	}
	var inventory Inventory
	xml.Unmarshal(byteValue, &inventory)
	return inventory
}

func RenderXML(inventory Inventory, multiline bool) string {
	tmp := struct {
		Inventory
		XMLName struct{} `xml:"INVENTORY"`
	}{Inventory: inventory}

	indentation := ""
	if multiline {
		indentation = "  "
	}
	b, err := xml.MarshalIndent(tmp, "", indentation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting result inventory to XML: %v\n", err)
		os.Exit(4)
	}
	xmlString := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	if multiline {
		xmlString += "\n"
	}
	xmlString += string(b)
	xmlString += "\n"
	return xmlString
}
