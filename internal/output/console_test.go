package output

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xorima/hub-sphere/internal/model"
)

var (
	pr1 = []string{"this is my first pr https://link.com", "this is my next pr https://google.com"}
	pr2 = []string{"this is my second pr https://link.com", "this is my next pr https://github.com"}
)

type mockWriter struct {
	err error
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), m.err
}

func TestNewConsoleOutput(t *testing.T) {
	t.Run("it should create a new console outputter", func(t *testing.T) {
		o := NewConsoleOutput()
		assert.NotNil(t, o)
	})
}

func TestConsoleOutputWrite(t *testing.T) {
	t.Run("it should output with a new line and tabbed in entries to the console stdout without error", func(t *testing.T) {
		var entries = make(model.Entries)
		entries["my-primary-repo"] = pr1
		entries["my-seconard-repo"] = pr2
		o := NewConsoleOutput()
		assert.NoError(t, o.Write(entries))
	})
	t.Run("it should output with a new line and tabbed in entries", func(t *testing.T) {
		var writer = &strings.Builder{}
		var entries = make(model.Entries)
		entries["my-main-repo"] = pr1
		entries["my-other-repo"] = pr2
		o := NewConsoleOutput(withWriter(writer))
		assert.NoError(t, o.Write(entries))
		want := `
my-main-repo:
	this is my first pr https://link.com
	this is my next pr https://google.com
my-other-repo:
	this is my second pr https://link.com
	this is my next pr https://github.com`
		assert.Equal(t, want, writer.String())
	})
	t.Run("it should error if the outputter returns an error", func(t *testing.T) {
		var entries = make(model.Entries)
		entries["hello"] = pr1
		entries["world"] = pr2
		o := NewConsoleOutput(withWriter(&mockWriter{err: assert.AnError}))
		assert.ErrorIs(t, o.Write(entries), assert.AnError)

	})
}
