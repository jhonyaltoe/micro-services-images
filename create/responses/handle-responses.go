package responses

type HandleResponse struct {
	Message string      `json:"mensage"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Index   string      `json:"index,omitempty"`
}

type msgsHandler struct {
	FailureBindJSON string
	ValidationErr   string
}

var MsgHandler = &msgsHandler{
	FailureBindJSON: "Invalid JSON format",
	ValidationErr:   "Validation error",
}
