// Package charmatcher will contain methods to compare two
// strings to eachother, on different levels
package charmatcher

import (
	"errors"
	"fmt"
	"strconv"
)

//
//
//          N777777777NO
//        N7777777777777N
//       M777777777777777N
//       $N877777777D77777M
//      N M77777777ONND777M
//      MN777777777NN  D777
//    N7ZN777777777NN ~M7778
//   N777777777777MMNN88777N
//   N777777777777MNZZZ7777O
//   DZN7777O77777777777777
//    N7OONND7777777D77777N
//     8$M++++?N???$77777$
//      M7++++N+M77777777N
//       N77O777777777777$                              M
//         DNNM$$$$777777N                              D
//        N$N:=N$777N7777M                             NZ
//       77Z::::N777777777                          ODZZZ
//      77N::::::N77777777M                         NNZZZ$
//    $777:::::::77777777MN                        ZM8ZZZZZ
//    777M::::::Z7777777Z77                        N++ZZZZNN
//   7777M:::::M7777777$777M                       $++IZZZZM
//  M777$:::::N777777$M7777M                       +++++ZZZDN
//    NN$::::::7777$$M777777N                      N+++ZZZZNZ
//      N::::::N:7$O:77777777                      N++++ZZZZN
//      M::::::::::::N77777777+                   +?+++++ZZZM
//      8::::::::::::D77777777M                    O+++++ZZ
//       ::::::::::::M777777777N                      O+?D
//       M:::::::::::M77777777778                     77=
//       D=::::::::::N7777777777N                    777
//      INN===::::::=77777777777N                  I777N
//     ?777N========N7777777777787M               N7777
//     77777$D======N77777777777N777N?         N777777
//    I77777$$$N7===M$$77777777$77777777$MMZ77777777N
//     $$$$$$$$$$$NIZN$$$$$$$$$M$$7777777777777777ON
//      M$$$$$$$$M    M$$$$$$$$N=N$$$$7777777$$$ND
//     O77Z$$$$$$$     M$$$$$$$$MNI==$DNNNNM=~N
//  7 :N MNN$$$$M$      $$$777$8      8D8I
//    NMM.:7O           777777778
//                      7777777MN
//                      M NO .7:
//                      M   :   M
//                           8
//

// CharMatcher struct represents a Character Matcher object;
// used for comparison operations between slices of runes
type CharMatcher struct {
	Chars []rune
	Input string
}

// NewCharMatcher will create a character matcher object based on
// the supplied input. Its methods will handle comparisons to other
// strings and CharMatchers
func NewCharMatcher(input interface{}) (*CharMatcher, error) {
	switch v := input.(type) {
	case string:
		return &CharMatcher{
			Chars: []rune(v),
			Input: v,
		}, nil
	case []rune:
		return &CharMatcher{
			Chars: v,
			Input: string(v),
		}, nil
	default:
		return nil, errors.New(`Couldn't match input type`)
	}

}

// Matches method will compare two strings ensuring they match
// in length and exact content
func (c *CharMatcher) Matches(input interface{}) (bool, error) {
	var inputRune []rune

	switch v := input.(type) {
	case string:
		inputRune = []rune(v)
	case []rune:
		inputRune = v
	case *CharMatcher:
		inputRune = v.Chars
	default:
		return false, errors.New(`Couldn't match input type`)
	}

	if len(inputRune) != len(c.Chars) {
		return false, nil
	}

	for idx, v := range c.Chars {
		if v != inputRune[idx] {
			return false, nil
		}
	}
	return true, nil
}

// MatchesAllOf method will return a boolean whether the input
// is present in the CharMatcher string (regardless where),
// provided it's present in its entirety
func (c *CharMatcher) MatchesAllOf(input interface{}) bool {
	var inputRune []rune

	switch v := input.(type) {
	case string:
		inputRune = []rune(v)
	case []rune:
		inputRune = v
	case *CharMatcher:
		inputRune = v.Chars
	default:
		panic(fmt.Errorf(`Couldn't match input type for %q`, input))
	}

	var spree bool = false
	var counter int = 0

	for idx, v := range inputRune {

		for _, t := range c.Chars {
			if t == v && idx == counter {
				spree = true
				counter++
				break
			}
			if spree == true && counter < idx {
				return false
			}
		}
		if spree != true {
			return false
		}

	}
	return spree
}

// ContainsAnyOf method will check whether any of the input's
// individual characters are present in the CharMatcher
func (c *CharMatcher) ContainsAnyOf(input string) bool {
	inputRune := []rune(input)

	for _, cVal := range c.Chars {
		for _, iVal := range inputRune {
			if cVal == iVal {
				return true
			}
		}
	}

	return false
}

// Contains method will return a boolean whether the input
// is present in the CharMatcher string (regardless where),
// provided it's present in its entirety. Similar implementation
// of MatchesAllOf method; and only accepts a string as input
func (c *CharMatcher) Contains(input string) bool {
	inputRune := []rune(input)
	var spree bool = false

	for idx, cVal := range c.Chars {
		if cVal == inputRune[0] {
			for i := 0; i < len(inputRune); i++ {
				if inputRune[i] == c.Chars[idx] {
					spree = true
				} else {
					spree = false
					return spree
				}
				idx++

			}
		}
	}
	return spree
}

// Index method returns a boolean on whether the input character
// (rune) and index (int) provided matches the character in the
// same index, in the CharMatcher object
func (c *CharMatcher) Index(input rune, index int) bool {
	if c.Chars[index] == input {
		return true
	}
	return false
}

// IndexOf method returns an int and boolean on whether the input
// character (rune) provided matches a similar character in the
// CharMatcher object; returns -1, false in case of no matches
func (c *CharMatcher) IndexOf(input rune) (int, bool) {
	for idx, v := range c.Chars {
		if v == input {
			return idx, true
		}
	}
	return -1, false
}

// IndexesOf method returns multiple matches of an input character
// (rune) in the CharMatcher object, in a slice of ints containing their
// indexes -- or a slice of ints containing -1 if there's no match
func (c *CharMatcher) IndexesOf(input rune) []int {
	var matched []int

	for idx, v := range c.Chars {
		if v == input {
			matched = append(matched, idx)
		}
	}
	if len(matched) == 0 {
		matched = append(matched, -1)
	}
	return matched
}

// CompareTo method will return the difference in numeric values between
// two strings;
// comparing `abc` to `abd` = 1
func (c *CharMatcher) CompareTo(input string, radix int) int {
	inputNumID, err := strconv.ParseInt(input, radix, 64)

	if err != nil {
		panic(err)
	}

	selfNumID, err := strconv.ParseInt(string(c.Chars), radix, 64)

	if err != nil {
		panic(err)
	}

	return int(inputNumID) - int(selfNumID)

}
