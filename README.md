# go-lark-project

访问飞书项目

## Install

```shell
go get github.com/Remedios14/go-lark-project
```

## Usage

### Example: create lark project client

- with plugin id and plugin secret

```go
cli := lark_project.New(WithPluginCredential("<PLUGIN_ID>", "<PLUGIN_SECRET>"))
```

### Example: Query User Detail

```go
cli := lark_project.New(WithPluginCredential("<PLUGIN_ID>", "<PLUGIN_SECRET>"))
resp, _, err := cli.User.QueryUserDetail(context.Background(), &lark_project.QueryUserDetailReq{
    Emails: []string{"<EMAIL>"},
})
if err != nil {
    panic(err)
}
for _, user := range resp.Data {
    fmt.Printf("user key: %s\n", user.UserKey)
}
```
