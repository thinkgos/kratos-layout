package router

import (
	v1 "github.com/thinkgos/kratos-layout/genproto/helloworld/v1"
	"google.golang.org/grpc"
)

func RegisterGrpc(srv grpc.ServiceRegistrar, d *DependencyService) {
	v1.RegisterGreeterServer(srv, d.GreeterService)
}
