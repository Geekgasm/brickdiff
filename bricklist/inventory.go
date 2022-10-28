/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

import (
	"fmt"
)

type Inventory struct {
	Items []Item `xml:"ITEM"`
}

type Item struct {
	ItemType  string `xml:"ITEMTYPE"`
	ItemID    string `xml:"ITEMID"`
	Color     string `xml:"COLOR"`
	Maxprice  string `xml:"MAXPRICE,omitempty"`
	MinQTY    int    `xml:"MINQTY"`
	Condition string `xml:"CONDITION,omitempty"`
	Remarks   string `xml:"REMARKS,omitempty"`
	Notify    string `xml:"NOTIFY,omitempty"`
}

type ItemKey struct {
	ItemType string
	ItemID   string
	Color    string
}

type ItemMap map[ItemKey]Item

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
		ItemType: item.ItemType,
		ItemID:   item.ItemID,
		Color:    item.Color,
	}
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
