package lark_project

import (
	"context"
	"net/http"
)

// QueryUserDetail 获取用户详情
// docs: https://project.feishu.cn/openapp/help/articles/363737543290
func (r *UserService) QueryUserDetail(ctx context.Context, request *QueryUserDetailReq) (*QueryUserDetailResp, *Response, error) {
	var (
		scope    = "User"
		funcName = "QueryUserDetail"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodPost,
			URL:             r.cli.OpenBaseURL + "/open_api/user/query",
			Body:            request,
			NeedPluginToken: true,
		}
	)
	var resp = new(QueryUserDetailResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type QueryUserDetailReq struct {
	UserKeys []string `json:"user_keys,omitempty"`
	OutIDs   []string `json:"out_ids,omitempty"`
	Emails   []string `json:"emails,omitempty"`
}

type QueryUserDetailResp struct {
	Data []*UserDetail `json:"data,omitempty"`
	ErrorData
}

type UserDetail struct {
	UserKey   string         `json:"user_key,omitempty"`
	Username  string         `json:"username,omitempty"`
	Email     string         `json:"email,omitempty"`
	NameCn    string         `json:"name_cn,omitempty"`
	NameEn    string         `json:"name_en,omitempty"`
	AvatarURL string         `json:"avatar_url,omitempty"`
	OutID     string         `json:"out_id,omitempty"`
	Status    UserStatus     `json:"status,omitempty"`
	Channels  []*UserChannel `json:"channels,omitempty"`
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "activated"
	UserStatusFrozen   UserStatus = "frozen"
	UserStatusResigned UserStatus = "resigned"
)

type UserChannel struct {
	TenantGroupID int    `json:"tenant_group_id,omitempty"`
	TenantName    string `json:"tenant_name,omitempty"`
}
