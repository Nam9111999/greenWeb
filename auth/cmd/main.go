package main

import (
	"context"
	"github.com/soheilhy/cmux"
	"gorm.io/gorm"
	"green.env.com/auth/client/mysql"
	"green.env.com/auth/config"
	serviceHttp "green.env.com/auth/delivery/http"
	"green.env.com/auth/repository"
	"green.env.com/auth/usecase"
	"green.env.com/auth/util"
	"net"
	"time"
)

const VERSION = "1.0.0"

func main() {
	//var cfg = config.GetConfig()

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			util.GetLogger().Fatal(err.Error())
		}

		time.Local = loc
	}

	client := mysql.GetClient
	repo := repository.New(client)
	usecase.New(repo)

}

func executeServer(useCase *usecase.UseCase, client func(ctx context.Context) *gorm.DB) {
	cfg := config.GetConfig()

	// migration
	//migration.Up(client(context.Background()))

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		util.GetLogger().Fatal(err.Error())
	}

	m := cmux.New(l)
	//grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	// gRPC
	//{
	//	s := grpc.NewServer()
	//
	//	grpcServer := &serviceGRPC.TeqService{UseCase: useCase}
	//	proto.RegisterTeqServiceServer(s, grpcServer)
	//
	//	go func() {
	//		errs <- s.Serve(grpcL)
	//	}()
	//}

	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		util.GetLogger().Fatal(err.Error())
	}
}
