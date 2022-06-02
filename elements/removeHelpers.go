package elements

// Hilfsfunktionen für Tree.Remove().

// Liefert den Elternknoten des Knotens mit dem angegebenen Key.
func (element *Element) GetParent(searchElement *Element) *Element {

	// Basisfälle: In diesen Fällen kann das gesuchte Element nicht enthalten sein.
	if element == nil || element.IsEmpty() || element.IsLeaf() {
		return nil
	}
	// Falls das folgende zutrifft, sind wir zu weit:
	if element == searchElement {
		return nil
	}
	// Wenn eines der Kinder == searchElement ist, haben wir Parent gefunden.
	if element.left == searchElement || element.right == searchElement {
		return element
	}

	// Normales Suchschema: Je nach Key links oder rechts weitersuchen.
	if searchElement.Key() < element.Key() {
		return element.left.GetParent(searchElement)
	}
	return element.right.GetParent(searchElement)
}

// Liefert das Element mit dem angegebenen Key.
// Wie Get(), aber liefert das Element, nicht den Datensatz.
func (element *Element) GetElement(key string) *Element {
	if element.IsEmpty() {
		return nil
	}
	if element.Key() == key {
		return element
	}
	if key < element.Key() {
		return element.left.GetElement(key)
	}
	return element.right.GetElement(key)
}

// Liefert true, falls das Element ein Blatt ist.
func (element Element) IsLeaf() bool {
	return element.left.IsEmpty() && element.right.IsEmpty()
}

// Löscht das Kind von element, das gleich toRemove ist.
func (element *Element) RemoveChild(toRemove *Element) {
	if toRemove == element.left {
		element.left.Clear()
	}
	if toRemove == element.right {
		element.right.Clear()
	}
}

// Liefert true, wenn das Element ein Halbblatt ist, also nur ein nicht-leeres Kind hat.
func (element Element) IsSemiLeaf() bool {
	return (element.left.IsEmpty() && !element.right.IsEmpty()) || (!element.left.IsEmpty() && element.right.IsEmpty())
}

// Liefert das einzige Kind, falls das Element ein Halbblatt ist.
// Hat undefiniertes Verhalten, falls das Element kein Halbblatt ist.
// (D.h. es geschieht keine Fehlerbehandlung.)
func (element *Element) GetNonEmptyChild() *Element {
	if element.left.IsEmpty() {
		return element.right
	}
	return element.left
}

// Ersetzt das angegebene Kind durch den angegebenen Knoten.
// Auch hier keine Fehlerbehandlung, d.h. das angegebene Kind muss wirklich eines der Kinder sein,
// sonst verhält sich auch diese Funktion nicht korrekt.
func (element *Element) ReplaceChild(elementToReplace, newChild *Element) {
	if element.left == elementToReplace {
		element.left = newChild
	} else {
		element.right = newChild
	}
}
