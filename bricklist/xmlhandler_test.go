/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package bricklist

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

var item1 = Item{
	ItemType:  "P",
	ItemID:    "2420",
	Color:     "69",
	Condition: "X",
	Maxprice:  "-1.0000",
	MinQTY:    10,
	Remarks:   "Item 1",
	Notify:    "N",
}
var item2 = Item{
	ItemType:  "P",
	ItemID:    "2420",
	Color:     "6",
	Condition: "U",
	Maxprice:  "-1.0000",
	MinQTY:    6,
	Remarks:   "Item 2",
	Notify:    "N",
}
var item3 = Item{
	ItemType:  "P",
	ItemID:    "2431",
	Color:     "4",
	Condition: "X",
	Maxprice:  "0.2000",
	MinQTY:    10,
	Remarks:   "Item 3",
	Notify:    "N",
}
var item4 = Item{
	ItemType: "P",
	ItemID:   "2431",
	Color:    "5",
	MinQTY:   10,
}

func TestReadXmlList(t *testing.T) {
	inventory := ReadXmlList("../testdata/four_items.xml")
	itemMap := createItemMap(inventory)

	if len(itemMap) != 4 {
		t.Fatalf(`Expected that item map has length 4, but it has length %v.`, len(itemMap))
	}
	assertMapEntry(t, itemMap, item1)
	assertMapEntry(t, itemMap, item2)
	assertMapEntry(t, itemMap, item3)
	assertMapEntry(t, itemMap, item4)
}

func TestRenderXmlMultiline(t *testing.T) {
	inventory := Inventory{Items: []Item{item1, item2, item3, item4}}
	expectedXml := readFile(t, "../testdata/four_items.xml")
	renderedXml := RenderSingleXML(inventory, true)
	if renderedXml != expectedXml {
		t.Fatalf(`Expected xml: "%v", but got "%v".`, expectedXml, renderedXml)
	}
}

func TestRenderXmlSingleline(t *testing.T) {
	inventory := Inventory{Items: []Item{item1, item2, item3, item4}}
	expectedXml := readFile(t, "../testdata/four_items.xml")
	re := regexp.MustCompile(`\s*\n\s*`)
	expectedXml = re.ReplaceAllString(expectedXml, "") + "\n"
	renderedXml := RenderSingleXML(inventory, false)
	if renderedXml != expectedXml {
		t.Fatalf(`Expected xml: "%v", but got "%v".`, expectedXml, renderedXml)
	}
}

func readFile(t *testing.T, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf(`Error opening file %v: %v.`, filename, err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	return string(buf)
}
