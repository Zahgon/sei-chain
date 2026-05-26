package mock

// MockEmptyAcknowledgement implements the exported.Acknowledgement interface and always returns an empty byte string as Response
type MockEmptyAcknowledgement struct {
	Response []byte
}

// NewMockEmptyAcknowledgement returns a new instance of MockEmptyAcknowledgement
func NewMockEmptyAcknowledgement() MockEmptyAcknowledgement {
	_ = "STUB: not implemented"
	return *new(MockEmptyAcknowledgement)
}

// Success implements the Acknowledgement interface
func (ack MockEmptyAcknowledgement) Success() bool {
	_ = "STUB: not implemented"

	// Acknowledgement implements the Acknowledgement interface
	return false
}

func (ack MockEmptyAcknowledgement) Acknowledgement() []byte { _ = "STUB: not implemented"; return nil }
