package microservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qbox/ke-base/sdk/proto"
	"github.com/qbox/ke-base/sdk/rpc"
	"github.com/qbox/ke-base/sdk/util"
)

type Client interface {
	ListResourceSpecs(ctx context.Context) (ret []proto.ResourceSpec, err error)
	PostApp(ctx context.Context, spaceName string, app proto.App) (ret proto.App, err error)
	ListApp(ctx context.Context, spaceName string) (ret []proto.App, err error)
	GetApp(ctx context.Context, spaceName, appName string) (ret proto.App, err error)
	DeleteApp(ctx context.Context, spaceName, appName string) (err error)

	PostService(ctx context.Context, spaceName, appName string, svc proto.MicroService) (ret proto.MicroService, err error)
	ListService(ctx context.Context, spaceNamee, appName string) (ret []proto.MicroService, err error)
	GetService(ctx context.Context, spaceName, appName, svcName string) (ret proto.MicroService, err error)
	DeleteService(ctx context.Context, spaceName, appName, svcName string) (err error)
	UpgradeService(ctx context.Context, spaceName, appName, svcName string, svc proto.MicroServiceUpgradeArgs) (ret proto.MicroService, err error)
	UpdateServicePorts(ctx context.Context, spaceName, appName, svcName string, svc proto.MicroServiceUpdatePortsArgs) (ret proto.MicroService, err error)
	ScaleService(ctx context.Context, spaceName, appName, svcName string, num int32) (err error)
	StartService(ctx context.Context, spaceName, appName, svcName string) (err error)
	StopService(ctx context.Context, spaceName, appName, svcName string) (err error)
	ListServiceRevisions(ctx context.Context, spaceName, appName, svcName string) (ret []proto.MicroserviceRevision, err error)
	RollbackService(ctx context.Context, spaceName, appName, svcName string, revision int32) (err error)
}

type client struct {
	Client rpc.Client
	Host   string
}

func (p *client) ListResourceSpecs(ctx context.Context) (ret []proto.ResourceSpec, err error) {
	url := fmt.Sprintf("%s/v1/resourcespecs", p.Host)
	err = p.Client.Call(ctx, &ret, "GET", url)
	return
}
func (p *client) PostApp(ctx context.Context, spaceName string, app proto.App) (ret proto.App, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps", p.Host, spaceName)
	err = p.Client.CallWithJson(ctx, &ret, "POST", url, app)
	return
}
func (p *client) ListApp(ctx context.Context, spaceName string) (ret []proto.App, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps", p.Host, spaceName)
	err = p.Client.Call(ctx, &ret, "GET", url)
	return
}
func (p *client) GetApp(ctx context.Context, spaceName, appName string) (ret proto.App, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s", p.Host, spaceName, appName)
	err = p.Client.Call(ctx, &ret, "GET", url)
	return
}
func (p *client) DeleteApp(ctx context.Context, spaceName, appName string) (err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s", p.Host, spaceName, appName)
	err = p.Client.Call(ctx, nil, "DELETE", url)
	return
}

func (p *client) PostService(ctx context.Context, spaceName, appName string, svc proto.MicroService) (ret proto.MicroService, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices", p.Host, spaceName, appName)
	err = p.Client.CallWithJson(ctx, &ret, "POST", url, svc)
	return
}
func (p *client) ListService(ctx context.Context, spaceName, appName string) (ret []proto.MicroService, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices", p.Host, spaceName, appName)
	err = p.Client.Call(ctx, &ret, "GET", url)
	return
}
func (p *client) GetService(ctx context.Context, spaceName, appName, svcName string) (ret proto.MicroService, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s", p.Host, spaceName, appName, svcName)
	err = p.Client.Call(ctx, &ret, "GET", url)
	return
}
func (p *client) DeleteService(ctx context.Context, spaceName, appName, svcName string) (err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s", p.Host, spaceName, appName, svcName)
	err = p.Client.Call(ctx, nil, "DELETE", url)
	return
}
func (p *client) UpgradeService(ctx context.Context, spaceName, appName, svcName string, svc proto.MicroServiceUpgradeArgs) (ret proto.MicroService, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/upgrade", p.Host, spaceName, appName, svcName)
	err = p.Client.CallWithJson(ctx, &ret, "PUT", url, svc)
	return
}
func (p *client) UpdateServicePorts(ctx context.Context, spaceName, appName, svcName string, svc proto.MicroServiceUpdatePortsArgs) (ret proto.MicroService, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/ports", p.Host, spaceName, appName, svcName)
	err = p.Client.CallWithJson(ctx, &ret, "PUT", url, svc)
	return
}
func (p *client) ScaleService(ctx context.Context, spaceName, appName, svcName string, num int32) (err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/scale/%d", p.Host, spaceName, appName, svcName, num)
	err = p.Client.Call(ctx, nil, "PUT", url)
	return
}
func (p *client) StartService(ctx context.Context, spaceName, appName, svcName string) (err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/start", p.Host, spaceName, appName, svcName)
	err = p.Client.Call(ctx, nil, "PUT", url)
	return
}
func (p *client) StopService(ctx context.Context, spaceName, appName, svcName string) (err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/stop", p.Host, spaceName, appName, svcName)
	err = p.Client.Call(ctx, nil, "PUT", url)
	return
}
func (p *client) ListServiceRevisions(ctx context.Context, spaceName, appName, svcName string) (ret []proto.MicroserviceRevision, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/revisions", p.Host, spaceName, appName, svcName)
	err = p.Client.Call(ctx, &ret, "GET", url)
	return
}
func (p *client) RollbackService(ctx context.Context, spaceName, appName, svcName string, revision int32) (err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/apps/%s/microservices/%s/revisions/%d/rollback", p.Host, spaceName, appName, svcName, revision)
	err = p.Client.Call(ctx, nil, "PUT", url)
	return
}

func New(host, token string) Client {
	transport := util.NewAuthTokenTransport(token, http.DefaultTransport)
	return NewWithTransport(host, transport)
}

func NewWithTransport(host string, tr http.RoundTripper) Client {
	c := rpc.Client{&http.Client{Transport: tr}}
	return &client{
		Client: c,
		Host:   util.CleanHost(host),
	}
}
