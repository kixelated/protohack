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

func (w *Writer) Line(parts ...string) *Writer {
	text := strings.Join(parts, "")

	w.b.WriteString(strings.Repeat("\t", w.indent))
	w.b.WriteString(text)
	w.b.WriteString("\n")

	return w
}

func (w *Writer) Comment(text string) *Writer {
	// TODO break up long lines
	w.Line("// " + text)
	return w
}

func (w *Writer) In() *Writer {
	w.indent += 1
	return w
}

func (w *Writer) Out() *Writer {
	w.indent -= 1
	return w
}

func (w *Writer) String() (out string) {
	return w.b.String()
}
