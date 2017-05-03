package codegen

import (
	"bytes"
	"strings"
)

type Writer struct {
	b      bytes.Buffer
	indent int
}

func (w *Writer) Reset() {
	w.b.Reset()
}

func (w *Writer) Line(parts ...string) {
	text := strings.Join(parts, "")

	w.b.WriteString(strings.Repeat("\t", w.indent))
	w.b.WriteString(text)
	w.b.WriteString("\n")
}

func (w *Writer) Comment(text string) {
	// TODO break up long lines
	w.Line("// " + text)
}

func (w *Writer) In() {
	w.indent += 1
}

func (w *Writer) Out() {
	w.indent -= 1
}

func (w *Writer) String() (out string) {
	return w.b.String()
}
