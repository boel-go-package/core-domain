package message

type Message struct {
	StatusCode string      `json:"status_code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta"`
}

func ReturnMessage(status_code string, status string, message string, data interface{}, meta interface{}) *Message {
	return &Message{
		StatusCode: status_code,
		Status:     status,
		Message:    message,
		Data:       data,
		Meta:       meta,
	}
}

func Success(status_code string, message string, data interface{}, meta interface{}) *Message {
	return ReturnMessage(status_code, "success", message, data, meta)
}

func Failed(status_code string, message string, data interface{}, meta interface{}) *Message {
	return ReturnMessage(status_code, "failed", message, data, meta)
}

func Custom(status_code string, status string, message string, data interface{}, meta interface{}) *Message {
	return ReturnMessage(status_code, status, message, data, meta)
}
