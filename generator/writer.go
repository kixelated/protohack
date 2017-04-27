package generator

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"strings"
)

type writer struct {
	head bytes.Buffer
	body bytes.Buffer

	pkg     string
	imports map[string]string
	indent  int
}

func newWriter() (w *writer) {
	return &writer{
		imports: make(map[string]string),
	}
}

func (w *writer) Package(name string) {
	w.head.WriteString("package " + name + "\n")
}

func (w *writer) Import(ref string) (alias string) {
	alias, ok := w.imports[ref]
	if !ok {
		alias = path.Dir(ref)
		w.imports[ref] = alias

		w.head.WriteString("import \"" + ref + "\"\n")
	}

	return alias
}

func (w *writer) Line(line string, args ...interface{}) {
	line = fmt.Sprintf(line, args...)
	line = strings.Repeat("\t", w.indent) + line + "\n"
	w.body.WriteString(line)
}

func (w *writer) In() {
	w.indent += 1
}

func (w *writer) Out() {
	w.indent -= 1
}

func (w *writer) Output() (data string, err error) {
	var buf bytes.Buffer

	_, err = io.Copy(&buf, &w.head)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(&buf, &w.body)
	if err != nil {
		return "", err
	}

	data = buf.String()

	return data, nil
}
