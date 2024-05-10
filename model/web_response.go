package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	total  int         `json:total,omitempty`
	Data   interface{} `json:"data"`
}

func ResponseToClient(code int, status string, data interface{}) WebResponse {
	return WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
