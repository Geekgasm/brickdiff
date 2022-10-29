/*
Copyright © 2022 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_GivenListWithTwoItems_whenSubtractingListWithOneItem_thenListHasOnlyOneItem(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x10_2x6.xml")
	inventory2 := ReadXmlList("../testdata/item_2x6.xml")
	expectedResult := ReadXmlList("../testdata/item_1x10.xml")
	result := SubtractInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}

func Test_GivenListWithOneItemQ10_whenSubtractingListWithOneItemQ6_thenListHasOnlyOneItemQ4(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x10.xml")
	inventory2 := ReadXmlList("../testdata/item_1x6.xml")
	expectedResult := ReadXmlList("../testdata/item_1x4.xml")
	result := SubtractInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}

func Test_GivenListWithOneItemQ10_whenSubtractingListWithOneItemDifferentConditionQ6_thenListHasOnlyOneItemQ4(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x10.xml")
	inventory2 := ReadXmlList("../testdata/item_1x6used.xml")
	expectedResult := ReadXmlList("../testdata/item_1x4.xml")
	result := SubtractInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}

func Test_GivenListWithOneItemQ6_whenSubtractingListWithOneItemQ10_thenListIsEmpty(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x6.xml")
	inventory2 := ReadXmlList("../testdata/item_1x10.xml")
	result := SubtractInventories(inventory1, inventory2)
	assert.Assert(t, len(result.Items) == 0)
}

func Test_GivenListWithOneItemQ10_whenSubtractingListWithOneItemQ10_thenListIsEmpty(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x10.xml")
	result := SubtractInventories(inventory1, inventory1)
	assert.Assert(t, len(result.Items) == 0)
}

func Test_GivenDisjointLists_whenSubtractingLists_thenResultIsMinuend(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x10_2x6.xml")
	inventory2 := ReadXmlList("../testdata/item_3x10_4x10.xml")
	expectedResult := inventory1
	result := SubtractInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}

func Test_GivenListWithOneItemQ4_whenAddingListWithOneItemQ6_thenResultHasItemQ10(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x4.xml")
	inventory2 := ReadXmlList("../testdata/item_1x6.xml")
	expectedResult := ReadXmlList("../testdata/item_1x10.xml")
	result := AddInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}

func Test_GivenListWithOneItemQ4_whenAddingListWithOneItemQ6used_thenResultHasItemQ10(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x4.xml")
	inventory2 := ReadXmlList("../testdata/item_1x6used.xml")
	expectedResult := ReadXmlList("../testdata/item_1x10.xml")
	result := AddInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}

func Test_GivenListWithItem1Q10Item2Q6_whenAddingListWithItem2Q6Item3Q10_thenResultHasItem1Q10Item2Q12Item3Q10(t *testing.T) {
	inventory1 := ReadXmlList("../testdata/item_1x10_2x6.xml")
	inventory2 := ReadXmlList("../testdata/item_2x6_3x10.xml")
	expectedResult := ReadXmlList("../testdata/item_1x10_2x12_3x10.xml")
	result := AddInventories(inventory1, inventory2)
	assert.DeepEqual(t, createItemMap(result), createItemMap(expectedResult))
}
