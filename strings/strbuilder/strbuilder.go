// Package strbuilder serves as a helper to build strings; containing helpful
// methods to facilitate standard actions
package strbuilder

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	cm "github.com/ZalgoNoise/meta/strings/charmatcher"
)

const (
	defaultCapacity  int   = 16
)

// StringBuilder struct represents an object to build strings,
// referring to multiple (private) methods to achieve it; and
// focusing on slices of runes for wider parsing
type StringBuilder struct {
	Output   []rune
	Capacity int
}

// NewStringBuilder function creates a new StringBuilder object
// based on multiple types of input
func NewStringBuilder(input interface{}) *StringBuilder {
	switch v := input.(type) {
	case int:
		return &StringBuilder{
			Output:   []rune{},
			Capacity: v,
		}
	case string:
		return &StringBuilder{
			Output:   []rune(v),
			Capacity: len([]rune(v)) + defaultCapacity,
		}
	case []rune:
		return &StringBuilder{
			Output:   v,
			Capacity: len(v) + defaultCapacity,
		}
	default:
		return &StringBuilder{
			Output:   []rune{},
			Capacity: defaultCapacity,
		}
	}
}

// CompareTo method will take an input StringBuilder object and use
// a CharMatcher to compare the content of both
func (b *StringBuilder) CompareTo(another *StringBuilder) (bool, error) {
	matcher, err := cm.NewCharMatcher(b.Output)

	if err != nil {
		return false, err
	}

	return matcher.Matches(another.Output)
}

// Append method will add more content to the existing StringBuilder's Output
func (b *StringBuilder) Append(another interface{}) error {
	var add []rune
	switch v := another.(type) {
	case int:
		add = []rune(strconv.Itoa(v))

	case string:
		add = []rune(v)

	case []rune:
		add = v

	case *StringBuilder:
		add = v.Output

	default:
		return errors.New(`Couldn't match input type`)
	}

	b.Output = append(b.Output, add...)
	b.Capacity = b.Capacity + len(add)

	return nil
}

// Delete method removes a portion of the string by taking a starting and ending point,
// index-wise
func (b *StringBuilder) Delete(start, end int) {
	if start < 0 || end < 0 {
		return
	}
	count := len(b.Output)

	if end >= count {
		end = count
	} else {
		end = end + 1
	}

	head := b.Output[:start]
	tail := b.Output[end:]

	if len(tail) > 0 {
		b.Output = append(head, tail...)
		return
	}
	b.Output = head
	return

}

// DeleteAt method will remove one single character from the StringBuilder's Output;
// it is the equivalent to using Delete(start, end) where start and end are the same value
func (b *StringBuilder) DeleteAt(index int) {
	if index < 0 {
		return
	}
	count := len(b.Output)

	if index >= count {
		index = (count - 1)
	}

	head := b.Output[:index]
	tail := b.Output[(index + 1):]

	if len(tail) > 0 {
		b.Output = append(head, tail...)
		return
	}
	b.Output = head
	return
}

// Replace method will place in characters in a StringBuilder's Output, as per provided
// start and ending indexes. Returns an error in case the difference between the starting and
// ending points; and the length of the input string would result in an index out of bounds error
func (b *StringBuilder) Replace(start, end int, input *StringBuilder) error {
	if start < 0 || end < 0 {
		return fmt.Errorf("ERROR: Index out of bounds exception for the provided start / end indexes:\n"+
			" Start / end indexes cannot be negative numbers. Start: %v ; End: %v", start, end)

	}
	count := len(b.Output)

	if end > count {
		end = count
	}

	diff := math.Abs(float64((end - start) + 1))

	head := b.Output[:start]

	if int(diff) > len(input.Output) {
		// index out of bounds exception
		return errors.New(`ERROR: Index out of bounds exception for the provided input length: 
	Difference between replaced bit length (` + strconv.Itoa(int(diff)) + `) is greater than the replacement string's length (` + strconv.Itoa(len(input.Output)) + `)`)
	}

	var mid []rune
	for i := 0; i < int(diff); i++ {
		mid = append(mid, input.Output[i])
	}

	head = append(head, mid...)

	tail := b.Output[(end + 1):]

	if len(tail) > 0 {
		b.Output = append(head, tail...)
		return nil
	}
	b.Output = head
	return nil
}

