package lark_project

type ErrorData struct {
	Err     map[string]interface{} `json:"error,omitempty"`
	ErrCode int64                  `json:"err_code,omitempty"` // 错误码，非 0 表示失败
	ErrMsg  string                 `json:"err_msg,omitempty"`  // 错误描述
}
