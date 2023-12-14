package lark_project

import (
	"context"
	"net/http"
	"time"
)

// HttpClient http client interface
type HttpClient interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

// LarkProject client struct
type LarkProject struct {
	pluginID        string
	pluginSecret    string
	pluginToken     string
	userPluginToken string
	timeout         time.Duration
	OpenBaseURL     string
	WwwBaseURL      string
	isEnableLogID   bool
	apiMiddlewares  []ApiMiddleware

	httpClient    HttpClient
	logger        Logger
	logLevel      LogLevel
	store         Store
	mock          *Mock
	wrapDoRequest ApiEndpoint

	// service
	Auth     *AuthService
	Space    *SpaceService
	User     *UserService
	WorkItem *WorkItemService
	Field    *FieldService
}

func (r *LarkProject) init() {
	r.wrapDoRequest = chainApiMiddleware(r.apiMiddlewares...)(r.rawRequest)

	r.Auth = &AuthService{cli: r}
	r.Space = &SpaceService{cli: r}
	r.User = &UserService{cli: r}
	r.WorkItem = &WorkItemService{cli: r}
	r.Field = &FieldService{cli: r}
}

func (r *LarkProject) clone() *LarkProject {
	r2 := &LarkProject{
		pluginID:        r.pluginID,
		pluginSecret:    r.pluginSecret,
		pluginToken:     r.pluginToken,
		userPluginToken: r.userPluginToken,
		timeout:         r.timeout,
		OpenBaseURL:     r.OpenBaseURL,
		WwwBaseURL:      r.WwwBaseURL,
		isEnableLogID:   r.isEnableLogID,
		apiMiddlewares:  r.apiMiddlewares,
		httpClient:      r.httpClient,
		logger:          r.logger,
		logLevel:        r.logLevel,
		store:           r.store,
		mock:            r.mock,
		wrapDoRequest:   r.wrapDoRequest,
	}
	r2.init()
	return r2
}

type AuthService struct{ cli *LarkProject }
type SpaceService struct{ cli *LarkProject }
type UserService struct{ cli *LarkProject }
type WorkItemService struct{ cli *LarkProject }
type FieldService struct{ cli *LarkProject }
