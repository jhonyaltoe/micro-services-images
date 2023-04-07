package responses

type DbResponse struct {
	Message string   `json:"mensage"`
	Data interface{} `json:"data.omitempty"`
}

type msgsDb struct {
	SuccessUpdate string
	FailureUpdate string
	InternalError string
}

var MsgDb = &msgsDb{
	SuccessUpdate: "Successfully updated",
	FailureUpdate: "No updates ocurred",
	InternalError: "Internal server error",
}
