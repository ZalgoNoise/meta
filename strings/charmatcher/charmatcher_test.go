package charmatcher

import (
	"testing"
)

func TestNewCharMatcher(t *testing.T) {
	var funcName string = `TestNewCharMatcher`

	tests := []struct {
		input interface{}
		ok    bool
	}{
		{
			input: `Test string`,
			ok:    true,
		}, {
			input: `日本語`,
			ok:    true,
		}, {
			input: []rune(`Test runes`),
			ok:    true,
		}, {
			input: []rune(`日本語`),
			ok:    true,
		}, {
			input: 1234,
			ok:    false,
		}, {
			input: 1.234,
			ok:    false,
		}, {
			input: nil,
			ok:    false,
		},
	}

	for _, test := range tests {
		_, err := NewCharMatcher(test.input)

		if err != nil && test.ok != false {
			t.Errorf(`%q(%q) = %q, expected error to be %v`, funcName, test.input, err, test.ok)
		}
	}
}

func TestMatches(t *testing.T) {
	var funcName string = `TestMatches`

	tests := []struct {
		base    *CharMatcher
		input   interface{}
		matches bool
		ok      bool
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input:   `日本語`,
			matches: true,
			ok:      true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input:   `日語本`,
			matches: false,
			ok:      true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input:   []rune(`日本語`),
			matches: true,
			ok:      true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input:   []rune(`日語本`),
			matches: false,
			ok:      true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			matches: true,
			ok:      true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input: func() *CharMatcher {
				m, _ := NewCharMatcher(`日語本`)
				return m
			}(),
			matches: false,
			ok:      true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input:   123,
			matches: false,
			ok:      false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`日本語`)
				return m
			}(),
			input:   1.23,
			matches: false,
			ok:      false,
		},
	}

	for _, test := range tests {

		if _, err := test.base.Matches(test.input); err != nil && test.ok != false {
			t.Errorf(`%q(%q) = %q, expected error to be %v`, funcName, test.input, err, test.ok)
		}
	}
}

func TestMatchesAllOf(t *testing.T) {
	var funcName string = `TestMatchesAllOf`

	tests := []struct {
		base  *CharMatcher
		input interface{}
		ok    bool
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `abc`,
			ok:    true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `ab日語本`,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `123`,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `ABC`,
			ok:    false,
		},
	}

	for _, test := range tests {
		if ok := test.base.MatchesAllOf(test.input); ok != test.ok {
			t.Errorf(`%q(%q) = %v, expected match to be %v`, funcName, test.input, ok, test.ok)
		}
	}
}

func TestContainsAnyOf(t *testing.T) {
	var funcName string = `TestContainsAnyOf`

	tests := []struct {
		base  *CharMatcher
		input string
		ok    bool
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `日本語x`,
			ok:    true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `日語本`,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `ABCdEFG`,
			ok:    true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `ABCDEFG`,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `01234`,
			ok:    false,
		},
	}

	for _, test := range tests {
		if ok := test.base.ContainsAnyOf(test.input); ok != test.ok {
			t.Errorf(`%q(%q) = %v, expected match to be %v`, funcName, test.input, ok, test.ok)
		}
	}

}

func TestContains(t *testing.T) {
	var funcName string = `TestContains`

	tests := []struct {
		base  *CharMatcher
		input string
		ok    bool
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `abc`,
			ok:    true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `ab日語本`,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `123`,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: `ABC`,
			ok:    false,
		},
	}

	for _, test := range tests {
		if ok := test.base.Contains(test.input); ok != test.ok {
			t.Errorf(`%q(%q) = %v, expected match to be %v`, funcName, test.input, ok, test.ok)
		}
	}
}

func TestIndex(t *testing.T) {
	var funcName string = `TestIndex`

	tests := []struct {
		base  *CharMatcher
		input rune
		index int
		ok    bool
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'e',
			index: 4,
			ok:    true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'c',
			index: 20,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'a',
			index: 0,
			ok:    true,
		},
	}

	for _, test := range tests {
		if ok := test.base.Index(test.input, test.index); ok != test.ok {
			t.Errorf(`%q(%q) = %v, expected match to be %v`, funcName, test.input, ok, test.ok)
		}
	}
}

func TestIndexOf(t *testing.T) {
	var funcName string = `TestIndexOf`

	tests := []struct {
		base  *CharMatcher
		input rune
		index int
		ok    bool
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'e',
			index: 4,
			ok:    true,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'A',
			index: -1,
			ok:    false,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'a',
			index: 0,
			ok:    true,
		},
	}

	for _, test := range tests {
		if idx, ok := test.base.IndexOf(test.input); ok != test.ok && idx != test.index {
			t.Errorf(`%q(%q) = %v, expected match to be %v`, funcName, test.input, ok, test.ok)
		}
	}

}

func TestIndexesOf(t *testing.T) {
	var funcName string = `TestIndexesOf`

	tests := []struct {
		base  *CharMatcher
		input rune
		index []int
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'e',
			index: []int{4},
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcdefghijklmnopqrstuvwxyz`)
				return m
			}(),
			input: 'A',
			index: []int{-1},
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`Hello world`)
				return m
			}(),
			input: 'l',
			index: []int{2, 3, 9},
		},
	}

	for _, test := range tests {
		idx := test.base.IndexesOf(test.input)

		for i, v := range idx {

			if v != test.index[i] {
				t.Errorf(`%q(%q) = %v, expected match to be %v; %v != %v`, funcName, test.input, idx, test.index, v, test.index[i])
			}

		}

	}
}

func TestCompareTo(t *testing.T) {
	var funcName string = `TestCompareTo`

	tests := []struct {
		base     *CharMatcher
		input    string
		expected int
	}{
		{
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcde`)
				return m
			}(),
			input:    `abcde`,
			expected: 0,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcde`)
				return m
			}(),
			input:    `abcdf`,
			expected: 1,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcde`)
				return m
			}(),
			input:    `abcdd`,
			expected: -1,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcde`)
				return m
			}(),
			input:    `a`,
			expected: -17325400,
		}, {
			base: func() *CharMatcher {
				m, _ := NewCharMatcher(`abcde`)
				return m
			}(),
			input:    `abcdef`,
			expected: 606389365,
		},
	}

	for _, test := range tests {
		result := test.base.CompareTo(test.input, 36)

		if result != test.expected {
			t.Errorf(`%q(%q) = %v, expected match to be %v`, funcName, test.input, result, test.expected)
		}

	}
}
