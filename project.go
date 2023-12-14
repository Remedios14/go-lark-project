package lark_project

import (
	"net/http"
	"time"
)

// New Lark-Project client
func New(options ...ClientOptionFunc) *LarkProject {
	return newClient("", options)
}

// ClientOptionFunc new Lark-Project client option
type ClientOptionFunc func(*LarkProject)

// WithPluginCredential set lark-project credential
func WithPluginCredential(pluginID, pluginSecret string) ClientOptionFunc {
	return func(r *LarkProject) {
		r.pluginID = pluginID
		r.pluginSecret = pluginSecret
	}
}

// WithTimeout set timeout
func WithTimeout(timeout time.Duration) ClientOptionFunc {
	return func(r *LarkProject) {
		r.timeout = timeout
	}
}

// WithNetHttpClient set net/http client
func WithNetHttpClient(cli *http.Client) ClientOptionFunc {
	return func(r *LarkProject) {
		r.httpClient = newDefaultHttpClient(cli)
	}
}

// WithHttpClient set http client
func WithHttpClient(cli HttpClient) ClientOptionFunc {
	return func(r *LarkProject) {
		r.httpClient = cli
	}
}

// WithLogger set logger
func WithLogger(logger Logger, level LogLevel) ClientOptionFunc {
	return func(lark *LarkProject) {
		lark.logger = logger
		lark.logLevel = level
	}
}

// WithStore set store TODO: 提供一个使用 redis 存储的 store
func WithStore(store Store) ClientOptionFunc {
	return func(r *LarkProject) {
		r.store = store
	}
}

func newClient(token string, options []ClientOptionFunc) *LarkProject {
	r := &LarkProject{
		pluginToken: token,
		timeout:     time.Second * 3,
		store:       NewStoreMemory(),
		mock:        new(Mock),
		OpenBaseURL: "https://project.feishu.cn",
		WwwBaseURL:  "https://www.feishu.cn",
	}

	for _, option := range options {
		if option != nil {
			option(r)
		}
	}

	if r.httpClient == nil {
		r.httpClient = newDefaultHttpClient(&http.Client{Timeout: r.timeout})
	}

	r.init()

	return r
}
