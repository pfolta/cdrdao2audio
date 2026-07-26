// Copyright (c) 2026 Peter Folta
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package formatter

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Format string

const (
	JSON Format = "json"
	TEXT Format = "text"
)

var ErrUnknownFormat = errors.New("unknown format")

// Formatter writes values to an [io.Writer].
type Formatter interface {
	// Write writes v to w using the formatter's output format.
	Write(w io.Writer, v any) error
}

// NewFormatter creates a formatter for the requested format.
func NewFormatter(format Format) (Formatter, error) {
	switch Format(strings.ToLower(string(format))) {
	case JSON:
		return NewJSONFormatter(), nil

	case TEXT:
		return NewTextFormatter(), nil

	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownFormat, format)
	}
}

// Formats returns the list of supported output formats.
func Formats() []string {
	return []string{
		string(TEXT),
		string(JSON),
	}
}
