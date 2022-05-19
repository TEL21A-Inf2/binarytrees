package elements

// Interface, das die Eigenschaften von geeigneten Elementdaten beschreibt.
type ElementData interface {
	Key() string
	Value() []string
}
