package lark_project

import (
	"context"
	"net/http"
)

// GetSpaceTeams 获取项目下团队人员
// docs: https://project.feishu.cn/openapp/help/articles/838384798514
func (r *SpaceService) GetSpaceTeams(ctx context.Context, request *GetSpaceTeamsReq) (*GetSpaceTeamsResp, *Response, error) {
	var (
		scope    = "Space"
		funcName = "GetSpaceTeams"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodGet,
			URL:             r.cli.OpenBaseURL + "/open_api/:project_key/teams/all",
			UserKey:         request.UserKey,
			Body:            request,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	var resp = new(GetSpaceTeamsResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type GetSpaceTeamsReq struct {
	ProjectKey string `path:"project_key"` // ProjectKey 或 SimpleName
	UserKey    string
}

type GetSpaceTeamsResp struct {
	Data []*SpaceTeam `json:"data,omitempty"`
	ErrorData
}

type SpaceTeam struct {
	TeamID         int      `json:"team_id,omitempty"`
	TeamName       string   `json:"team_name,omitempty"`
	UserKeys       []string `json:"user_keys,omitempty"`      // UserKey 列表
	Administrators []string `json:"administrators,omitempty"` // UserKey 列表
}
