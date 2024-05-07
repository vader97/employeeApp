package customerrors

type RestErr struct {
	Message    string `json:"message"`
	ErrorCode  string `json:"errorCode"`
	StatusCode int    `json:"-"`
}
