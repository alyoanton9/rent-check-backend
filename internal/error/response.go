package error

type UniqueErrorResponse struct {
	Msg   string `json:"msg"`
	Field string `json:"field"`
}
