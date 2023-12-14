package lark_project

import (
	"context"
	"net/http"
)

// GetWorkItemAll 获取空间下工作项类型
// docs: https://project.feishu.cn/openapp/help/articles/925668155213
func (r *SpaceService) GetWorkItemAll(ctx context.Context, request *GetSpaceWorkItemReq) (*GetSpaceWorkItemResp, *Response, error) {
	var (
		scope    = "Space"
		funcName = "GetWorkItemAll"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodGet,
			URL:             r.cli.OpenBaseURL + "/open_api/:project_key/work_item/all-types",
			UserKey:         request.UserKey,
			Body:            request,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	var resp = new(GetSpaceWorkItemResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type GetSpaceWorkItemReq struct {
	ProjectKey string `path:"project_key"` // ProjectKey 或 SimpleName
	UserKey    string
}

type GetSpaceWorkItemResp struct {
	Data []*SpaceWorkItem `json:"data,omitempty"`
	ErrorData
}

type SpaceWorkItem struct {
	ApiName   string `json:"api_name,omitempty"`
	IsDisable int    `json:"is_disable,omitempty"`
	TypeKey   string `json:"type_key,omitempty"` // 工作项类型 Key
	Name      string `json:"name,omitempty"`     // 工作项类型名称
}
