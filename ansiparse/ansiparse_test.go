package ansiparse

import "testing"

var MeasureTextAreaTests = []struct {
	text     string
	expected measuredText
}{
	{"test 1", measuredText{1, 6}},
	{"foo", measuredText{1, 3}},
	{"foo\nbar", measuredText{2, 3}},
	{"🇪🇸", measuredText{1, 2}},
	{"こんにちは", measuredText{1, 10}},
}

func TestMeasueTextArea(t *testing.T) {
	for _, tt := range MeasureTextAreaTests {
		t.Run(tt.text, func(t *testing.T) {
			got := measureTextArea(tt.text)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
