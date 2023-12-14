package lark_project

import (
	"context"
)

// Mock Mock struct
type Mock struct {
	mockRawRequest func(ctx context.Context, req *RawRequestReq, resp interface{}) (response *Response, err error)
}
