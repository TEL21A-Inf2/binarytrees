package elements

import (
	"fmt"

	"github.com/tel21a-inf2/binarytrees/data"
)

// Datentyp für Elemente eines Baumes.
type Element struct {
	data        *data.DictEntry
	left, right *Element
}

// Konstruktor für Baum-Elemente.
func NewElement() *Element {
	return &Element{nil, nil, nil}
}

// Liefert true, falls dieses Element leer ist.
func (element Element) IsEmpty() bool {
	// Es wird nur geprüft, ob data==nil ist.
	// Wir werden bei den Baum-Implementierungen dafür sorgen,
	// dass in diesem Fall auch die beiden Kind-Pointer nil sind.
	return element.data == nil
}

// Leert das Element.
// D.h. löscht die Daten, leert rekursiv die Kinder und löscht sie.
func (element *Element) Clear() {
	// Diese Funktion ist wichtig, um keine Speicherlecks zu hinterlassen.
	element.data = nil
	if element.left != nil {
		element.left.Clear()
		element.left = nil
	}
	if element.right != nil {
		element.right.Clear()
		element.right = nil
	}
}

// Setzt ein neues Datenelement.
// Falls das Element bisher leer war, werden zwei neue leere Kinder angehängt.
func (element *Element) setData(newData *data.DictEntry) {
	if element.IsEmpty() {
		element.left = NewElement()
		element.right = NewElement()
	}
	element.data = newData
}

// Zugriffsfunktion für den Schlüssel des Elements.
func (element Element) Key() string {
	return element.data.Key()
}

// Zugriffsfunktion für den Wert des Elements.
func (element Element) Value() []string {
	return element.data.Value()
}

// Zugriffsfunktion für die gesamten Daten des Elements.
func (element Element) Data() *data.DictEntry {
	return element.data
}

// String-Darstellung von Elementen.
func (element Element) String() string {
	if element.IsEmpty() {
		return "<empty>"
	}
	return element.data.String()
}

// Gibt den Baum in In-Order-Darstellung aus.
func (element Element) InOrderString() string {
	if element.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%s%s\n%s", element.left.InOrderString(), element.data, element.right.InOrderString())
}

// Fügt ein Element zum Baum hinzu.
func (element *Element) Add(newData *data.DictEntry) {
	if element.IsEmpty() {
		element.setData(newData)
		return
	}
	if newData.Key() < element.Key() {
		element.left.Add(newData)
	} else {
		element.right.Add(newData)
	}
}

// Liefert einen Pointer auf den Datentsatz mit dem angegebenen Key.
func (element Element) Get(key string) *data.DictEntry {
	if element.IsEmpty() {
		return nil
	}
	if element.Key() == key {
		return element.Data()
	}
	if key < element.Key() {
		return element.left.Get(key)
	}
	return element.right.Get(key)
}

// Liefert den Wert zum angegebenen Key.
func (element Element) GetValue(key string) ([]string, error) {
	data := element.Get(key)
	if data == nil {
		return make([]string, 0), fmt.Errorf("Key nicht im Baum enthalten")
	}
	return data.Value(), nil
}

// Liefert die Anzahl der Elemente im Baum.
func (element Element) Size() int {
	if element.IsEmpty() {
		return 0
	}
	return 1 + element.left.Size() + element.right.Size()
}

func (element Element) Height() int {
	if element.IsEmpty() {
		return 0
	}
	leftHeight, rightHeight := element.left.Height(), element.right.Height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
