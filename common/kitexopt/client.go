// Description: rpc客户端公共配置

package kitexopt

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	registryEtcd "github.com/kitex-contrib/registry-etcd"
)

type CommonGrpcClientSuite struct {
	CurrentServiceName string
	RegistryAddr       []string
	RegistryUsername   string
	RegistryPassword   string
}

// client.Suite is an interface that can be used to configure the kitex client.
var _ client.Suite = CommonGrpcClientSuite{}

func (s CommonGrpcClientSuite) Options() []client.Option {
	r, err := registryEtcd.NewEtcdResolver(s.RegistryAddr, registryEtcd.WithAuthOpt(s.RegistryUsername, s.RegistryPassword))
	if err != nil {
		panic(err)
	}
	opts := []client.Option{
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithSuite(tracing.NewClientSuite()),
	)

	return opts
}
