package Utils

type ErrorResponse struct {
	Message   string      `json:"message"`
	ErrorCode int         `json:"errorCode"`
	Result    interface{} `json:"result"`
	TargetUrl interface{} `json:"targetUrl"`
	Success   bool        `json:"success"`
	Error     struct {
		Code             int         `json:"code"`
		Message          string      `json:"message"`
		Details          string      `json:"details"`
		ValidationErrors interface{} `json:"validationErrors"`
	} `json:"error"`
	UnAuthorizedRequest bool `json:"unAuthorizedRequest"`
	Abp                 bool `json:"__abp"`
}
