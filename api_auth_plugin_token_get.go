package lark_project

import (
	"context"
	"time"
)

// GetPluginToken get plugin token
// docs: https://project.feishu.cn/openapp/help/articles/418746063091#d2dARb49NsDrOJVcAop9
func (r *AuthService) GetPluginToken(ctx context.Context) (*TokenExpire, *Response, error) {
	r.cli.log(ctx, LogLevelInfo, "[LarkProject] Auth#GetPluginToken call api")

	val, ttl, err := r.cli.store.Get(ctx, genPluginTokenKey(r.cli.pluginID))
	if err != nil && err != ErrStoreNotFound {
		r.cli.log(ctx, LogLevelError, "[LarkProject] Auth#GetPluginToken get token from store failed: %s", err)
	} else if val != "" && ttl > 0 {
		return &TokenExpire{Token: val, ExpireTime: int64(ttl.Seconds())}, &Response{}, nil
	}

	uri := r.cli.OpenBaseURL + "/bff/v2/authen/plugin_token"

	req := &RawRequestReq{
		Scope:  "Auth",
		API:    "GetPluginToken",
		Method: "POST",
		URL:    uri,
		Body: getPluginTokenReq{
			PluginID:     r.cli.pluginID,
			PluginSecret: r.cli.pluginSecret,
		},
	}
	resp := new(getPluginTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[LarkProject] Auth#GetPluginToken POST failed: %s", err)
		return nil, response, err
	} else if resp.Error != nil && resp.Error.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[LarkProject] Auth#GetPluginToken POST %s failed, code: %d, msg: %s", uri, resp.Error.Code, resp.Error.Msg)
		return nil, response, NewError("Token", "GetPluginToken", resp.Error.Code, resp.Error.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[LarkProject] Auth#GetPluginToken request_id: %s, response: %s", response.RequestID, jsonString(resp))

	err = r.cli.store.Set(ctx, genPluginTokenKey(r.cli.pluginID), resp.Data.Token, time.Duration(resp.Data.ExpireTime)*time.Second)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[LarkProject] Auth#GetPluginToken set token to storage failed: %s", err)
	}
	return resp.Data, response, nil
}

type TokenExpire struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expire_time"`
}

type GetTokenError struct {
	Code       int64                  `json:"code,omitempty"`
	Msg        string                 `json:"msg,omitempty"`
	DisplayMsg map[string]interface{} `json:"display_msg,omitempty"`
}

type getPluginTokenReq struct {
	PluginID     string `json:"plugin_id,omitempty"`
	PluginSecret string `json:"plugin_secret,omitempty"`
}

type getPluginTokenResp struct {
	Data  *TokenExpire   `json:"data,omitempty"`
	Error *GetTokenError `json:"error,omitempty"`
}

func genPluginTokenKey(pluginID string) string {
	return "lark_project:plugin_token:" + pluginID
}
