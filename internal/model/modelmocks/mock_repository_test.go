package modelmocks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xorima/hub-sphere/internal/model"
)

func TestNewMockRepository(t *testing.T) {
	t.Run("it should implement the repository interface", func(t *testing.T) {
		mockRepo := NewMockRepository("repoName")
		assert.Implements(t, (*model.Repository)(nil), mockRepo)
	})
}

func TestMockRepositoryGetName(t *testing.T) {
	t.Run("it should return the name", func(t *testing.T) {
		mockRepo := NewMockRepository(t.Name())
		assert.Equal(t, t.Name(), mockRepo.GetName())
	})
}
