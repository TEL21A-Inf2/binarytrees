package elements

import (
	"fmt"

	"github.com/tel21a-inf2/binarytrees/data"
)

func ExampleElement_DictEntry() {
	element1 := NewElement()
	fmt.Println(element1.IsEmpty())

	data1 := data.NewDictEntry("Haus", "house", "building")
	element1.setData(data1)
	fmt.Println(element1.IsEmpty())
	fmt.Println(element1.Key())
	fmt.Println(element1.Data())

	// Output:
	// true
	// false
	// Haus
	// Haus: house, building
}
