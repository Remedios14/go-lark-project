package lark_project

import (
	"context"
	"net/http"
)

func (r *FieldService) GetFieldList(ctx context.Context, request *GetFieldListReq) (*GetFieldListResp, *Response, error) {
	var (
		scope    = "Field"
		funcName = "GetFieldList"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodPost,
			URL:             r.cli.OpenBaseURL + "/open_api/:project_key/field/all",
			Body:            request,
			UserKey:         request.UserKey,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	resp := new(GetFieldListResp)
	httpResp, err := r.cli.RawRequest(ctx, rawReq, resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type GetFieldListReq struct {
	ProjectKey      string `path:"project_key"`
	UserKey         string
	WorkItemTypeKey string `json:"work_item_type_key,omitempty"`
}

type GetFieldListResp struct {
	Data []*Field `json:"data"`
	ErrorData
}

type Field struct {
	FieldKey      string `json:"field_key,omitempty"`
	FieldAlias    string `json:"field_alias,omitempty"`
	FieldTypeKey  string `json:"field_type_key,omitempty"`
	FieldName     string `json:"field_name,omitempty"`
	IsCustomField bool   `json:"is_custom_field,omitempty"`
	IsObsoleted   bool   `json:"is_obsoleted,omitempty"`
}
