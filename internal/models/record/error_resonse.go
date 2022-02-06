package record

type ErrorResponse struct {
	Code int    `json:"code" `
	Msg  string `json:"msg" `
}
