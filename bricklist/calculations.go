/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

func SubtractInventories(inventory1 Inventory, inventory2 Inventory) Inventory {
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

func AddInventories(inventory1 Inventory, inventory2 Inventory) Inventory {
	inventory2Map := createItemMap(inventory2)
	resultInventory := Inventory{}
	for i := 0; i < len(inventory1.Items); i++ {
		item1 := inventory1.Items[i]
		item2, found := inventory2Map[getKey(item1)]
		if found {
			// item is in both inventories, add item with added quantities
			resultItem := item1
			resultItem.MinQTY = item1.MinQTY + item2.MinQTY
			resultInventory.Items = append(resultInventory.Items, resultItem)
		} else {
			// item not in inventory2 -> just add the original item1
			resultInventory.Items = append(resultInventory.Items, item1)
		}
	}
	inventory1Map := createItemMap(inventory1)
	// Now add all items from inventory2 that were not in inventory1
	for i := 0; i < len(inventory2.Items); i++ {
		itemToAdd := inventory2.Items[i]
		_, found := inventory1Map[getKey(itemToAdd)]
		if !found {
			resultInventory.Items = append(resultInventory.Items, itemToAdd)
		}
	}
	return resultInventory
}

func UnionInventories(inventory1 Inventory, inventory2 Inventory) Inventory {
	inventory2Map := createItemMap(inventory2)
	resultInventory := Inventory{}
	for i := 0; i < len(inventory1.Items); i++ {
		item1 := inventory1.Items[i]
		item2, found := inventory2Map[getKey(item1)]
		if found {
			// item is in both inventories, add item with maximum quantity
			resultItem := item1
			resultItem.MinQTY = max(item1.MinQTY, item2.MinQTY)
			resultInventory.Items = append(resultInventory.Items, resultItem)
		} else {
			// item not in inventory2 -> just add the original item1
			resultInventory.Items = append(resultInventory.Items, item1)
		}
	}
	inventory1Map := createItemMap(inventory1)
	// Now add all items from inventory2 that were not in inventory1
	for i := 0; i < len(inventory2.Items); i++ {
		itemToAdd := inventory2.Items[i]
		_, found := inventory1Map[getKey(itemToAdd)]
		if !found {
			resultInventory.Items = append(resultInventory.Items, itemToAdd)
		}
	}
	return resultInventory
}

func IntersectInventories(inventory1 Inventory, inventory2 Inventory) Inventory {
	inventory2Map := createItemMap(inventory2)
	resultInventory := Inventory{}
	for i := 0; i < len(inventory1.Items); i++ {
		item1 := inventory1.Items[i]
		item2, found := inventory2Map[getKey(item1)]
		if found {
			resultItem := item1
			resultItem.MinQTY = min(item1.MinQTY, item2.MinQTY)
			resultInventory.Items = append(resultInventory.Items, resultItem)
		}
	}
	return resultInventory
}

func Multiply(inventory Inventory, factor int) Inventory {
	resultInventory := Inventory{}
	for i := 0; i < len(inventory.Items); i++ {
		resultItem := inventory.Items[i]
		resultItem.MinQTY = resultItem.MinQTY * factor
		resultInventory.Items = append(resultInventory.Items, resultItem)
	}
	return resultInventory
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
