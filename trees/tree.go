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
		return
	}

	if elementToRemove.IsSemiLeaf() {
		if elementToRemove == tree.root {
			tree.root = tree.root.GetNonEmptyChild()
			return
		}
		child := elementToRemove.GetNonEmptyChild()
		parentOfElementToRemove.ReplaceChild(elementToRemove, child)
		return
	}

	inOrderSuccessor := elementToRemove.GetInOrderSuccessor()
	tree.root.Swap(elementToRemove, inOrderSuccessor)
	// Ab dem rechten Kind des elementToRemove rekursiv weitermachen.
	// Hier wenden wir einen (schmutzigen) Trick an und konstruieren Hilfsweise einen Baum,
	// dessen Wurzel das rechte Kind des elementToRemove ist.
	t := Tree{elementToRemove.GetRight()}
	t.Remove(key)
}
