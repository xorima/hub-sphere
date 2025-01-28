package modelmocks

type MockRepository struct {
	name string
}

func NewMockRepository(name string) *MockRepository {
	return &MockRepository{name: name}
}

func (m *MockRepository) GetName() string {
	return m.name
}
