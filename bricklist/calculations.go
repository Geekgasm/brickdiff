/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

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
