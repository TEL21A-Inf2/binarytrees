package data

import "fmt"

func ExampleDictEntry() {
	e1 := NewDictEntry("Haus")
	fmt.Println(e1)

	e1.AddTranslation("house")
	fmt.Println(e1)

	e1.AddTranslation("building")
	fmt.Println(e1)

	fmt.Println(e1.Key())
	fmt.Println(e1.Value())

	e2 := NewDictEntry("Katze", "cat", "tiger")
	fmt.Println(e2)

	// Output:
	// Haus: <Keine Übersetzung>
	// Haus: house
	// Haus: house, building
	// Haus
	// [house building]
	// Katze: cat, tiger
}
