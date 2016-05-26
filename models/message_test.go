package models

import "testing"

var testCases = []struct {
	input  []string
	output string
}{
	{
		input:  []string{""},
		output: "",
	},
	{
		input:  []string{"John"},
		output: "John",
	},
	{
		input:  []string{"John", "Steph"},
		output: "John and Steph",
	},
	{
		input:  []string{"John", "Steph", "Bob"},
		output: "John, Steph and Bob",
	},
	{
		input:  []string{"John", "Steph", "Bob", "Derp"},
		output: "John, Steph, Bob and Derp",
	},
}

func TestToSentence(t *testing.T) {
	for _, test := range testCases {
		got := toSentence(test.input)
		expected := test.output

		if got != expected {
			t.Errorf("ERROR: Got %s, expected %s", got, expected)
		}

	}
}
