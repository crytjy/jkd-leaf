package network

type Processor interface {
	// Route must goroutine safe
	Route(msgId uint16, msg interface{}, userData interface{}) error
	// Unmarshal must goroutine safe
	Unmarshal(data []byte) (interface{}, error, uint16)
	// Marshal must goroutine safe
	Marshal(msg interface{}) ([][]byte, error)
}
