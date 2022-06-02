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
