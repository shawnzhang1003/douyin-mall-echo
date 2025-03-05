package mtl

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName string, metricsPort string, registryAddr string) {
	Registry = prometheus.NewRegistry()
	// go运行时指标
	Registry.MustRegister(collectors.NewGoCollector())
	// 进程指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	/*
		r, _ := etcd.NewEtcdRegistry([]string{registryAddr})

		addr, _ := net.ResolveTCPAddr("tcp", metricsPort)

		registryInfo := &registry.Info{
			ServiceName: "prometheus",
			Addr:        addr,
			Weight:      1,
			Tags:        map[string]string{"service": serviceName},
		}

		err := r.Register(registryInfo)
		if err != nil {
			log.Print(err)
		}

		server.RegisterShutdownHook(func() {
			r.Deregister(registryInfo) //nolint:errcheck
		})
	*/
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil) //nolint:errcheck
}
