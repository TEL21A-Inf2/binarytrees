package trees

import (
	"fmt"

	"github.com/tel21a-inf2/binarytrees/data"
	"github.com/tel21a-inf2/binarytrees/elements"
)

type Tree struct {
	root *elements.Element
}

// Konstruktor für Bäume
func NewTree() *Tree {
	return &Tree{elements.NewElement()}
}

// Fügt ein Element zum Baum hinzu.
func (tree *Tree) Add(newData *data.DictEntry) {
	tree.root.Add(newData)
}

// Liefert den Wert zum angegebenen Key.
func (tree Tree) GetValue(key string) ([]string, error) {
	return tree.root.GetValue(key)
}

// Liefert einen Mermaid-String für den Baum.
func (tree Tree) MermaidString() string {
	return fmt.Sprintf("graph TD\n%s", tree.root.MermaidStrings())
}

// Entfernt den Knoten mit dem angegebenen Element aus dem Baum.
func (tree *Tree) Remove(key string) {
	elementToRemove := tree.root.GetElement(key)
	parentOfElementToRemove := tree.root.GetParent(elementToRemove)

	if elementToRemove.IsLeaf() {
		if elementToRemove == tree.root {
			elementToRemove.Clear()
			return
		}
		parentOfElementToRemove.RemoveChild(elementToRemove)
	}
}
