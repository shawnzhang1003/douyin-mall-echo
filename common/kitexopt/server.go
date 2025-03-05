// Description: rpc服务端公共配置
package kitexopt

import (
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/config-etcd/etcd"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	registryEtcd "github.com/kitex-contrib/registry-etcd"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       []string
	RegistryUsername   string
	RegistryPassword   string
}

// server.Suite is an interface that can be used to configure the kitex server.
var _ server.Suite = CommonServerSuite{}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	// r should not be reused.
	r, err := registryEtcd.NewEtcdRegistry(s.RegistryAddr, registryEtcd.WithAuthOpt(s.RegistryUsername, s.RegistryPassword))
	if err != nil {
		log.Fatal(err)
	}

	opts = append(opts, server.WithRegistry(r))

	// er, err := etcd.NewEtcdResolver(s.RegistryAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	etcdClient, err := etcd.NewClient(etcd.Options{})
	if err != nil {
		log.Fatal(err)
	} else {
		opts = append(opts, server.WithSuite(etcdServer.NewSuite(s.CurrentServiceName, etcdClient)))
	}
	// 提供rpc调用的链路追踪
	_ = provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)
	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		// 禁用kitex server端的prometheus监控
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	)

	return opts
}
