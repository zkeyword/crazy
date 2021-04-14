package response

// Response response model
type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// WithData set model success and data
func (res *Response) WithData(data interface{}) *Response {
	res.Code = "200"
	res.Msg = "success"
	res.Data = data
	return res
}

// WithError set error message
func (res *Response) WithError(errCode string, Message string) *Response {
	res.Code = errCode
	res.Msg = Message
	return res
}
