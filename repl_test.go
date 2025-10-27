package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "",
			expected: []string{""},
		},
		{
			input: " hello world",
			expected: []string{"hello","world"},
		},
		{
			input: " this is a    test   ",
			expected: []string{"this","is","a","test"},
		},
	}


	for _, c:= range cases {
		actual := CleanInput(c.input)

		if len(c.expected) != len(actual){
			t.Errorf("Failed test, clean input length(%v) not equal to expected(%v)", len(actual), len(cases))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("Failed test, input word(%v) does not match expected word(%v)", word, expectedWord)
			}
		}
	}
}

// func testCommands(t *testing.T){
// 	cases := []struct {
// 		input string
// 		expected string
// 	}{
// 		{
// 			input: "",
// 			expected: "Unknown command",
// 		},
// 		{
// 			input: "help",
// 			expected: "help",
// 		},
// 		{
// 			input: "exit",
// 			expected: "exit",
// 		},
// 	}
// 	for i :=1; i >0; i++{
// 		fmt.Printf("Pokedex >")
// 		for _, command := range cases{
// 			input := command.input
// 				if input == commands[input].name{
// 					t.Error("command not being read correctly")
// 				}
// 			}
// 		}
// }
