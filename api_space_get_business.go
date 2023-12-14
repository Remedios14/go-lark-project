package lark_project

import (
	"context"
	"net/http"
)

// GetSpaceBusinessAll 获取空间下业务线详情
// docs: https://project.feishu.cn/openapp/help/articles/139893277257
func (r *SpaceService) GetSpaceBusinessAll(ctx context.Context, request *GetSpaceBusinessReq) (*GetSpaceBusinessResp, *Response, error) {
	var (
		scope    = "Space"
		funcName = "GetSpaceBusinessAll"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodGet,
			URL:             r.cli.OpenBaseURL + "/open_api/:project_key/business/all",
			UserKey:         request.UserKey,
			Body:            request,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	var resp = new(GetSpaceBusinessResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type GetSpaceBusinessReq struct {
	ProjectKey string `path:"project_key"` // ProjectKey 或 SimpleName
	UserKey    string
}

type GetSpaceBusinessResp struct {
	Data []*SpaceBusinessItem `json:"data,omitempty"`
	ErrorData
}

type SpaceBusinessItem struct {
	ID           string                             `json:"id,omitempty"`
	Name         string                             `json:"name,omitempty"`
	RoleOwners   map[string]*SpaceBusinessRoleOwner `json:"role_owners,omitempty"` // 默认角色及负责人，Map 的 key 为角色标识（ RoleOwner.role ）
	Watchers     []string                           `json:"watchers,omitempty"`    // UserKey 列表
	LevelID      int                                `json:"level_id,omitempty"`    // 层级 id，最顶层为 1，子节点依次递增
	Parent       string                             `json:"parent,omitempty"`      // ProjectKey
	Disabled     bool                               `json:"disabled,omitempty"`
	Labels       []string                           `json:"labels,omitempty"`
	Order        float32                            `json:"order,omitempty"`
	Project      string                             `json:"project,omitempty"`       // ProjectKey
	SuperMasters []string                           `json:"super_masters,omitempty"` // UserKey 列表
	Children     []*SpaceBusinessItem               `json:"children,omitempty"`      // 子业务线
}

type SpaceBusinessRoleOwner struct {
	Role   string   `json:"role,omitempty"`
	Name   string   `json:"name,omitempty"`
	Owners []string `json:"owners,omitempty"` // UserKey 列表
}
