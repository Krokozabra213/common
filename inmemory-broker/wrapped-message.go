package inmemorybroker

type WrappedMessage struct {
	ID      uint64
	Status  uint8
	Message interface{}
}

func NewWrappedMessage(id uint64, status uint8, message interface{}) *WrappedMessage {
	wrappedMessage := &WrappedMessage{
		ID:      id,
		Status:  status,
		Message: message,
	}
	// TODO: Validation
	return wrappedMessage
}
