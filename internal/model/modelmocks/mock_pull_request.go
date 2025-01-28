package modelmocks

type MockPullRequest struct {
	title   string
	htmlUrl string
}

func NewMockPullRequest(title, htmlUrl string) *MockPullRequest {
	return &MockPullRequest{title: title, htmlUrl: htmlUrl}
}

func (m *MockPullRequest) GetTitle() string {
	return m.title
}

func (m *MockPullRequest) GetHTMLURL() string {
	return m.htmlUrl
}
