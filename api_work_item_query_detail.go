package lark_project

import (
	"context"
	"net/http"
)

// QueryWorkItemDetail 获取工作项详情
// docs: https://project.feishu.cn/openapp/help/articles/840014289372
func (r *WorkItemService) QueryWorkItemDetail(ctx context.Context, req *QueryWorkItemDetailReq) (*QueryWorkItemDetailResp, *Response, error) {
	var (
		scope    = "WorkItem"
		funcName = "QueryWorkItemDetail"
		rawReq   = &RawRequestReq{
			Scope:           scope,
			API:             funcName,
			Method:          http.MethodPost,
			URL:             r.cli.OpenBaseURL + "/open_api/:project_key/work_item/:work_item_type_key/query",
			UserKey:         req.UserKey,
			Body:            req,
			NeedPluginToken: true,
			NeedUserKey:     true,
		}
	)
	var resp = new(QueryWorkItemDetailResp)
	httpResp, err := r.cli.wrapDoRequest(ctx, rawReq, &resp)
	if err != nil {
		return resp, httpResp, err
	} else if resp.ErrCode != 0 {
		return resp, httpResp, NewError(scope, funcName, resp.ErrCode, resp.ErrMsg)
	}
	return resp, httpResp, err
}

type QueryWorkItemDetailReq struct {
	ProjectKey      string `path:"project_key"`
	WorkItemTypeKey string `path:"work_item_type_key"`
	UserKey         string
	WorkItemIDs     []int64              `json:"work_item_ids,omitempty"`
	Fields          []string             `json:"fields,omitempty"` // 需要返回的字段，可以以 - 开头，表示不返回该字段；但不能两种方法混用
	Expand          *WorkItemQueryExpand `json:"expand,omitempty"`
}

type WorkItemQueryExpand struct {
	NeedWorkflow bool `json:"need_workflow,omitempty"` // 是否需要工作流信息（仅流节点）
	// 参考官方文档已经废弃
	//NeedMultiText        bool `json:"need_multi_text,omitempty"`        // 是否需要富文本
	//RelationFieldsDetail bool `json:"relation_fields_detail,omitempty"` // 是否需要关联字段详情
}

type QueryWorkItemDetailResp struct {
	Data []*WorkItemDetail `json:"data,omitempty"`
	ErrorData
}

type WorkItemDetail struct {
	ID              int64             `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	WorkItemTypeKey string            `json:"work_item_type_key,omitempty"`
	ProjectKey      string            `json:"project_key,omitempty"`
	TemplateID      int64             `json:"template_id,omitempty"`
	TemplateType    string            `json:"template_type,omitempty"`
	Pattern         string            `json:"pattern,omitempty"`       // 工作项模式，分为节点模式 Node 和状态模式 State
	SubStage        string            `json:"sub_stage,omitempty"`     // 仅需求
	CurrentNodes    []*WorkItemNode   `json:"current_nodes,omitempty"` // 状态模式节点为空
	StateTimes      []*StateTime      `json:"state_times,omitempty"`
	CreatedBy       string            `json:"created_by,omitempty"` // UserKey
	UpdatedBy       string            `json:"updated_by,omitempty"` // UserKey
	DeletedBy       string            `json:"deleted_by,omitempty"` // UserKey
	CreatedAt       int64             `json:"created_at,omitempty"` // 时间戳，毫秒
	UpdatedAt       int64             `json:"updated_at,omitempty"` // 时间戳，毫秒
	DeletedAt       int64             `json:"deleted_at,omitempty"` // 时间戳，毫秒
	Fields          []*FieldValuePair `json:"fields,omitempty"`
	WorkItemStatus  *WorkItemStatus   `json:"work_item_status,omitempty"`
}

type WorkItemNode struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Owners    []string `json:"owners,omitempty"` // UserKey 列表
	Milestone bool     `json:"milestone,omitempty"`
}

type StateTime struct {
	StateKey  string `json:"state_key,omitempty"`
	StartTime int64  `json:"start_time,omitempty"` // 时间戳，毫秒
	EndTime   int64  `json:"end_time,omitempty"`   // 时间戳，毫秒
	Name      string `json:"name,omitempty"`
}

type FieldValuePair struct {
	FieldKey        string      `json:"field_key,omitempty"`
	FieldAlias      string      `json:"field_alias,omitempty"`
	FieldValue      interface{} `json:"field_value,omitempty"`
	FieldTypeKey    string      `json:"field_type_key,omitempty"`
	HelpDescription string      `json:"help_description,omitempty"`
}

type WorkItemStatus struct {
	StatusKey       string            `json:"status_key,omitempty"`
	IsArchivedState bool              `json:"is_archived_state,omitempty"`
	IsInitState     bool              `json:"is_init_state,omitempty"`
	UpdatedAt       int64             `json:"updated_at,omitempty"` // 时间戳，毫秒
	UpdatedBy       string            `json:"updated_by,omitempty"` // UserKey
	History         []*WorkItemStatus `json:"history,omitempty"`
}

type RoleOwnersValue struct {
	Role   string   `json:"role"`
	Owners []string `json:"owners"` // UserKey 列表
}

type SelectValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Use github.com/mitchellh/mapstructure to decode

type RoleOwnersValueStruct []*RoleOwnersValue
type SelectValueStruct SelectValue
