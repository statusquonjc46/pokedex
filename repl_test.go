package main

import "testing"

func testCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello,   world  !",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "PIKACHU CHARMANDER",
			expected: []string{"pikachu", "charmander"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expecting length '%d', but got '%d'", c, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expecting word '%s', but received '%s'", expectedWord, word)
			}
		}
	}
}
