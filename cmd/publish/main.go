package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/publish/dal"
	publish "github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish/publishservice"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/bound"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/middleware"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
	tracer2 "github.com/dzc1997/DouyinSimplifyEdition/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.PublishServiceName)
	dal.Init()
	oss.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.PublishAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.PublishServiceName}), //server name                                               // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),                                                // middleWare
		server.WithServiceAddr(addr),                                                                      //address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),                                //limit
		server.WithMuxTransport(),                                                                         //Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                                                   //tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                                               //BoundHandler
		server.WithRegistry(r),                                                                            // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
