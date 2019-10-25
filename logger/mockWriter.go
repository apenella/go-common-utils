package log

// MockWrite implements io.Writer and is used to test messageOutput
type MockWriter struct {
	buff []byte
}
// Writes implementation
func (m *MockWriter) Write(p []byte) (n int, err error) {
	m.buff = p

	return len(p), nil
}
// Read implementation
func (m *MockWriter) Read(p []byte) (n int, err error) {
	p = m.buff[0:len(p)-1]

	return len(p), nil
}