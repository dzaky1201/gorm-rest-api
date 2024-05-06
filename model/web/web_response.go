package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func JsonResponse(code int, message string, data interface{}) WebResponse {
	return WebResponse{
		Code:   code,
		Status: message,
		Data:   data,
	}
}
