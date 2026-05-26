package types

// NewResultAcknowledgement returns a new instance of Acknowledgement using an Acknowledgement_Result
// type in the Response field.
func NewResultAcknowledgement(result []byte) Acknowledgement {
	_ = "STUB: not implemented"
	return *new(Acknowledgement)
}

// NewErrorAcknowledgement returns a new instance of Acknowledgement using an Acknowledgement_Error
// type in the Response field.
// NOTE: Acknowledgements are written into state and thus, changes made to error strings included in packet acknowledgements
// risk an app hash divergence when nodes in a network are running different patch versions of software.
func NewErrorAcknowledgement(err string) Acknowledgement {
	_ = "STUB: not implemented"
	return *new(Acknowledgement)
}

// ValidateBasic performs a basic validation of the acknowledgement
func (ack Acknowledgement) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Success implements the Acknowledgement interface. The acknowledgement is
// considered successful if it is a ResultAcknowledgement. Otherwise it is
// considered a failed acknowledgement.
func (ack Acknowledgement) Success() bool { _ = "STUB: not implemented"; return false }

// Acknowledgement implements the Acknowledgement interface. It returns the
// acknowledgement serialised using JSON.
func (ack Acknowledgement) Acknowledgement() []byte { _ = "STUB: not implemented"; return nil }
