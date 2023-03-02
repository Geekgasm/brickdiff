/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

import (
	"testing"
)

func TestCreateItemKey(t *testing.T) {
	var someItem = Item{
		ItemType:  "type",
		ItemID:    "ID",
		Color:     "blue",
		Condition: "any",
		Maxprice:  "99",
		MinQTY:    3,
		Remarks:   "none",
		Notify:    "no",
	}
	key := getKey(someItem)
	if key.ItemID != someItem.ItemID {
		t.Fatalf(`key.ItemID was %v, expected %v`, key.ItemID, someItem.ItemID)
	}
	if key.ItemType != someItem.ItemType {
		t.Fatalf(`key.ItemType was %v, expected %v`, key.ItemType, someItem.ItemType)
	}
	if key.Color != someItem.Color {
		t.Fatalf(`key.Color was %v, expected %v`, key.Color, someItem.Color)
	}
}

func TestCreateItemMap(t *testing.T) {
	var item1 = Item{
		ItemType:  "type1",
		ItemID:    "ID1",
		Color:     "blue",
		Condition: "any",
		Maxprice:  "99",
		MinQTY:    3,
		Remarks:   "none",
		Notify:    "no",
	}
	var item2 = Item{
		ItemType:  "type2",
		ItemID:    "ID2",
		Color:     "red",
		Condition: "any",
		Maxprice:  "99",
		MinQTY:    3,
		Remarks:   "none",
		Notify:    "no",
	}
	itemMap := createItemMap(Inventory{Items: []Item{item1, item2}})

	if len(itemMap) != 2 {
		t.Fatalf(`Expected that item map has length 2, but it has length %v.`, len(itemMap))
	}
	assertMapEntry(t, itemMap, item1)
	assertMapEntry(t, itemMap, item2)
}

func assertMapEntry(t *testing.T, itemMap ItemMap, expectedItem Item) {
	item, found := itemMap[getKey(expectedItem)]
	if !found {
		t.Fatalf(`Expected item with key %v in map, but it wasn't found.`, getKey(expectedItem))
	}
	if item != expectedItem {
		t.Fatalf(`Expected item %v for key %v in map, but it found %v instead.`, expectedItem, getKey(expectedItem), item)
	}
}
