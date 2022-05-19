package elements

// Interface, das die Eigenschaften von geeigneten Elementdaten beschreibt.
type ElementData[KeyType comparable, ValueType any] interface {
	Key() KeyType
	Value() ValueType
	String() string
}