// Insert method will place in a set of characters on the Output of a StringBuilder object
func (b *StringBuilder) Insert(offset int, input *StringBuilder) error {
	if offset < 0 {
		return fmt.Errorf("ERROR: Index out of bounds exception for the provided offset index:\n"+
			" Offset index cannot be a negative number. Offset: %v", offset)

	}
	count := len(b.Output)

	if int(offset) > count {
		// index out of bounds exception
		return errors.New(`ERROR: Index out of bounds exception for the provided offset: 
	Offset set too high (` + strconv.Itoa(offset) + `) for the actual length of the string (` + strconv.Itoa(len(input.Output)) + `)`)
	}

	head := b.Output[:offset]

	if (offset + len(input.Output)) > count {
		b.Output = append(head, input.Output...)
		return nil
	}

	head = append(head, input.Output...)
	tail := b.Output[(offset + len(input.Output)):]
	b.Output = append(head, tail...)
	return nil

}

// Substring method will reduce the Output value of a StringBuilder object based on
// the start and end indexes provided
func (b *StringBuilder) Substring(start, end int) {
	if start < 0 || end < 0 {
		return
	}

	count := len(b.Output)

	if end > count || end < 0 {
		end = count
	}

	b.Output = b.Output[start:end]
	return
}

// GetSubstring method is similar to the Substring() method, however it returns a
// string instead of defining the StringBuilder object's Output value
func (b *StringBuilder) GetSubstring(start, end int) []rune {
	if start < 0 || end < 0 {
		return b.Output
	}

	count := len(b.Output)

	if end > count || end < 0 {
		end = count
	}

	return b.Output[start:end]

}

// PadStart method will add padding to the beginning Output of a StringBuilder object, depending on
// the desired string length and the padding character provided.
func (b *StringBuilder) PadStart(length int, pad interface{}) error {

	var char []rune
	switch v := pad.(type) {
	case int:
		char = []rune(strconv.Itoa(v))

	case string:
		char = []rune(v)

	case []rune:
		char = v

	case rune:
		char = []rune{v}

	case *StringBuilder:
		char = v.Output

	default:
		return errors.New(`Couldn't match input type`)
	}

	if length < len(b.Output) {
		// index out of bounds exception
		return errors.New(`Padding is shorter than string's length`)
	} else if length == len(b.Output) {
		return nil
	}

	padAmount := length - len(b.Output)
	var padding []rune

	for i := 0; i < padAmount; i++ {
		padding = append(padding, char...)
	}

	b.Output = append(padding, b.Output...)

	return nil

}

// Fields method will breakdown the content in the StringBuilder object
// by its whitespace, returning the number of fields detected and a slice
// of byte arrays (for each field and its content)
func (b *StringBuilder) Fields() (n int, fields [][]byte) {

	var buf []byte
	counter := 0

	for i := 0; i < len(b.Output); i++ {

		// " " == 32
		// "\t" == 9
		if b.Output[i] == 32 || b.Output[i] == 9 {

			if len(buf) == 0 {
				continue
			}
			fields = append(fields, buf)
			buf = []byte{}
			counter++
			continue
		}
		buf = append(buf, byte(b.Output[i]))
	}
	if len(buf) > 0 {
		fields = append(fields, buf)
	}

	return len(fields), fields
}

// FieldsBy method will breakdown the content in the StringBuilder object
// by the input rune provided, returning the number of fields detected
// and a slice of byte arrays (for each field and its content)
func (b *StringBuilder) FieldsBy(sep rune) (n int, fields [][]byte) {

	var buf []byte
	counter := 0

	for i := 0; i < len(b.Output); i++ {

		// " " == 32
		// "\t" == 9
		if b.Output[i] == sep {

			if len(buf) == 0 {
				continue
			}
			fields = append(fields, buf)
			buf = []byte{}
			counter++
			continue
		}
		buf = append(buf, byte(b.Output[i]))
	}
	if len(buf) > 0 {
		fields = append(fields, buf)
	}

	return len(fields), fields
}

// FieldsRows method will breakdown the content in the StringBuilder object
// by the input runes provided, returning the number of rows and fields detected
// and a slice of slice of byte arrays (for each row, and each field and its content)
func (b *StringBuilder) FieldsRows(row, fld rune) (r, f int, rows [][][]byte) {

	var buf []byte
	var fields [][]byte

	for i := 0; i < len(b.Output); i++ {
		if b.Output[i] == row {
			if len(fields) == 0 {
				continue
			}
			rows = append(rows, fields)
			fields = [][]byte{}
			continue
		}
		// " " == 32
		// "\t" == 9
		// "\n" == 10
		if b.Output[i] == fld {

			if len(buf) == 0 {
				continue
			}
			fields = append(fields, buf)
			buf = []byte{}
			continue
		}
		buf = append(buf, byte(b.Output[i]))
	}

	if len(buf) > 0 {
		fields = append(fields, buf)
	}

	if len(fields) > 0 {
		rows = append(rows, fields)
	}

	return len(rows), len(fields), rows
}


// String method returns the Output value of a StringBuilder object, as a string type
func (b *StringBuilder) String() string {
	return string(b.Output)
}
