package elements

import "github.com/tel21a-inf2/binarytrees/data"

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
	return element.data.Word
}

// Zugriffsfunktion für die gesamten Daten des Elements.
func (element Element) Data() *data.DictEntry {
	return element.data
}