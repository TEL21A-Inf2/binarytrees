package data

// Datentyp für Elemente eines Wörterbuchs.
type DictEntry struct {
	Word         string   // Wort in Quellsprache
	Translations []string // Mögliche Übersetzungen
}

// Konstruktor für einen neuen Eintrag eines Wörterbuchs.
// Das Quell-Wort muss angegeben werden.
// Es können auch schon Übersetzungen angegeben werden,
// diese können aber auch später mittels AddTranslation hinzugefügt werden.
func NewDictEntry(word string, translations ...string) *DictEntry {
	return &DictEntry{word, translations}
}

// Fügt eine Übersetzung hinzu.
func (entry *DictEntry) AddTranslation(translation string) {
	entry.Translations = append(entry.Translations, translation)
}

// Liefert die Anzahl der vorhandenen Übersetzungen.
func (entry DictEntry) Count() int {
	return len(entry.Translations)
}

// Liefert true, falls der Eintrag leer ist.
func (entry DictEntry) IsEmpty() bool {
	return entry.Count() == 0
}

// String-Ausgabe für Wörterbucheinträge.
func (entry DictEntry) String() string {
	result := entry.Word + ": "
	if entry.IsEmpty() {
		result += "<Keine Übersetzung>"
	} else {
		for _, t := range entry.Translations {
			result += t + ", "
		}
		result = result[:len(result)-2]
	}
	return result
}

// Generisch benannte Zugriffsfunktion für den Schlüssel eines Eintrags.
func (entry DictEntry) Key() string {
	return entry.Word
}

// Generisch benannte Zugriffsfunktion für den Wert eines Eintrags.
func (entry DictEntry) Value() []string {
	return entry.Translations
}
