package lark_project

import (
	"encoding/json"
	"fmt"
)

// Error ...
type Error struct {
	Scope    string
	FuncName string
	Code     int64
	Msg      string
}

// Error ...
func (r *Error) Error() string {
	if r.Code == 0 {
		return ""
	}
	return fmt.Sprintf("request %s#%s failed: code: %d, msg: %s", r.Scope, r.FuncName, r.Code, r.Msg)
}

// NewError ...
func NewError(scope, funcName string, code int64, msg string) error {
	return &Error{
		Scope:    scope,
		FuncName: funcName,
		Code:     code,
		Msg:      msg,
	}
}

// GetErrorCode ...
func GetErrorCode(err error) int64 {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.Code
		}
		return -1
	}
	return 0
}

func jsonString(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

// mask key
const (
	httpRequestHeaderAuthorization = "Authorization"
	httpRequestHeaderHelpdeskAuth  = "X-Lark-Helpdesk-Authorization"
	RpcLogIDKey                    = "K_LOGID"
	HttpHeaderLogIDKey             = "X-Tt-Logid"
	HttpHeaderPluginToken          = "X-PLUGIN-TOKEN"
	HttpHeaderUserKey              = "X-USER-KEY"
)

func jsonHeader(headers map[string]string) string {
	val := make(map[string]string, len(headers))
	for k, v := range headers {
		if k == httpRequestHeaderAuthorization || k == httpRequestHeaderHelpdeskAuth {
			val[k] = maskString(v, 9, '*') // `Bearer xx******`
		} else {
			val[k] = v
		}
	}
	bs, _ := json.Marshal(val)
	return string(bs)
}

func maskString(s string, prefixCount int, mask rune) string {
	ss := []rune(s)
	res := make([]rune, len(ss))
	for i, v := range ss {
		if i < prefixCount {
			res[i] = v
		} else {
			res[i] = mask
		}
	}
	return string(res)
}
