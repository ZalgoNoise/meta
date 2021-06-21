package strbuilder

import (
	"testing"
)

var (
	empty       string  = ``
	base        string  = `abc`
	cttOne      string  = `def`
	cttTwo      string  = `ghi`
	cttThree    string  = `日本語`
	numOne      int     = 123
	numTwo      int     = 456
	numThree    int     = 789
	strconvNums string  = `123456789`
	errVal      float64 = 1.23456
)

func TestNewStringBuilder(t *testing.T) {
	var funcName string = `TestNewStringBuilder`
	tests := []struct {
		input    interface{}
		expected []rune
		cap      int
	}{
		{
			input:    100,
			expected: []rune{},
			cap:      100,
		}, {
			input:    `Test`,
			expected: []rune{'T', 'e', 's', 't'},
			cap:      20,
		}, {
			input:    `日本語`,
			expected: []rune{'日', '本', '語'},
			cap:      19,
		}, {
			input:    []rune(`Test`),
			expected: []rune{'T', 'e', 's', 't'},
			cap:      20,
		}, {
			input:    nil,
			expected: []rune{},
			cap:      16,
		},
	}

	for _, test := range tests {
		str := NewStringBuilder(test.input)

		if str.Capacity != test.cap {
			t.Errorf(`%q(%q) = cap: %v, expected cap: %v`, funcName, test.input, str.Capacity, test.cap)
		}

		for idx, v := range str.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected cap: %v`, funcName, test.input, v, test.expected[idx])
			}
		}
	}
}

func TestStringBuilderCompareTo(t *testing.T) {
	var funcName string = `TestStringBuilderCompareTo`

	tests := []struct {
		base     *StringBuilder
		input    *StringBuilder
		expected bool
	}{
		{
			base:     NewStringBuilder(`abc`),
			input:    NewStringBuilder(`abc`),
			expected: true,
		}, {
			base:     NewStringBuilder(`abc`),
			input:    NewStringBuilder(`cde`),
			expected: false,
		}, {
			base:     NewStringBuilder(`日本語`),
			input:    NewStringBuilder(`日本語`),
			expected: true,
		}, {
			base:     NewStringBuilder(`日本語`),
			input:    NewStringBuilder(`日語本`),
			expected: false,
		},
	}

	for _, test := range tests {

		if ok, err := test.base.CompareTo(test.input); err != nil && ok != test.expected {
			t.Errorf(`%q(%q) = %v, expected %v`, funcName, test.input, ok, test.expected)
		}
	}

}

func TestAppendStringBuilders(t *testing.T) {
	var funcName string = `TestAppendStringBuilders`
	tests := []struct {
		base     *StringBuilder
		input    []*StringBuilder
		expected []rune
	}{
		{
			base: NewStringBuilder(empty),
			input: []*StringBuilder{
				NewStringBuilder(base),
			},
			expected: []rune(base),
		}, {
			base: NewStringBuilder(nil),
			input: []*StringBuilder{
				NewStringBuilder(base),
				NewStringBuilder(cttOne),
				NewStringBuilder(cttTwo),
				NewStringBuilder(cttThree),
			},
			expected: []rune(base + cttOne + cttTwo + cttThree),
		},
	}
	for _, test := range tests {
		for _, x := range test.input {
			if err := test.base.Append(x); err != nil {
				t.Errorf(`%q(%q) = %v`, funcName, x, err)
			}
		}
		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q() = %q, expected %q`, funcName, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestAppendStrings(t *testing.T) {
	var funcName string = `TestAppendStrings`
	tests := []struct {
		base     *StringBuilder
		input    []string
		expected []rune
	}{
		{
			base: NewStringBuilder(base),
			input: []string{
				cttOne,
				cttTwo,
				cttThree,
			},
			expected: []rune(base + cttOne + cttTwo + cttThree),
		},
	}
	for _, test := range tests {
		for _, x := range test.input {
			if err := test.base.Append(x); err != nil {
				t.Errorf(`%q(%q) = %v`, funcName, x, err)
			}
		}
		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q() = %q, expected %q`, funcName, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestAppendInts(t *testing.T) {
	var funcName string = `TestAppendInts`
	tests := []struct {
		base     *StringBuilder
		input    []int
		expected []rune
	}{
		{
			base: NewStringBuilder(base),
			input: []int{
				numOne,
				numTwo,
				numThree,
			},
			expected: []rune(base + strconvNums),
		},
	}
	for _, test := range tests {
		for _, x := range test.input {
			if err := test.base.Append(x); err != nil {
				t.Errorf(`%q(%q) = %v`, funcName, x, err)
			}
		}
		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q() = %q, expected %q`, funcName, string(v), string(test.expected[idx]))
			}
		}
	}
}
func TestAppendRunes(t *testing.T) {
	var funcName string = `TestAppendRunes`
	tests := []struct {
		base     *StringBuilder
		input    [][]rune
		expected []rune
	}{
		{
			base: NewStringBuilder(base),
			input: [][]rune{
				[]rune(cttOne),
				[]rune(cttTwo),
				[]rune(cttThree),
			},
			expected: []rune(base + cttOne + cttTwo + cttThree),
		},
	}
	for _, test := range tests {
		for _, x := range test.input {
			if err := test.base.Append(x); err != nil {
				t.Errorf(`%q(%q) = %v`, funcName, x, err)
			}
		}
		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q() = %q, expected %q`, funcName, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestAppendErr(t *testing.T) {
	var funcName string = `TestAppendErr`
	tests := []struct {
		base     *StringBuilder
		input    []float64
		expected []rune
		ok       bool
	}{
		{
			base: NewStringBuilder(base),
			input: []float64{
				errVal,
			},
			expected: []rune(base),
			ok:       false,
		},
	}
	for _, test := range tests {
		for _, x := range test.input {
			if err := test.base.Append(x); err != nil && test.ok != false {
				t.Errorf(`%q(%v) = %v`, funcName, x, err)
			}
		}
		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q() = %q, expected %q`, funcName, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestDelete(t *testing.T) {
	var funcName string = `TestDelete`
	tests := []struct {
		base     *StringBuilder
		start    int
		end      int
		expected []rune
	}{
		{
			base:     NewStringBuilder(`Hello world`),
			start:    0,
			end:      2,
			expected: []rune(`lo world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    2,
			end:      30,
			expected: []rune(`He`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    -1,
			end:      5,
			expected: []rune(`Hello world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    8,
			end:      -5,
			expected: []rune(`Hello world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    0,
			end:      0,
			expected: []rune(`ello world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    6,
			end:      6,
			expected: []rune(`Hello orld`),
		},
	}

	for _, test := range tests {
		test.base.Delete(test.start, test.end)

		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.base.Output, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestDeleteAt(t *testing.T) {
	var funcName string = `TestDeleteAt`
	tests := []struct {
		base     *StringBuilder
		index    int
		expected []rune
	}{
		{
			base:     NewStringBuilder(`Hello world`),
			index:    0,
			expected: []rune(`ello world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			index:    1,
			expected: []rune(`Hllo world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			index:    500,
			expected: []rune(`Hello worl`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			index:    -4,
			expected: []rune(`Hello world`),
		},
	}

	for _, test := range tests {
		test.base.DeleteAt(test.index)

		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.base.Output, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestReplace(t *testing.T) {
	var funcName string = `TestReplace`
	tests := []struct {
		base     *StringBuilder
		start    int
		end      int
		sub      *StringBuilder
		expected []rune
		ok       bool
	}{
		{
			base:     NewStringBuilder(`Hello world`),
			start:    2,
			end:      4,
			sub:      NewStringBuilder(base),
			expected: []rune(`He` + base + ` world`),
			ok:       true,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    1,
			end:      9,
			sub:      NewStringBuilder(strconvNums),
			expected: []rune(`H` + strconvNums + `d`),
			ok:       true,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    4,
			end:      4,
			sub:      NewStringBuilder(`x`),
			expected: []rune(`Hellx world`),
			ok:       true,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    0,
			end:      4,
			sub:      NewStringBuilder(`x`),
			expected: []rune(`Hello world`),
			ok:       false,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    0,
			end:      -2,
			sub:      NewStringBuilder(base),
			expected: []rune(`Hello world`),
			ok:       false,
		},
	}

	for _, test := range tests {
		if err := test.base.Replace(test.start, test.end, test.sub); err != nil && test.ok != false {
			t.Errorf(`%q(start: %v, end: %v, sub: %q) has errored: %q`, funcName, test.start, test.end, test.sub, err)
		}

		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.base.Output, string(v), string(test.expected[idx]))
			}
		}
	}
}

func TestInsert(t *testing.T) {
	var funcName string = `TestInsert`
	tests := []struct {
		base     *StringBuilder
		start    int
		sub      *StringBuilder
		expected []rune
		ok       bool
	}{
		{
			base:     NewStringBuilder(`Hello world`),
			start:    2,
			sub:      NewStringBuilder(`yya`),
			expected: []rune(`Heyya world`),
			ok:       true,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    -2,
			sub:      NewStringBuilder(`yya`),
			expected: []rune(`Hello world`),
			ok:       false,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    10,
			sub:      NewStringBuilder(`yya`),
			expected: []rune(`Hello worlyya`),
			ok:       true,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    11,
			sub:      NewStringBuilder(`yya`),
			expected: []rune(`Hello worldyya`),
			ok:       true,
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    12,
			sub:      NewStringBuilder(`yya`),
			expected: []rune(`Hello world`),
			ok:       false,
		},
	}

	for _, test := range tests {
		if err := test.base.Insert(test.start, test.sub); err != nil && test.ok != false {
			t.Errorf(`%q(start: %v, sub: %q) has errored: %q`, funcName, test.start, test.sub, err)
		}
	}

}

func TestSubstring(t *testing.T) {
	var funcName string = `TestSubstring`
	tests := []struct {
		base     *StringBuilder
		start    int
		end      int
		expected []rune
	}{
		{
			base:     NewStringBuilder(`Hello world`),
			start:    2,
			end:      5,
			expected: []rune(`llo `),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    0,
			end:      400,
			expected: []rune(`Hello world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    5,
			end:      400,
			expected: []rune(` world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    -3,
			end:      400,
			expected: []rune(`Hello world`),
		},
	}

	for _, test := range tests {
		test.base.Substring(test.start, test.end)
		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.base.Output, string(v), string(test.expected[idx]))
			}
		}
	}

}

func TestGetSubstring(t *testing.T) {
	var funcName string = `TestGetSubstring`
	tests := []struct {
		base     *StringBuilder
		start    int
		end      int
		expected []rune
	}{
		{
			base:     NewStringBuilder(`Hello world`),
			start:    2,
			end:      5,
			expected: []rune(`llo `),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    0,
			end:      400,
			expected: []rune(`Hello world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    5,
			end:      400,
			expected: []rune(` world`),
		}, {
			base:     NewStringBuilder(`Hello world`),
			start:    -3,
			end:      400,
			expected: []rune(`Hello world`),
		},
	}

	for _, test := range tests {
		output := test.base.GetSubstring(test.start, test.end)
		for idx, v := range output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.base.Output, string(v), string(test.expected[idx]))
			}
		}
	}

}

func TestPadStart(t *testing.T) {
	var funcName string = `TestPadStart`
	tests := []struct {
		base     *StringBuilder
		length   int
		pad      interface{}
		expected []rune
		ok       bool
	}{
		{
			base:     NewStringBuilder(strconvNums),
			length:   12,
			pad:      0,
			expected: []rune(`000123456789`),
			ok:       true,
		}, {
			base:     NewStringBuilder(strconvNums),
			length:   15,
			pad:      `A`,
			expected: []rune(`AAAAAA123456789`),
			ok:       true,
		}, {
			base:     NewStringBuilder(strconvNums),
			length:   7,
			pad:      0,
			expected: []rune(strconvNums),
			ok:       false,
		}, {
			base:     NewStringBuilder(strconvNums),
			length:   12,
			pad:      []rune(`X`),
			expected: []rune(`XXX123456789`),
			ok:       true,
		}, {
			base:     NewStringBuilder(strconvNums),
			length:   12,
			pad:      'Y',
			expected: []rune(`YYY123456789`),
			ok:       true,
		}, {
			base:     NewStringBuilder(strconvNums),
			length:   12,
			pad:      NewStringBuilder(`#`),
			expected: []rune(`###123456789`),
			ok:       true,
		}, {
			base:     NewStringBuilder(strconvNums),
			length:   12,
			pad:      1.23456,
			expected: []rune(`123456789`),
			ok:       false,
		}}

	for _, test := range tests {
		if err := test.base.PadStart(test.length, test.pad); err != nil && test.ok != false {
			t.Errorf(`%q(%q) = %q`, funcName, test.base.Output, err)
		}

		for idx, v := range test.base.Output {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.base.Output, string(v), string(test.expected[idx]))
			}
		}
	}

}

func TestFields(t *testing.T) {
	var funcName string = `TestFields`
	tests := []struct {
		input    *StringBuilder
		expected [][]rune
		count    int
	}{
		{
			input: NewStringBuilder(`AAA BBB CCC`),
			expected: [][]rune{
				[]rune(`AAA`),
				[]rune(`BBB`),
				[]rune(`CCC`),
			},
			count: 3,
		}, {
			input: NewStringBuilder("AAA\tBBB\tCCC"),
			expected: [][]rune{
				[]rune(`AAA`),
				[]rune(`BBB`),
				[]rune(`CCC`),
			},
			count: 3,
		}, {
			input: NewStringBuilder("A A A\tBBB\tCCC"),
			expected: [][]rune{
				[]rune(`A`),
				[]rune(`A`),
				[]rune(`A`),
				[]rune(`BBB`),
				[]rune(`CCC`),
			},
			count: 5,
		},
	}

	for _, test := range tests {
		n, fields := test.input.Fields()

		if n != test.count {
			t.Errorf(`%q(%q) = %v fields, expected %v`, funcName, test.input.String(), n, test.count)
		}

		for a := 0; a < len(fields); a++ {
			for b := 0; b < len(fields[a]); b++ {
				if fields[a][b] != test.expected[a][b] {
					t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.input.String(), fields[a], test.expected[a])
				}
			}
		}
	}
}

func TestFieldsBy(t *testing.T) {
	var funcName string = `TestFieldsBy`
	tests := []struct {
		input     *StringBuilder
		separator rune
		expected  [][]rune
		count     int
	}{
		{
			input:     NewStringBuilder(`AAAxBBBxCCC`),
			separator: 'x',
			expected: [][]rune{
				[]rune(`AAA`),
				[]rune(`BBB`),
				[]rune(`CCC`),
			},
			count: 3,
		}, {
			input:     NewStringBuilder("AxAxAxBBBxCCC"),
			separator: 'x',
			expected: [][]rune{
				[]rune(`A`),
				[]rune(`A`),
				[]rune(`A`),
				[]rune(`BBB`),
				[]rune(`CCC`),
			},
			count: 5,
		},
	}

	for _, test := range tests {
		n, fields := test.input.FieldsBy(test.separator)

		if n != test.count {
			t.Errorf(`%q(%q) = %v fields, expected %v`, funcName, test.input.String(), n, test.count)
		}

		for a := 0; a < len(fields); a++ {
			for b := 0; b < len(fields[a]); b++ {
				if fields[a][b] != test.expected[a][b] {
					t.Errorf(`%q(%q) = %q, expected %q`, funcName, test.input.String(), fields[a], test.expected[a])
				}
			}
		}
	}
}

func TestFieldsRows(t *testing.T) {
	var funcName string = `TestFieldsRows`
	tests := []struct {
		input    *StringBuilder
		sep1     rune
		sep2     rune
		expected [][][]rune
		len1     int
		len2     int
	}{
		{
			input: NewStringBuilder(`+AxAx+BxBx+CxCx`),
			sep1:  '+',
			sep2:  'x',
			expected: [][][]rune{
				[][]rune{
					[]rune(`A`),
					[]rune(`A`),
				},
				[][]rune{
					[]rune(`B`),
					[]rune(`B`),
				},
				[][]rune{
					[]rune(`C`),
					[]rune(`C`),
				},
			},
			len1: 3,
			len2: 2,
		}, {
			input: NewStringBuilder(`+AAAxAAx+BBxBBBx+CxCCCx+DDDDDxDDx`),
			sep1:  '+',
			sep2:  'x',
			expected: [][][]rune{
				[][]rune{
					[]rune(`AAA`),
					[]rune(`AA`),
				},
				[][]rune{
					[]rune(`BB`),
					[]rune(`BBB`),
				},
				[][]rune{
					[]rune(`C`),
					[]rune(`CCC`),
				},
				[][]rune{
					[]rune(`DDDDD`),
					[]rune(`DD`),
				},
			},
			len1: 4,
			len2: 2,
		},
	}

	for _, test := range tests {
		r, f, rows := test.input.FieldsRows(test.sep1, test.sep2)

		if r != test.len1 {
			t.Errorf("%q(%q) = %v rows, expected %v", funcName, test.input.String(), r, test.len1)
		}

		if f != test.len2 {
			t.Errorf("%q(%q) = %v fields, expected %v", funcName, test.input.String(), r, test.len2)
		}

		for a := 0; a < len(rows); a++ {
			for b := 0; b < len(rows[a]); b++ {
				for c := 0; c < len(rows[a][b]); c++ {
					if rows[a][b][c] != test.expected[a][b][c] {
						t.Errorf("%q(%q) = %q, expected %q", funcName, test.input.String(), rows[a][b], test.expected[a][b])
					}
				}
			}
		}
	}
}

func TestString(t *testing.T) {
	var funcName string = `TestString`
	tests := []struct {
		input    *StringBuilder
		expected []rune
	}{
		{
			input:    NewStringBuilder(`Hello world`),
			expected: []rune(`Hello world`),
		}, {
			input:    NewStringBuilder(strconvNums),
			expected: []rune(strconvNums),
		}, {
			input:    NewStringBuilder(cttThree),
			expected: []rune(cttThree),
		}, {
			input: func() *StringBuilder {
				b := NewStringBuilder(nil)                  // ""
				b.Append(base)                              // "abc"
				b.Append(cttOne)                            // "abcdef"
				b.Append(cttTwo)                            // "abcdefghi"
				b.Delete(7, 9)                              // "abcdefg"
				b.Replace(4, 6, NewStringBuilder(cttThree)) // "abcd日本語"
				return b
			}(),
			expected: []rune(`abcd日本語`),
		},
	}
	for _, test := range tests {
		out := []rune(test.input.String())

		for idx, v := range out {
			if v != test.expected[idx] {
				t.Errorf(`%q(%q) = %q, expected %q`, funcName, out, string(v), string(test.expected[idx]))
			}
		}
	}
}
