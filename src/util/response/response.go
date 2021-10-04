package response

type response struct {
	MessageType string      `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Message     string      `json:"message,omitempty" xml:"message,omitempty"`
	Data        interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func New(
	messageType,
	message string,
	data interface{},
) *response {
	return &response{messageType, message, data}
}

func NewErr(
	message string,
	data interface{},
) *response {
	return &response{"error", message, data}
}

func NewSatisfactory(
	message string,
	data interface{},
) *response {
	return &response{"message", message, data}
}
