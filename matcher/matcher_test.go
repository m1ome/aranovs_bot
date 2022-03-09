package matcher

import "testing"

func TestMatches(t *testing.T) {
	positiveMatches := []string{"да", "Да", "Да!!!", "да))))"}
	for _, s := range positiveMatches {
		sticker := Matches(s)
		if sticker == nil {
			t.Fatalf("expect matches positive")
		}
	}

	negativeMatches := []string{"нет", "Нет", "Нет!!!", "нЕт!))!"}
	for _, s := range negativeMatches {
		sticker := Matches(s)
		if sticker == nil {
			t.Fatalf("expect matches negative")
		}
	}

	falseMatcher := []string{"Даша", "Некто"}
	for _, s := range falseMatcher {
		sticker := Matches(s)
		if sticker != nil {
			t.Fatalf("expect not matches anything")
		}
	}

}
