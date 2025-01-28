package modelmocks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xorima/hub-sphere/internal/model"
)

func TestNewMockPullRequest(t *testing.T) {
	t.Run("it should implement the Pull Request interface", func(t *testing.T) {
		mockPR := NewMockPullRequest("title", "htmlUrl")
		assert.Implements(t, (*model.PullRequest)(nil), mockPR)
	})
}

func TestGetHTMLURL(t *testing.T) {
	t.Run("it should return the html url", func(t *testing.T) {
		mockPR := NewMockPullRequest("title", t.Name())
		assert.Equal(t, t.Name(), mockPR.GetHTMLURL())
	})
}

func TestMockPullRequestGetTitle(t *testing.T) {
	t.Run("it should return the title", func(t *testing.T) {
		mockPR := NewMockPullRequest(t.Name(), "htmlUrl")
		assert.Equal(t, t.Name(), mockPR.GetTitle())

	})
}
