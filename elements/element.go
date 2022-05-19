package elements

// Datentyp für Elemente eines Baumes.
type Element[KeyType comparable, ValueType any] struct {
	data        ElementData[KeyType, ValueType]
	left, right *Element[KeyType, ValueType]
}

// Konstruktor für Baum-Elemente.
func NewElement[KeyType comparable, ValueType any]() *Element[KeyType, ValueType] {
	return &Element[KeyType, ValueType]{nil, nil, nil}
}

// Liefert true, falls dieses Element leer ist.
func (element Element[KeyType, ValueType]) IsEmpty() bool {
	// Es wird nur geprüft, ob data==nil ist.
	// Wir werden bei den Baum-Implementierungen dafür sorgen,
	// dass in diesem Fall auch die beiden Kind-Pointer nil sind.
	return element.data == nil
}

// Leert das Element.
// D.h. löscht die Daten, leert rekursiv die Kinder und löscht sie.
func (element *Element[KeyType, ValueType]) Clear() {
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
func (element *Element[KeyType, ValueType]) setData(newData ElementData[KeyType, ValueType]) {
	if element.IsEmpty() {
		element.left = NewElement[KeyType, ValueType]()
		element.right = NewElement[KeyType, ValueType]()
	}
	element.data = newData
}

// Zugriffsfunktion für den Schlüssel des Elements.
func (element Element[KeyType, ValueType]) Key() KeyType {
	return element.data.Key()
}

// Zugriffsfunktion für den Wert des Elements.
func (element Element[KeyType, ValueType]) Value() ValueType {
	return element.data.Value()
}

// Zugriffsfunktion für die gesamten Daten des Elements.
func (element Element[KeyType, ValueType]) Data() ElementData[KeyType, ValueType] {
	return element.data
}
