package lark_project

import (
	"context"
	"net/http"
)

// GetSpaceList 获取空间列表
// docs: https://project.feishu.cn/openapp/help/articles/071694862282
func (r *SpaceService) GetSpaceList(ctx context.Context, request *GetSpaceListReq) (*GetSpaceListResp, *Response, error) {
	var (
		scope    = "Space"
		funcName = "GetSpaceList"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodPost,
			URL:             r.cli.OpenBaseURL + "/open_api/projects",
			Body:            request,
			UserKey:         request.UserKey,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	resp := new(GetSpaceListResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type GetSpaceListReq struct {
	UserKey string `json:"user_key"`
}

type GetSpaceListResp struct {
	Data []string `json:"data"` // ProjectKey 列表
	ErrorData
}
