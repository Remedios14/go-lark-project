package lark_project

import (
	"context"
	"net/http"
)

// GetSpaceDetail 获取空间详情
//
// 注意事项：
// - 获取到的 Data 总总是以 ProjectKey 为 key，即使请求时使用的是 SimpleName
// - 获取到的 Administrators 是 UserKey 的列表
//
// docs: https://project.feishu.cn/openapp/help/articles/867915863696
func (r *SpaceService) GetSpaceDetail(ctx context.Context, request *GetSpaceInfoReq) (*GetSpaceInfoResp, *Response, error) {
	var (
		scope    = "Space"
		funcName = "GetSpaceDetail"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodPost,
			URL:             r.cli.OpenBaseURL + "/open_api/projects/detail",
			Body:            request,
			UserKey:         request.UserKey,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	var resp = new(GetSpaceInfoResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type GetSpaceInfoReq struct {
	UserKey     string   `json:"user_key"`
	ProjectKeys []string `json:"project_keys,omitempty"`
	SimpleNames []string `json:"simple_names,omitempty"`
}

type GetSpaceInfoResp struct {
	Data map[string]*SpaceInfoItem `json:"data,omitempty"` // key: ProjectKey
	ErrorData
}

type SpaceInfoItem struct {
	ProjectKey     string   `json:"project_key,omitempty"`
	Name           string   `json:"name,omitempty"`
	SimpleName     string   `json:"simple_name,omitempty"`
	Administrators []string `json:"administrators,omitempty"` // UserKey 列表
}
