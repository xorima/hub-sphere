package output

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/xorima/hub-sphere/internal/model"
)

type ConsoleOutput struct {
	writer io.Writer
}

type ConsoleOptsFunc = func(o *ConsoleOutput)

func withWriter(w io.Writer) ConsoleOptsFunc {
	return func(o *ConsoleOutput) {
		o.writer = w
	}
}

func NewConsoleOutput(opts ...ConsoleOptsFunc) *ConsoleOutput {
	o := &ConsoleOutput{writer: os.Stdout}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *ConsoleOutput) Write(entries model.Entries) error {
	var keys []string
	for k := range entries {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, parent := range keys {
		_, err := fmt.Fprintf(o.writer, "\n%s:\n\t%s", parent, strings.Join(entries[parent], "\n\t"))
		if err != nil {
			return err
		}
	}
	return nil
}
