package responses

type HandleResponse struct {
	Message string      `json:"mensage"`
	Data    interface{} `json:"data,omitempty"`
}

type msgsHandler struct {
	FailureBindJSON string
}

var MsgHandler = &msgsHandler{
	FailureBindJSON: "Invalid JSON format",
}
