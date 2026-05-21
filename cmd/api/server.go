package api

import (
	"fmt"

	"github.com/xxx-newbee/product/internal/config"
	"github.com/xxx-newbee/product/internal/server"
	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configY  string
	StartCmd = &cobra.Command{
		Use:     "service",
		Short:   "start api server",
		Example: "go run product.go service -c /your/config/file.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.Flags().StringVarP(&configY, "config", "c", "etc/product.yaml", "the config file")
}

func setup() {
	conf.MustLoad(configY, &config.C)
}

func run() {
	sctx := svc.NewServiceContext(config.C)

	s := zrpc.MustNewServer(config.C.RpcServerConf, func(grpcServer *grpc.Server) {
		product.RegisterProductServer(grpcServer, server.NewProductServer(sctx))

		if config.C.Mode == service.DevMode || config.C.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer func() {
		s.Stop()
		sctx.MemoryQueue.Shutdown()
		sctx.RedisQueue.Shutdown()
	}()

	fmt.Printf("Starting rpc server at %s...\n", config.C.ListenOn)
	s.Start()
}
