// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

type MyString string

func (s MyString) Read(p []byte) (int, error) {
	smaller := len(s)
	if len(p) < smaller {
		smaller = len(p)
	}
	for i := 0; i < smaller; i++ {
		p[i] = s[i]
	}
	return smaller, nil
}

// Reverse returns the given string in the reverse order.
func Reverse(input string) string {
	values := []rune(input)
	l := len(input)
	for i := 0; i < l/2; i++ {
		values[i], values[l-1-i] = values[l-1-i], values[i]
	}
	return string(values)
}
