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
	"io"

	"github.com/goccy/go-yaml"
)

func defaultYAMLOptions() []yaml.EncodeOption {
	return []yaml.EncodeOption{
		// Indent using 2 spaces
		yaml.Indent(2),
	}
}

// YAMLFormatter writes values as complete YAML documents.
type YAMLFormatter struct {
	opts []yaml.EncodeOption
}

// NewYAMLFormatter creates a formatter that writes YAML.
func NewYAMLFormatter(opts ...yaml.EncodeOption) *YAMLFormatter {
	return &YAMLFormatter{
		opts: append(defaultYAMLOptions(), opts...),
	}
}

// Write writes v as a YAML document beginning with "---" and
// ending with a newline.
func (f *YAMLFormatter) Write(w io.Writer, v any) error {
	if _, err := io.WriteString(w, "---\n"); err != nil {
		return err
	}

	enc := yaml.NewEncoder(w, f.opts...)
	defer enc.Close()

	return enc.Encode(v)
}
