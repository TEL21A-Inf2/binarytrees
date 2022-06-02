package trees

import (
	"fmt"

	"github.com/tel21a-inf2/binarytrees/data"
)

func ExampleTree_GetValue() {
	t1 := NewTree()

	data1 := data.NewDictEntry("Haus", "house", "building")
	t1.Add(data1)
	data2 := data.NewDictEntry("Katze", "cat", "tiger")
	t1.Add(data2)
	data3 := data.NewDictEntry("Fahrrad", "bicycle")
	t1.Add(data3)
	fmt.Println(t1.GetValue("Haus"))
	fmt.Println(t1.GetValue("Katze"))
	fmt.Println(t1.GetValue("Fahrrad"))
	fmt.Println(t1.GetValue("Nichts"))

	// Output:
	// [house building] <nil>
	// [cat tiger] <nil>
	// [bicycle] <nil>
	// [] Key nicht im Baum enthalten
}

func ExampleTree_MermaidString() {
	t1 := NewTree()

	data1 := data.NewDictEntry("Haus", "house", "building")
	t1.Add(data1)
	data2 := data.NewDictEntry("Katze", "cat", "tiger")
	t1.Add(data2)
	data3 := data.NewDictEntry("Fahrrad", "bicycle")
	t1.Add(data3)
	fmt.Println(t1.MermaidString())

	// Output:
	// graph TD
	//   Haus --> Fahrrad
	//   Haus --> Katze
}

func ExampleElement_Remove_Leaf() {
	t1 := NewTree()

	data1 := data.NewDictEntry("Haus", "house", "building")
	t1.Add(data1)
	data2 := data.NewDictEntry("Katze", "cat", "tiger")
	t1.Add(data2)
	data3 := data.NewDictEntry("Fahrrad", "bicycle")
	t1.Add(data3)

	// Fahrrad entfernen:
	t1.Remove("Fahrrad")

	// Haus und Katze sollten noch da sein, Fahrrad nicht mehr.
	fmt.Println(t1.GetValue("Haus"))
	fmt.Println(t1.GetValue("Katze"))
	fmt.Println(t1.GetValue("Fahrrad"))

	// Katze entfernen:
	t1.Remove("Katze")

	// Haus sollte noch da sein, Katze nicht mehr
	fmt.Println(t1.GetValue("Haus"))
	fmt.Println(t1.GetValue("Katze"))

	// Katze entfernen:
	t1.Remove("Haus")
	fmt.Println(t1.GetValue("Haus"))

	// Output:
	// [house building] <nil>
	// [cat tiger] <nil>
	// [] Key nicht im Baum enthalten
	// [house building] <nil>
	// [] Key nicht im Baum enthalten
	// [] Key nicht im Baum enthalten
}

func ExampleElement_Remove_SemiLeaf() {
	t1 := NewTree()

	data1 := data.NewDictEntry("Haus", "house", "building") // Die Wurzel ist ein Halbblatt, nur der linke Nachfolger existiert.
	t1.Add(data1)
	data2 := data.NewDictEntry("Fahrrad", "bicycle")
	t1.Add(data2)
	data3 := data.NewDictEntry("Auto", "car") // Auto ist ein Halb-Blatt, es hat nur einen rechten Nachfolger.
	t1.Add(data3)
	data4 := data.NewDictEntry("Blume", "flower")
	t1.Add(data4)
	data5 := data.NewDictEntry("Bauer", "farmer") // Links unterhalb von Blume
	t1.Add(data5)
	data6 := data.NewDictEntry("Esel", "donkey") // Rechts unterhalb von Blume
	t1.Add(data6)

	t1.Remove("Auto")

	// Wir lassen uns die Fehlerwerte von GetValue geben.
	_, e1 := t1.GetValue("Haus")
	_, e2 := t1.GetValue("Fahrrad")
	_, e3 := t1.GetValue("Auto")
	_, e4 := t1.GetValue("Blume")
	_, e5 := t1.GetValue("Bauer")
	_, e6 := t1.GetValue("Esel")

	// Außer Auto sollten alle Fehlerwerte nil sein.
	fmt.Println(e1 == nil)
	fmt.Println(e2 == nil)
	fmt.Println(e3 != nil)
	fmt.Println(e4 == nil)
	fmt.Println(e5 == nil)
	fmt.Println(e6 == nil)

	// Wurzel zu einem Halbblatt machen und entfernen.
	t1.Remove("Haus")

	// Wir lassen uns noch einmal die Fehlerwerte von GetValue geben.
	_, e1 = t1.GetValue("Haus")
	_, e2 = t1.GetValue("Fahrrad")
	_, e3 = t1.GetValue("Blume")
	_, e4 = t1.GetValue("Bauer")
	_, e5 = t1.GetValue("Esel")

	// Dieses Mal sollte der erste Fehler existieren.
	fmt.Println(e1 != nil)
	fmt.Println(e2 == nil)
	fmt.Println(e3 == nil)
	fmt.Println(e4 == nil)
	fmt.Println(e5 == nil)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}

func ExampleElement_Remove_Inner() {
	t1 := NewTree()

	data1 := data.NewDictEntry("Haus", "house", "building")
	t1.Add(data1)
	data2 := data.NewDictEntry("Katze", "cat", "tiger")
	t1.Add(data2)
	data3 := data.NewDictEntry("Fahrrad", "bicycle")
	t1.Add(data3)
	data4 := data.NewDictEntry("Blume", "flower")
	t1.Add(data4)
	data5 := data.NewDictEntry("Bauer", "farmer")
	t1.Add(data5)
	data6 := data.NewDictEntry("Esel", "donkey")
	t1.Add(data6)

	t1.Remove("Blume")

	// Wir lassen uns die Fehlerwerte von GetValue geben.
	_, e1 := t1.GetValue("Haus")
	_, e2 := t1.GetValue("Fahrrad")
	_, e3 := t1.GetValue("Katze")
	_, e4 := t1.GetValue("Blume")
	_, e5 := t1.GetValue("Bauer")
	_, e6 := t1.GetValue("Esel")

	// Außer Auto sollten alle Fehlerwerte nil sein.
	fmt.Println(e1 == nil)
	fmt.Println(e2 == nil)
	fmt.Println(e3 == nil)
	fmt.Println(e4 != nil)
	fmt.Println(e5 == nil)
	fmt.Println(e6 == nil)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
