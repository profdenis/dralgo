package gen

import "testing"

func TestLookupLanguagesAndAliases(t *testing.T) {
	cases := map[string]bool{
		"puml":       true,
		"plantuml":   true,
		"python":     true,
		"go":         true,
		"js":         true,
		"javascript": true,
		"java":       true,
		"cobol":      false,
		"":           false,
	}

	for lang, want := range cases {
		_, ok := Lookup(lang)
		if ok != want {
			t.Errorf("Lookup(%q) ok = %v, want %v", lang, ok, want)
		}
	}
}

func TestLanguagesAreResolvable(t *testing.T) {
	for _, lang := range Languages() {
		if _, ok := Lookup(lang); !ok {
			t.Errorf("Lookup(%q) ok = false for a listed language", lang)
		}
	}
}
