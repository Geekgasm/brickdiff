/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package util

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Inventory struct {
	Items []Item `xml:"ITEM"`
}

type Item struct {
	ItemType  string `xml:"ITEMTYPE"`
	ItemID    string `xml:"ITEMID"`
	Color     string `xml:"COLOR"`
	Maxprice  string `xml:"MAXPRICE"`
	MinQTY    int    `xml:"MINQTY"`
	Condition string `xml:"CONDITION"`
	Remarks   string `xml:"REMARKS"`
	Notify    string `xml:"NOTIFY"`
}

type ItemKey struct {
	ItemType  string
	ItemID    string
	Color     string
	Condition string
}

type ItemMap map[ItemKey]Item

func ReadList(list string) Inventory {
	xmlFile, err := os.Open(list)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %v: %v\n", list, err)
		os.Exit(2)
	}

	// fmt.Printf("Successfully Opened %v\n", list)
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing XML file %v: %v\n", list, err)
		os.Exit(3)
	}
	var inventory Inventory
	xml.Unmarshal(byteValue, &inventory)
	return inventory
}

func CreateSubtractionInventory(inventory1 Inventory, inventory2 Inventory) Inventory {
	inventory2Map := createItemMap(inventory2)
	resultInventory := Inventory{}
	for i := 0; i < len(inventory1.Items); i++ {
		item1 := inventory1.Items[i]
		item2, found := inventory2Map[getKey(item1)]
		if found {
			quantity1 := item1.MinQTY
			quantity2 := item2.MinQTY
			if quantity1 > quantity2 {
				resultItem := item1
				resultItem.MinQTY = quantity1 - quantity2
				resultInventory.Items = append(resultInventory.Items, resultItem)
			}
			// otherwise the result is 0 or negative, so item is not added
		} else {
			// item not in inventory2 -> just add the original item1
			resultInventory.Items = append(resultInventory.Items, item1)
		}
	}
	return resultInventory
}

func createItemMap(inventory Inventory) ItemMap {
	itemMap := ItemMap{}
	for i := 0; i < len(inventory.Items); i++ {
		item := inventory.Items[i]
		itemMap[getKey(item)] = item
	}
	return itemMap
}

func getKey(item Item) ItemKey {
	return ItemKey{
		ItemType:  item.ItemType,
		ItemID:    item.ItemID,
		Color:     item.Color,
		Condition: item.Condition,
	}
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

func PrintInventory(inventory Inventory) {
	for i := 0; i < len(inventory.Items); i++ {
		fmt.Println("ITEMTYPE: " + inventory.Items[i].ItemType)
		fmt.Println("ITEMID: " + inventory.Items[i].ItemID)
		fmt.Println("COLOR: " + inventory.Items[i].Color)
		fmt.Println("MAXPRICE: " + inventory.Items[i].Maxprice)
		fmt.Printf("MINQTY: %v\n", inventory.Items[i].MinQTY)
		fmt.Println("CONDITION: " + inventory.Items[i].Condition)
		fmt.Println("REMARKS: " + inventory.Items[i].Remarks)
		fmt.Println("NOTIFY: " + inventory.Items[i].Notify)
		fmt.Println()
	}
}
