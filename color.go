package color

import (
	"fmt"
	"io"
)

// Color escaspe code.
const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Writer is used to write colored text.
type Writer struct {
	W     io.Writer
	Color int
}

// New returns color writer.
func New(w io.Writer, color int) io.Writer {
	return &Writer{W: w, Color: color}
}

func (c *Writer) Write(b []byte) (int, error) {
	var n int
	if _, err := io.WriteString(c.W, fmt.Sprintf("\033[%vm", c.Color)); err != nil {
		return n, err
	}
	n0, err := c.W.Write(b)
	n += n0
	if err != nil {
		return n, err
	}
	if _, err := io.WriteString(c.W, "\033[0m"); err != nil {
		return n, err
	}
	return len(b), nil
}

func Wrap(s string, color int) string {
	return fmt.Sprintf("\033[%vm%s\033[0m", color, s)
}
