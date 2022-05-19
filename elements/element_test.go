package elements

import (
	"fmt"

	"github.com/tel21a-inf2/binarytrees/data"
)

func ExampleElement_DictEntry() {
	element1 := NewElement[string, []string]()
	fmt.Println(element1.IsEmpty())

	data1 := data.NewDictEntry("Haus", "house", "building")
	element1.setData(data1)
	fmt.Println(element1.IsEmpty())
	fmt.Println(element1.Key())
	fmt.Println(element1.Value())
	fmt.Println(element1.Data())

	// Output:
	// true
	// false
	// Haus
	// [house building]
	// Haus: house, building
}

func ExampleElement_String() {
	element1 := NewElement[string, []string]()
	fmt.Println(element1)

	data1 := data.NewDictEntry("Haus", "house", "building")
	element1.setData(data1)
	fmt.Println(element1)

	// Output:
	// <empty>
	// Haus: house, building
}

func ExampleElement_InOrderString() {
	element1 := NewElement[string, []string]()
	fmt.Println(element1.InOrderString()) // Keine Ausgabe (leerer String)

	data1 := data.NewDictEntry("Haus", "house", "building")
	element1.setData(data1)
	fmt.Println(element1.InOrderString())

	data2 := data.NewDictEntry("Katze", "cat", "tiger")
	element1.right.setData(data2)
	fmt.Println(element1.InOrderString())

	data3 := data.NewDictEntry("Fahrrad", "bicycle")
	element1.left.setData(data3)
	fmt.Println(element1.InOrderString())

	// Output:
	// Haus: house, building
	//
	// Haus: house, building
	// Katze: cat, tiger
	//
	// Fahrrad: bicycle
	// Haus: house, building
	// Katze: cat, tiger
}
